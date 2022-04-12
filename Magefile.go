//+build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/sh"
	"strings"
)

// Runs go mod download and then installs the binary.
func Build() error {
	var err error

	cmds := []string{
		"telegraf-sparkplugb",
	}

	arches := []string{
		//"linux/386",
		"linux/amd64",
		//"linux/arm64",
		//"windows/386",
		//"windows/amd64",
	}

	//
	// Get dependencies
	//
	if err = sh.Run("go", "mod", "download"); err != nil {
		return err
	}

	for _, command_package := range cmds {
		for _, arch := range arches {
			p := strings.Split(arch, "/")
			err = build(command_package, p[0], p[1])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func build(pkg string, os string, arch string) error {
	source := fmt.Sprintf("./cmd/%s", pkg)
	dest := fmt.Sprintf("%s-%s-%s", pkg, os, arch)
	if os == "windows" {
		dest = dest + ".exe"
	}

	return sh.RunWith(map[string]string{"GOOS": os, "GOARCH": arch}, "go", "build", "-v", "-o", dest, source)
}

