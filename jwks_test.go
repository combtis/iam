// Copyright (c) 2026 Business Technologies
// SDPX-License-Identifier: MIT
package iam

import (
	"testing"
)

func Test(t *testing.T) {
	t.Log("iam testing...\n")

	jwks := NewJWKS()
	pub, priv, err := jwks.GenerateKey()
	if err == nil {
		//jwks.Load("testdata/private")
		_ = pub
		jwks.InsertPrivateKey(priv)
		t.Log(jwks.String())
	}
}
