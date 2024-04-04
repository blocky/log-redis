//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Build() {
	sh.Run("go", "build", "-o", "./bin/service", ".")
}
