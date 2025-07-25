package worker

import (
	"bufio"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/lxc/incus/v6/shared/subprocess"
	"github.com/lxc/incus/v6/shared/util"

	internalUtil "github.com/FuturFusion/migration-manager/internal/util"
)

//go:embed scripts/*
var embeddedScripts embed.FS

const (
	PARTITION_TYPE_UNKNOWN = iota
	PARTITION_TYPE_PLAIN
	PARTITION_TYPE_LVM
)

type LSBLKOutput struct {
	BlockDevices []struct {
		Name     string `json:"name"`
		Serial   string `json:"serial"`
		Children []struct {
			Name         string `json:"name"`
			FSType       string `json:"fstype"`
			PartLabel    string `json:"partlabel"`
			PartTypeName string `json:"parttypename"`
		} `json:"children"`
	} `json:"blockdevices"`
}

type LVSOutput struct {
	Report []struct {
		LV []struct {
			VGName string `json:"vg_name"`
			LVName string `json:"lv_name"`
		} `json:"lv"`
	} `json:"report"`
}

const chrootMountPath string = "/run/mount/target/"

func LinuxDoPostMigrationConfig(ctx context.Context, distro string, majorVersion int) error {
	slog.Info("Preparing to perform post-migration configuration of VM")

	// Determine the root partition.
	rootPartition, rootPartitionType, rootMountOpts, err := determineRootPartition()
	if err != nil {
		return err
	}

	// Activate VG prior to mounting, if needed.
	if rootPartitionType == PARTITION_TYPE_LVM {
		err := ActivateVG()
		if err != nil {
			return err
		}

		defer func() { _ = DeactivateVG() }()
	}

	// Mount the migrated root partition.
	err = DoMount(rootPartition, chrootMountPath, rootMountOpts)
	if err != nil {
		return err
	}

	defer func() { _ = DoUnmount(chrootMountPath) }()

	// Bind-mount /dev/, /proc/ and /sys/ into the chroot.
	err = DoMount("/dev/", filepath.Join(chrootMountPath, "dev"), []string{"-o", "bind"})
	if err != nil {
		return err
	}

	defer func() { _ = DoUnmount(filepath.Join(chrootMountPath, "dev")) }()

	err = DoMount("/proc/", filepath.Join(chrootMountPath, "proc"), []string{"-o", "bind"})
	if err != nil {
		return err
	}

	defer func() { _ = DoUnmount(filepath.Join(chrootMountPath, "proc")) }()

	err = DoMount("/sys/", filepath.Join(chrootMountPath, "sys"), []string{"-o", "bind"})
	if err != nil {
		return err
	}

	defer func() { _ = DoUnmount(filepath.Join(chrootMountPath, "sys")) }()

	// Mount additional file systems, such as /var/ on a different partition.
	for _, mnt := range getAdditionalMounts() {
		opts := []string{}
		if mnt["options"] != "" {
			opts = []string{"-o", mnt["options"]}
		}

		err := DoMount(mnt["device"], filepath.Join(chrootMountPath, mnt["path"]), opts)
		if err != nil {
			return err
		}

		defer func() { _ = DoUnmount(filepath.Join(chrootMountPath, mnt["path"])) }() //nolint: revive
	}

	// Install incus-agent into the VM.
	err = runScriptInChroot("install-incus-agent.sh")
	if err != nil {
		return err
	}

	// Remove any open-vm-tools packing that might be installed.
	if internalUtil.IsDebianOrDerivative(distro) {
		err := runScriptInChroot("debian-purge-open-vm-tools.sh")
		if err != nil {
			return err
		}
	} else if internalUtil.IsRHELOrDerivative(distro) {
		err := runScriptInChroot("redhat-purge-open-vm-tools.sh")
		if err != nil {
			return err
		}
	} else if internalUtil.IsSUSEOrDerivative(distro) {
		err := runScriptInChroot("suse-purge-open-vm-tools.sh")
		if err != nil {
			return err
		}
	}

	// Add the virtio drivers if needed.
	if internalUtil.IsRHELOrDerivative(distro) || internalUtil.IsSUSEOrDerivative(distro) {
		err := runScriptInChroot("dracut-add-virtio-drivers.sh")
		if err != nil {
			return err
		}
	}

	c := internalUtil.UnixHTTPClient("/dev/incus/sock")
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://unix.socket/1.0/config/user.migration.hwaddrs", nil)
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer func() { _ = resp.Body.Close() }()
	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Setup udev rules to create network device aliases.
	err = runScriptInChroot("add-udev-network-rules.sh", string(out))
	if err != nil {
		return err
	}

	// Setup incus-agent service override for older versions of RHEL.
	if internalUtil.IsRHELOrDerivative(distro) && majorVersion <= 7 {
		err := runScriptInChroot("add-incus-agent-override-for-old-systemd.sh")
		if err != nil {
			return err
		}
	}

	slog.Info("Post-migration configuration complete!")
	return nil
}

func ActivateVG() error {
	_, err := subprocess.RunCommand("vgchange", "-a", "y")
	return err
}

func DeactivateVG() error {
	_, err := subprocess.RunCommand("vgchange", "-a", "n")
	return err
}

func DetermineWindowsPartitions() (base string, recovery string, err error) {
	partitions, err := scanPartitions("")
	if err != nil {
		return "", "", err
	}

	for _, dev := range partitions.BlockDevices {
		if dev.Serial != "incus_root" {
			continue
		}

		for _, child := range dev.Children {
			if child.PartLabel == "Basic data partition" && child.PartTypeName == "Microsoft basic data" {
				base = child.Name
			} else if child.PartTypeName == "Windows recovery environment" {
				recovery = child.Name
			}
		}
	}

	if base == "" || recovery == "" {
		b, err := json.Marshal(partitions)
		if err != nil {
			return "", "", err
		}

		return "", "", fmt.Errorf("Could not determine partitions: %v", string(b))
	}

	return base, recovery, nil
}

func determineRootPartition() (string, int, []string, error) {
	lvs, err := scanVGs()
	if err != nil {
		return "", PARTITION_TYPE_UNKNOWN, nil, err
	}

	// If a VG(s) exists, check if any LVs look like the root partition.
	if len(lvs.Report[0].LV) > 0 {
		err := ActivateVG()
		if err != nil {
			return "", PARTITION_TYPE_UNKNOWN, nil, err
		}

		defer func() { _ = DeactivateVG() }()

		for _, lv := range lvs.Report[0].LV {
			if looksLikeRootPartition(fmt.Sprintf("/dev/%s/%s", lv.VGName, lv.LVName), nil) {
				return fmt.Sprintf("/dev/%s/%s", lv.VGName, lv.LVName), PARTITION_TYPE_LVM, nil, nil
			}
		}
	}

	partitions, err := scanPartitions("")
	if err != nil {
		return "", PARTITION_TYPE_UNKNOWN, nil, err
	}

	for _, dev := range partitions.BlockDevices {
		if dev.Serial != "incus_root" {
			continue
		}

		// Loop through any partitions on /dev/sda and check if they look like the root partition.
		for _, p := range dev.Children {
			partition := fmt.Sprintf("/dev/%s", p.Name)
			if p.FSType == "btrfs" {
				btrfsSubvol, err := getBTRFSTopSubvol(partition)
				if err != nil {
					return "", PARTITION_TYPE_UNKNOWN, nil, err
				}

				opts := []string{"-o", fmt.Sprintf("subvol=%s", btrfsSubvol)}
				if looksLikeRootPartition(partition, opts) {
					return partition, PARTITION_TYPE_PLAIN, opts, nil
				}
			} else if looksLikeRootPartition(partition, nil) {
				return partition, PARTITION_TYPE_PLAIN, nil, nil
			}
		}
	}

	return "", PARTITION_TYPE_UNKNOWN, nil, fmt.Errorf("Failed to determine the root partition")
}

func runScriptInChroot(scriptName string, args ...string) error {
	// Get the embedded script's contents.
	script, err := embeddedScripts.ReadFile(filepath.Join("scripts/", scriptName))
	if err != nil {
		return err
	}

	// Write script to tmp file.
	err = os.WriteFile(filepath.Join(chrootMountPath, scriptName), script, 0o755)
	if err != nil {
		return err
	}

	defer func() { _ = os.Remove(filepath.Join(chrootMountPath, scriptName)) }()

	// Run the script within the chroot.
	cmd := []string{chrootMountPath, filepath.Join("/", scriptName)}
	cmd = append(cmd, args...)
	_, err = subprocess.RunCommand("chroot", cmd...)
	return err
}

func scanVGs() (LVSOutput, error) {
	ret := LVSOutput{}
	output, err := subprocess.RunCommand("lvs", "-o", "vg_name,lv_name", "--reportformat", "json")
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal([]byte(output), &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func scanPartitions(device string) (LSBLKOutput, error) {
	ret := LSBLKOutput{}
	args := []string{"-J", "-o", "NAME,FSTYPE,PARTLABEL,PARTTYPENAME,SERIAL"}
	if device != "" {
		args = append(args, device)
	}

	output, err := subprocess.RunCommand("lsblk", args...)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal([]byte(output), &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func looksLikeRootPartition(partition string, opts []string) bool {
	// Mount the potential root partition.
	err := DoMount(partition, chrootMountPath, opts)
	if err != nil {
		return false
	}

	defer func() { _ = DoUnmount(chrootMountPath) }()

	// If /usr/ and /etc/ exist, this is probably the root partition.
	return util.PathExists(filepath.Join(chrootMountPath, "usr")) && util.PathExists(filepath.Join(chrootMountPath, "etc"))
}

func getAdditionalMounts() []map[string]string {
	ret := []map[string]string{}

	fstab, err := os.Open(filepath.Join(chrootMountPath, "etc/fstab"))
	if err != nil {
		return ret
	}

	defer func() { _ = fstab.Close() }()

	sc := bufio.NewScanner(fstab)
	for sc.Scan() {
		text := strings.TrimSpace(sc.Text())

		if len(text) > 0 && !strings.HasPrefix(text, "#") {
			fields := regexp.MustCompile(`\s+`).Split(text, -1)
			if strings.HasPrefix(fields[1], "/boot") || strings.HasPrefix(fields[1], "/var") || strings.HasPrefix(fields[1], "/usr") {
				ret = append(ret, map[string]string{"device": fields[0], "path": fields[1], "options": fields[3]})
			}
		}
	}

	return ret
}

func getBTRFSTopSubvol(partition string) (string, error) {
	// Mount the partition so we can get the list of subvolumes.
	err := DoMount(partition, chrootMountPath, nil)
	if err != nil {
		return "", err
	}

	defer func() { _ = DoUnmount(chrootMountPath) }()

	// Get the subvolumes.
	output, err := subprocess.RunCommand("btrfs", "subvolume", "list", chrootMountPath)
	if err != nil {
		return "", err
	}

	// Get the top level subvolume.
	submatch := regexp.MustCompile(` top level 5 path (.+)`).FindStringSubmatch(output)
	if submatch != nil {
		return submatch[1], nil
	}

	return "", fmt.Errorf("Unable to determine top level subvolume for partition %s", partition)
}
