//go:build tools
// +build tools

package main

import (
	_ "github.com/BurntSushi/toml/cmd/tomlv"
	_ "github.com/fzipp/gocyclo"
	_ "github.com/go-critic/go-critic"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "golang.org/x/tools/cmd/goimports"
)
