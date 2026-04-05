// Copyright (c) 2026 Business Techologies
// SPDX-License-Identifier: MIT
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
	jwk.Kid = ts.Format("2006-01-02T15:04:05Z-0700")
}
