package config

import (
	"encoding/json"
	"fmt"
)

// Load a configuration
func Load(data []byte) (*Config, error) {
	var cfg Config
	err := json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("parsing configuration: %w", err)
	}

	return &cfg, nil
}

// Config values, as loaded from the configuration file
type Config struct {
	// Certs configures the demo certs to request and serve
	Certs []Cert `json:"certs"`

	// Acme configures the CA to use
	Acme Acme `json:"acme"`
}

type Cert struct {
	// Hostname that this certificate should contain.
	// Eg, revoked.example.org
	Hostname string `json:"hostname"`

	// Revoked set to true means this cert will be revoked
	Revoked bool `json:"revoked"`

	// Expired set to true means this cert will be expired
	Expired bool `json:"expired"`

	// Profile is the ACME profile to use
	// Eg, shortlived at Let's Encrypt will get a 7-day cert,
	// which is useful for speeding up expiry.
	Profile string `json:"profile"`

	// Root that this entry must chain to.
	// Used to select among multiple available chains from the ACME server.
	// Eg, ISRG Root X1
	Root string `json:"root"`

	// CertPath is where the certificate is stored
	// Loaded on startup if exists, otherwise will be created.
	// Overwritten on renewal.
	// Eg, /etc/demo/certs/revoked.crt
	CertPath string `json:"cert_path"`

	// KeyPath is where the corresponding key is stored
	// Loaded on startup if exists, otherwise will be generated.
	// Eg, /etc/demo/keys/rsa.key
	KeyPath string `json:"key_path"`

	// KeyType is the type of key to use: RSA2048 or P256.
	// Eg, P256
	KeyType string `json:"key_type"`
}

type Acme struct {
	// Directory URL of the server
	// Eg, https://acme-v02.api.letsencrypt.org/directory
	Directory string `json:"directory"`

	// KeyPath is the ACME account key
	// Generated if it doesn't exist
	// Eg, /etc/demo/keys/account.key
	KeyPath string `json:"key_path"`
}
