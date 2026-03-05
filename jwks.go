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

type (
	JWKS struct {
		Keys []Key `json:"keys"`
	}

	Key struct {
		Alg string `json:"alg"`
		E   string `json:"e"`
		Kid string `json:"kid"`
		Kty string `json:"kty"`
		N   string `json:"n"`
		Use string `json:"use"`
	}
)
