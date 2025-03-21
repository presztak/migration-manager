package sqlite_test

import (
	"context"
	"crypto/x509"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	dbschema "github.com/FuturFusion/migration-manager/internal/db"
	dbdriver "github.com/FuturFusion/migration-manager/internal/db/sqlite"
	"github.com/FuturFusion/migration-manager/internal/migration"
	endpointMock "github.com/FuturFusion/migration-manager/internal/migration/endpoint/mock"
	"github.com/FuturFusion/migration-manager/internal/migration/repo/sqlite"
	"github.com/FuturFusion/migration-manager/internal/ptr"
	"github.com/FuturFusion/migration-manager/internal/transaction"
	"github.com/FuturFusion/migration-manager/shared/api"
)

var (
	testSource = migration.Source{
		Name:       "TestSource",
		SourceType: api.SOURCETYPE_COMMON,
		Properties: []byte(`{}`),
		EndpointFunc: func(t api.Source) (migration.SourceEndpoint, error) {
			return &endpointMock.SourceEndpointMock{
				ConnectFunc: func(ctx context.Context) error {
					return nil
				},
				DoBasicConnectivityCheckFunc: func() (api.ExternalConnectivityStatus, *x509.Certificate) {
					return api.EXTERNALCONNECTIVITYSTATUS_OK, nil
				},
			}, nil
		},
	}

	testTarget = migration.Target{
		Name:       "TestTarget",
		TargetType: api.TARGETTYPE_INCUS,
		Properties: []byte(`{"endpoint": "https://localhost:6443"}`),
		EndpointFunc: func(t api.Target) (migration.TargetEndpoint, error) {
			return &endpointMock.TargetEndpointMock{
				ConnectFunc: func(ctx context.Context) error {
					return nil
				},
				DoBasicConnectivityCheckFunc: func() (api.ExternalConnectivityStatus, *x509.Certificate) {
					return api.EXTERNALCONNECTIVITYSTATUS_OK, nil
				},
				IsWaitingForOIDCTokensFunc: func() bool {
					return false
				},
			}, nil
		},
	}

	testBatch     = migration.Batch{ID: 1, Name: "TestBatch", TargetID: 1, StoragePool: "", IncludeExpression: "true", MigrationWindowStart: time.Time{}, MigrationWindowEnd: time.Time{}}
	instanceAUUID = uuid.Must(uuid.NewRandom())

	instanceA = migration.Instance{
		Instance: api.Instance{
			UUID:                  instanceAUUID,
			InventoryPath:         "/path/UbuntuVM",
			Annotation:            "annotation",
			MigrationStatus:       api.MIGRATIONSTATUS_NOT_ASSIGNED_BATCH,
			MigrationStatusString: api.MIGRATIONSTATUS_NOT_ASSIGNED_BATCH.String(),
			LastUpdateFromSource:  time.Now().UTC(),
			BatchID:               nil,
			GuestToolsVersion:     123,
			Architecture:          "x86_64",
			HardwareVersion:       "hw version",
			OS:                    "Ubuntu",
			OSVersion:             "24.04",
			Devices:               nil,
			Disks: []api.InstanceDiskInfo{
				{
					Name:                      "disk",
					DifferentialSyncSupported: true,
					SizeInBytes:               123,
				},
			},
			NICs: []api.InstanceNICInfo{
				{
					Network: "net",
					Hwaddr:  "mac",
				},
			},
			Snapshots: nil,
			CPU: api.InstanceCPUInfo{
				NumberCPUs:             2,
				CPUAffinity:            []int32{},
				NumberOfCoresPerSocket: 2,
			},
			Memory: api.InstanceMemoryInfo{
				MemoryInBytes:            4294967296,
				MemoryReservationInBytes: 4294967296,
			},
			UseLegacyBios:     false,
			SecureBootEnabled: false,
			TPMPresent:        false,
		},
		NeedsDiskImport: false,
		SecretToken:     uuid.Must(uuid.NewRandom()),
		SourceID:        1,
	}

	instanceBUUID = uuid.Must(uuid.NewRandom())
	instanceB     = migration.Instance{
		Instance: api.Instance{
			UUID:                  instanceBUUID,
			InventoryPath:         "/path/WindowsVM",
			Annotation:            "annotation",
			MigrationStatus:       api.MIGRATIONSTATUS_NOT_ASSIGNED_BATCH,
			MigrationStatusString: api.MIGRATIONSTATUS_NOT_ASSIGNED_BATCH.String(),
			LastUpdateFromSource:  time.Now().UTC(),
			BatchID:               nil,
			GuestToolsVersion:     123,
			Architecture:          "x86_64",
			HardwareVersion:       "hw version",
			OS:                    "Windows",
			OSVersion:             "11",
			Devices:               nil,
			Disks: []api.InstanceDiskInfo{
				{
					Name:                      "disk",
					DifferentialSyncSupported: false,
					SizeInBytes:               321,
				},
			},
			NICs: []api.InstanceNICInfo{
				{
					Network: "net1",
					Hwaddr:  "mac1",
				},
				{
					Network: "net2",
					Hwaddr:  "mac2",
				},
			},
			Snapshots: nil,
			CPU: api.InstanceCPUInfo{
				NumberCPUs:             2,
				CPUAffinity:            []int32{0, 1},
				NumberOfCoresPerSocket: 2,
			},
			Memory: api.InstanceMemoryInfo{
				MemoryInBytes:            4294967296,
				MemoryReservationInBytes: 4294967296,
			},
			UseLegacyBios:     false,
			SecureBootEnabled: true,
			TPMPresent:        true,
		},
		NeedsDiskImport: false,
		SecretToken:     uuid.Must(uuid.NewRandom()),
		SourceID:        1,
	}

	instanceCUUID = uuid.Must(uuid.NewRandom())
	instanceC     = migration.Instance{
		Instance: api.Instance{
			UUID:                  instanceCUUID,
			InventoryPath:         "/path/DebianVM",
			Annotation:            "annotation",
			MigrationStatus:       api.MIGRATIONSTATUS_NOT_ASSIGNED_BATCH,
			MigrationStatusString: api.MIGRATIONSTATUS_NOT_ASSIGNED_BATCH.String(),
			LastUpdateFromSource:  time.Now().UTC(),
			BatchID:               ptr.To(1),
			GuestToolsVersion:     123,
			Architecture:          "arm64",
			HardwareVersion:       "hw version",
			OS:                    "Debian",
			OSVersion:             "bookworm",
			Devices:               nil,
			Disks: []api.InstanceDiskInfo{
				{
					Name:                      "disk1",
					DifferentialSyncSupported: true,
					SizeInBytes:               123,
				},
				{
					Name:                      "disk2",
					DifferentialSyncSupported: true,
					SizeInBytes:               321,
				},
			},
			NICs:      nil,
			Snapshots: nil,
			CPU: api.InstanceCPUInfo{
				NumberCPUs:             4,
				CPUAffinity:            []int32{0, 1, 2, 3},
				NumberOfCoresPerSocket: 2,
			},
			Memory: api.InstanceMemoryInfo{
				MemoryInBytes:            4294967296,
				MemoryReservationInBytes: 4294967296,
			},
			UseLegacyBios:     true,
			SecureBootEnabled: false,
			TPMPresent:        false,
		},
		NeedsDiskImport: false,
		SecretToken:     uuid.Must(uuid.NewRandom()),
		SourceID:        1,
	}
)

func TestInstanceDatabaseActions(t *testing.T) {
	ctx := context.Background()

	// Create a new temporary database.
	tmpDir := t.TempDir()
	db, err := dbdriver.Open(tmpDir)
	require.NoError(t, err)

	t.Cleanup(func() {
		err = db.Close()
		require.NoError(t, err)
	})

	_, err = dbschema.EnsureSchema(db, tmpDir)
	require.NoError(t, err)

	sourceSvc := migration.NewSourceService(sqlite.NewSource(db))
	targetSvc := migration.NewTargetService(sqlite.NewTarget(db))

	instance := sqlite.NewInstance(db)
	instanceSvc := migration.NewInstanceService(instance, sourceSvc)

	batchSvc := migration.NewBatchService(sqlite.NewBatch(db), instanceSvc)

	// Cannot add an instance with an invalid source.
	_, err = instance.Create(ctx, instanceA)
	require.ErrorIs(t, err, migration.ErrConstraintViolation)

	// Add dummy source.
	_, err = sourceSvc.Create(ctx, testSource)
	require.NoError(t, err)

	// Add dummy target.
	_, err = targetSvc.Create(ctx, testTarget)
	require.NoError(t, err)

	// Add dummy batch.
	_, err = batchSvc.Create(ctx, testBatch)
	require.NoError(t, err)

	// Add instanceA.
	instanceA, err = instance.Create(ctx, instanceA)
	require.NoError(t, err)

	// Add instanceB.
	instanceB, err = instance.Create(ctx, instanceB)
	require.NoError(t, err)

	// Add instanceC.
	instanceC, err = instance.Create(ctx, instanceC)
	require.NoError(t, err)

	// Cannot delete a source or target if referenced by an instance.
	err = sourceSvc.DeleteByName(context.TODO(), testSource.Name)
	require.ErrorIs(t, err, migration.ErrConstraintViolation)
	err = targetSvc.DeleteByName(ctx, testTarget.Name)
	require.ErrorIs(t, err, migration.ErrConstraintViolation)

	// Ensure we have three instances.
	instances, err := instance.GetAll(ctx)
	require.NoError(t, err)
	require.Len(t, instances, 3)

	// Should get back instanceA unchanged.
	dbInstanceA, err := instance.GetByID(ctx, instanceA.UUID)
	require.NoError(t, err)
	require.Equal(t, instanceA, dbInstanceA)

	// Test updating an instance.
	instanceB.InventoryPath = "/foo/bar"
	instanceB.CPU.NumberCPUs = 8
	instanceB.MigrationStatus = api.MIGRATIONSTATUS_BACKGROUND_IMPORT
	instanceB.MigrationStatusString = instanceB.MigrationStatus.String()
	instanceB, err = instance.UpdateByID(ctx, instanceB)
	require.NoError(t, err)
	dbInstanceB, err := instance.GetByID(ctx, instanceB.UUID)
	require.NoError(t, err)
	require.Equal(t, instanceB, dbInstanceB)

	// Delete an instance.
	err = instance.DeleteByID(ctx, instanceA.UUID)
	require.NoError(t, err)
	_, err = instance.GetByID(ctx, instanceA.UUID)
	require.ErrorIs(t, err, migration.ErrNotFound)

	// Should have two instances remaining.
	instances, err = instance.GetAll(ctx)
	require.NoError(t, err)
	require.Len(t, instances, 2)

	// Can't delete an instance that doesn't exist.
	randomUUID, _ := uuid.NewRandom()
	err = instance.DeleteByID(ctx, randomUUID)
	require.ErrorIs(t, err, migration.ErrNotFound)

	// Can't update an instance that doesn't exist.
	_, err = instance.UpdateByID(ctx, instanceA)
	require.ErrorIs(t, err, migration.ErrNotFound)

	// Can't add a duplicate instance.
	_, err = instance.Create(ctx, instanceB)
	require.ErrorIs(t, err, migration.ErrConstraintViolation)

	// Can't delete a source that has at least one associated instance.
	err = sourceSvc.DeleteByName(ctx, testSource.Name)
	require.ErrorIs(t, err, migration.ErrConstraintViolation)

	// Can't delete a target that has at least one associated instance.
	err = targetSvc.DeleteByName(ctx, testTarget.Name)
	require.ErrorIs(t, err, migration.ErrConstraintViolation)
}

var overridesA = migration.Overrides{InstanceOverride: api.InstanceOverride{UUID: instanceAUUID, LastUpdate: time.Now().UTC(), Comment: "A comment", NumberCPUs: 8, MemoryInBytes: 4096, DisableMigration: true}}

func TestInstanceOverridesDatabaseActions(t *testing.T) {
	ctx := context.Background()

	// Create a new temporary database.
	tmpDir := t.TempDir()
	db, err := dbdriver.Open(tmpDir)
	require.NoError(t, err)

	t.Cleanup(func() {
		err = db.Close()
		require.NoError(t, err)
	})

	_, err = dbschema.EnsureSchema(db, tmpDir)
	require.NoError(t, err)

	sourceSvc := migration.NewSourceService(sqlite.NewSource(db))
	targetSvc := migration.NewTargetService(sqlite.NewTarget(db))

	instance := sqlite.NewInstance(db)

	// Cannot add an overrides if there's no corresponding instance.
	_, err = instance.CreateOverrides(ctx, overridesA)
	require.ErrorIs(t, err, migration.ErrConstraintViolation)

	// Add the corresponding instance.
	_, err = sourceSvc.Create(ctx, testSource)
	require.NoError(t, err)
	_, err = targetSvc.Create(ctx, testTarget)
	require.NoError(t, err)
	_, err = instance.Create(ctx, instanceA)
	require.NoError(t, err)

	// Add the overrides.
	overridesA, err = instance.CreateOverrides(ctx, overridesA)
	require.NoError(t, err)

	// Should get back overridesA unchanged.
	dbOverridesA, err := instance.GetOverridesByID(ctx, instanceA.UUID)
	require.NoError(t, err)
	require.Equal(t, overridesA, dbOverridesA)

	// The Instance's returned overrides should match what we set.
	dbInstanceA, err := instance.GetByID(ctx, instanceA.UUID)
	require.NoError(t, err)
	require.Equal(t, *dbInstanceA.Overrides, overridesA.InstanceOverride)

	// Test updating an overrides.
	overridesA.Comment = "An update"
	overridesA.DisableMigration = false
	overridesA, err = instance.UpdateOverridesByID(ctx, overridesA)
	require.NoError(t, err)
	dbOverridesA, err = instance.GetOverridesByID(ctx, instanceA.UUID)
	require.NoError(t, err)
	require.Equal(t, overridesA, dbOverridesA)

	// Can't add a duplicate overrides.
	_, err = instance.CreateOverrides(ctx, overridesA)
	require.ErrorIs(t, err, migration.ErrConstraintViolation)

	// Delete an overrides.
	err = instance.DeleteOverridesByID(ctx, instanceA.UUID)
	require.NoError(t, err)
	_, err = instance.GetOverridesByID(ctx, instanceA.UUID)
	require.ErrorIs(t, err, migration.ErrNotFound)

	// Can't delete an overrides that doesn't exist.
	randomUUID := uuid.Must(uuid.NewRandom())
	err = instance.DeleteOverridesByID(ctx, randomUUID)
	require.ErrorIs(t, err, migration.ErrNotFound)

	// Can't update an overrides that doesn't exist.
	_, err = instance.UpdateOverridesByID(ctx, overridesA)
	require.ErrorIs(t, err, migration.ErrNotFound)

	// Ensure deletion of instance fails, if an overrides is present
	// (cascading delete is handled by the business logic and not the DB layer).
	_, err = instance.CreateOverrides(ctx, overridesA)
	require.NoError(t, err)
	_, err = instance.GetOverridesByID(ctx, instanceA.UUID)
	require.NoError(t, err)
	err = instance.DeleteByID(ctx, instanceA.UUID)
	require.ErrorIs(t, err, migration.ErrConstraintViolation)
}

func TestInstanceGetAll(t *testing.T) {
	ctx := context.Background()

	// Create a new temporary database.
	tmpDir := t.TempDir()
	db, err := dbdriver.Open(tmpDir)
	require.NoError(t, err)

	t.Cleanup(func() {
		err = db.Close()
		require.NoError(t, err)
	})

	_, err = dbschema.EnsureSchema(db, tmpDir)
	require.NoError(t, err)

	dbWithTransaction := transaction.Enable(db)

	sourceSvc := migration.NewSourceService(sqlite.NewSource(dbWithTransaction))
	targetSvc := migration.NewTargetService(sqlite.NewTarget(dbWithTransaction))

	instance := sqlite.NewInstance(dbWithTransaction)

	// Add dummy source.
	_, err = sourceSvc.Create(ctx, testSource)
	require.NoError(t, err)

	// Add dummy target.
	_, err = targetSvc.Create(ctx, testTarget)
	require.NoError(t, err)

	const maxInstances = 100

	// Add instanceA.
	for i := 0; i < maxInstances; i++ {
		instanceN := instanceA
		instanceN.UUID = uuid.Must(uuid.NewRandom())
		instanceN.InventoryPath = fmt.Sprintf("/%d", i)

		_, err = instance.Create(ctx, instanceN)
		require.NoError(t, err)

		overrideN := overridesA
		overrideN.UUID = instanceN.UUID
		_, err = instance.CreateOverrides(ctx, overrideN)
		require.NoError(t, err)
	}

	ctx2 := context.Background()
	_ = transaction.Do(ctx2, func(ctx context.Context) error {
		instances, err := instance.GetAll(ctx2)
		require.NoError(t, err)
		require.Len(t, instances, maxInstances)

		return nil
	})
}
