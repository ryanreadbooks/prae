package main

import (
	_ "embed"

	"github.com/ryanreadbooks/prae/cmd"
)

//go:embed VERSION
var version string

func main() {
	cmd.Execute(version)
}
