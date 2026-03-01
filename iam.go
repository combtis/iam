// Copyright (c) 2026 Business Technologies
// SDPX-Licencse-Identifier: MIT
package iam

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"net/http"
	"os"
)

var (
	site = os.Getenv("IAM_SITE")
)

type (
	OpenIDConfig struct {
		Endpoint  string `json:"endpoint"`
		PublicKey []byte `json:"public_key"`
	}
)

func NewOpenIDConfig() *OpenIDConfig {
	oidc := &OpenIDConfig{}

	return oidc
}

func (oidc *OpenIDConfig) GenerateKey() (pub ed25519.PublicKey, priv ed25519.PrivateKey, err error) {
	pub, priv, err = ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return
}

func (oidc *OpenIDConfig) LoadPrivateKey(filename string) error {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0o644)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(f)
	f.Close()

	block, rest := pem.Decode(buf.Bytes())
	if block == nil {
		return errors.New("err: file not pem format")
	}
	_ = rest

	if block.Type != "PUBLIC KEY" {
		return errors.New("err: file not public key")
	}

	oidc.PublicKey = []byte(base64.RawURLEncoding.EncodeToString([]byte(block.Bytes)))

	return nil
}

func (oidc *OpenIDConfig) ToJSON() ([]byte, error) {
	return json.Marshal(oidc)
}

func OpenIDConfiguration(rw http.ResponseWriter, req *http.Request) {
	code := http.StatusOK

	if req.Method != http.MethodGet {
		code = http.StatusMethodNotAllowed
		http.Error(rw, http.StatusText(code), code)
	}

	if site == "" {
		code = http.StatusInternalServerError
		http.Error(rw, http.StatusText(code), code)
		return
	}

	rw.WriteHeader(code)
	rw.Write([]byte(`# test: not prod
endpoint: http://` + site + `/auth/admin
`))
}
