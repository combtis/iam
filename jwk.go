package iam

import (
	"time"
)

type (
	JWK struct {
		Kty string `json:"kty,omitempty"`
		Kid string `json:"kid,omitempty"`
	}
)

func NewJWK() *JWK {
	jwk := &JWK{}

	return jwk
}

func (jwk *JWK) SetKid(ts time.Time) {
	jwk.Kid = ts.Format("20060102150405")
}
