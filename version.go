// Package version records versioning information about this module.
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	major      = 0
	minor      = 1
	patch      = 2
	preRelease = "devel"
)

func showVersion() {
	version := fmt.Sprintf("v%d.%d.%d", major, minor, patch)
	if preRelease != "" {
		version += "-" + "devel"
	}

	fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), version)
}
