//go:build tools
// +build tools

package tools

import (
	_ "github.com/hexdigest/gowrap"
	_ "github.com/matryer/moq"
	_ "github.com/openfga/cli/cmd/fga"
)
