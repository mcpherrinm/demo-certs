package config

import (
	"reflect"
	"testing"
)

const testcfg = `
{
	"certs": [
		{
			"hostname": "expired.example.com",
			"expired": true,
			"profile": "shortlived",
			"root": "ISRG Root X1",
			"cert_path": "/etc/demo/certs/expired-x1.crt",
			"key_path": "/etc/demo/keys/expired-x1.key",	
			"key_type": "P256"
		}
	],
	"acme": {
		"directory": "https://acme-v02.api.letsencrypt.org/directory",
		"key_path": "/etc/demo/keys/account.key"
	}
}

`

func TestLoad(t *testing.T) {
	expected := Config{
		Certs: []Cert{
			{
				Hostname: "expired.example.com",
				Expired:  true,
				Profile:  "shortlived",
				Root:     "ISRG Root X1",
				CertPath: "/etc/demo/certs/expired-x1.crt",
				KeyPath:  "/etc/demo/keys/expired-x1.key",
				KeyType:  "P256",
			},
		},
		Acme: Acme{
			Directory: "https://acme-v02.api.letsencrypt.org/directory",
			KeyPath:   "/etc/demo/keys/account.key",
		},
	}

	cfg, err := Load([]byte(testcfg))
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, *cfg) {
		t.Fatal("expected", expected, "got", *cfg)
	}
}
