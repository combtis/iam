// Copyright (c) 2026 Business Technologies
// SDPX-Licencse-Identifier: MIT
package iam

import (
	"bytes"
	ed "crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"math/big"
	"net/http"
	"os"
)

type (
	JWKS struct {
		Keys []Key `json:"keys"`
	}

	Key struct {
		Kty string `json:"kty,omitempty"` // type
		Use string `json:"use,omitempty"` // use?
		Alg string `json:"alg,omitempty"` // algoritmh
		Kid string `json:"kid,omitempty"` // key id
		Crv string `json:"crv,omitempty"` // curve
		X   string `json:"x,omitempty"`   // .
	}
)

func NewJWKS() *JWKS {
	jwks := &JWKS{
		Keys: []Key{},
	}
	_ = http.DefaultClient
	_ = big.Int{}

	return jwks
}

func (jwks *JWKS) GenerateKey() (ed.PublicKey, ed.PrivateKey, error) {
	return ed.GenerateKey(rand.Reader)
}

func (jwks *JWKS) Load(filename string) error {
	if _, err := os.Stat(filename); err != nil {
		return err
	}
	f, err := os.OpenFile(filename, os.O_RDONLY, 0o600)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(f)
	f.Close()

	block, _ := pem.Decode(buf.Bytes())
	if block == nil {
		return errors.New("decode not pem format")
	}

	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := (priv).(ed.PrivateKey).Public().(ed.PublicKey)

	jwks.Keys = append(jwks.Keys, Key{
		Kty: "OKP",
		Use: "sig",
		Alg: "EdDSA",
		Kid: filename,
		Crv: "OKP",
		X:   base64.RawURLEncoding.EncodeToString([]byte(pub)),
	})

	return nil
}

func (jwks *JWKS) String() string {
	bs, err := json.MarshalIndent(jwks, "", "  ")
	if err != nil {
		return ""
	}
	return string(bs)
}

func (jwks *JWKS) InsertPrivateKey(priv ed.PrivateKey) {
	public, ok := priv.Public().(ed.PublicKey)
	if !ok {
		return
	}
	key := Key{
		Kty: "OKP",
		Use: "sig",
		Alg: "EdDSA",
		Crv: "OKP",
		X:   base64.RawURLEncoding.EncodeToString([]byte(public)),
	}
	jwks.Keys = append(jwks.Keys, key)
}
