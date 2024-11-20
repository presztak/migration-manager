package api

import (
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

// IncusTarget defines an Incus target for use by the migration manager.
//
// swagger:model
type IncusTarget struct {
	// A human-friendly name for this target
	// Example: MyTarget
	Name string `json:"name" yaml:"name"`

	// An opaque integer identifier for the target
	// Example: 123
	DatabaseID int `json:"databaseID" yaml:"databaseID"`

	// Hostname or IP address of the target endpoint
	// Example: https://incus.local:8443
	Endpoint string `json:"endpoint" yaml:"endpoint"`

	// base64-encoded TLS client key for authentication
	TLSClientKey string `json:"tlsClientKey" yaml:"tlsClientKey"`

	// base64-encoded TLS client certificate for authentication
	TLSClientCert string `json:"tlsClientCert" yaml:"tlsClientCert"`

	// OpenID Connect tokens
	OIDCTokens *oidc.Tokens[*oidc.IDTokenClaims] `json:"oidcTokens" yaml:"oidcTokens"`

	// If true, disable TLS certificate validation
	// Example: false
	Insecure bool `json:"insecure" yaml:"insecure"`

	// The Incus profile to use
	// Example: default
	IncusProfile string `json:"incusProfile" yaml:"incusProfile"`

	// The Incus project to use
	// Example: default
	IncusProject string `json:"incusProject" yaml:"incusProject"`

	// The storage pool that holds various migration ISO images.
	// Example: local
	StoragePool string `json:"storagePool" yaml:"storagePool"`

	// The name of the boot environment ISO image.
	// Example: migration-manager-minimal-boot.iso
	BootISOImage string `json:"bootIsoImage" yaml:"bootIsoImage"`

	// The name of the virtio drivers ISO image.
	// Example: virtio-win.iso
	DriversISOImage string `json:"driversIdoImage" yaml:"DriversISOImage"`
}
