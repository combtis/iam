// Copyright (c) 2026 Business Technologies
// SDPX-License-Identifier: MIT
package iam

import (
	"testing"
)

func Test(t *testing.T) {
	t.Log("iam testing...\n")

	jwks := NewJWKS()
	pub1, priv1, err := jwks.GenerateKey()
	if err == nil {
		//jwks.Load("testdata/private")
		_ = pub1
		jwks.InsertPrivateKey(priv1)
		t.Log(jwks.String())
	}
	pub2, priv2, err := jwks.GenerateKey()
	if err == nil {
		//jwks.Load("testdata/private")
		_ = pub2
		jwks.InsertPrivateKey(priv2)
		t.Log(jwks.String())
	}
}
