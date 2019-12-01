// +build mage

package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"

	// mage:import
	grimoire "github.com/VixsTy/grimoire"
	figure "github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

var Default = Build

var curDir = func() string {
	name, _ := os.Getwd()
	return name
}()

// Calculate file paths
var toolsBinDir = grimoire.NormalizePath(path.Join(curDir, "tools", "bin"))

func init() {
	time.Local = time.UTC

	// Add local bin in PATH
	err := os.Setenv("PATH", fmt.Sprintf("%s:%s", toolsBinDir, os.Getenv("PATH")))
	if err != nil {
		panic(err)
	}
}

func Build() {
	banner := figure.NewFigure(grimoire.MainDirectoryName(), "", true)
	banner.Print()

	fmt.Println("")
	color.Red("# Build Info ---------------------------------------------------------------")
	fmt.Printf("Go version : %s\n", runtime.Version())
	fmt.Printf("Git revision : %s\n", grimoire.Hash())
	fmt.Printf("Git branch : %s\n", grimoire.Branch())
	fmt.Printf("Tag : %s\n", grimoire.Tag())

	fmt.Println("")

	color.Red("# Core packages ------------------------------------------------------------")
	mg.SerialDeps(grimoire.Go.Deps, grimoire.Go.License, grimoire.Go.Format, grimoire.Go.Lint, Go.Test)

	fmt.Println("")
	color.Red("# Artifacts ----------------------------------------------------------------")
	mg.Deps(Bin.Calculator)
}

// Bob is a mage namespace which will manage binaries build actions
type Bin mg.Namespace

// Calculator build the Calculator binarie
func (Bin) Calculator() error {
	return grimoire.Build{}.Binary("github.com/VixsTy/calculator/cli/calculator", "calculator")
}

// Go is a mage namespace which will manage golang actions
type Go mg.Namespace

// Test run all test
func (Go) Test() error {
	color.Cyan("## Running unit tests")
	sh.Run("mkdir", "-p", "test-results/junit")
	return sh.RunV("gotestsum", "--junitfile", "test-results/junit/unit-tests.xml", "--", "-short", "-race", "-cover", "./...")
}
