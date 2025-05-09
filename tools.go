//go:build tools
// +build tools

package tools

import (
	_ "github.com/hexdigest/gowrap"
	_ "github.com/lxc/incus/v6/cmd/generate-database"
	_ "github.com/matryer/moq"
	_ "github.com/openfga/cli/cmd/fga"
	_ "golang.org/x/tools/cmd/goimports"
)
