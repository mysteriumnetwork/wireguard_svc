//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
	"path"
)

// Runs go mod download and then installs the binary.
func Build() error {
	env := map[string]string{
		"GOOS":   "windows",
		"GOARCH": "amd64",
	}
	path.Join("cmd", "main.go")
	return sh.RunWith(env, "go", "build", "-o", "build/wireguard_svc.exe", "cmd/main.go")
}
