// Copyright (c) 2026 Business Techologies
// SPDX-License-Identifier: MIT
package iam

import (
	"time"
)

type (
	JWK struct {
		Kty string `json:"kty,omitempty"` // type
		Use string `json:"use,omitempty"` // use: jwt, etc...
		Alg string `json:"alg,omitempty"` // algoritmh
		Kid string `json:"kid,omitempty"` // key id
		Crv string `json:"crv,omitempty"` // curve
		X   string `json:"x,omitempty"`   // .
	}
)

func NewJWK() *JWK {
	jwk := &JWK{}

	return jwk
}

func (jwk *JWK) Now() {
	jwk.SetKid(time.Now())
}

func (jwk *JWK) SetKid(ts time.Time) {
	jwk.Kid = ts.Format("2006-01-02T15:04:05Z-0700")
}
