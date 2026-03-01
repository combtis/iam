// Copyright (c) 2026 Business Technologies
// SDPX-Licencse-Identifier: MIT
package iam

import (
	"net/http"
	"os"
)

var (
	site = os.Getenv("IAM_SITE")
)

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
