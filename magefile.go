// +build mage

package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

const versionFileName = "cmd/version.go"

// Default target
// If not set, running mage will list available targets
//var Default = Build
var binary = "podcoff"
var packageName = "github.com/Strubbl/" + binary

func tidy() error {
	fmt.Println("+ mod tidy")
	return sh.Run("go", "mod", "tidy")
}

func verify() error {
	fmt.Println("+ mod verify")
	return sh.Run("go", "mod", "verify")
}

func UpdateVendor() error {
	mg.Deps(tidy)
	mg.Deps(verify)
	fmt.Println("+ updatevendor")
	return nil
}

func getLint() error {
	return sh.Run("go", "get", "-u", "golang.org/x/lint/golint")
}

func Lint() error {
	mg.Deps(getLint)
	fmt.Println("+ lint")
	if err := sh.RunV("golint", "-set_exit_status", "./..."); err != nil {
		return fmt.Errorf("error running lint: %v", err)
	}
	return nil
}

func Test() error {
	fmt.Println("+ test")
	return sh.Run("go", "test", "./...")
}

func Vet() error {
	fmt.Println("+ vet")
	if err := sh.RunV("go", "vet", "./..."); err != nil {
		return fmt.Errorf("error running vet: %v", err)
	}
	return nil
}

func Fmt() error {
	fmt.Println("+ fmt")
	failed := false
	files, err := filepath.Glob(filepath.Join("*.go"))
	if err != nil {
		fmt.Printf("error globbing file for gofmt: %v\n", err)
		return err
	}
	for _, f := range files {
		s, err := sh.Output("gofmt", "-l", f)
		if err != nil {
			fmt.Printf("error running gofmt: %v\n", err)
			failed = true
		}
		if s != "" {
			fmt.Println("The following files are not valid go format:")
			failed = true
			fmt.Println(s)
		}
	}
	if failed {
		return errors.New("improperly formatted go files")
	}
	return nil
}

func Build() error {
	mg.Deps(tidy)
	mg.Deps(verify)
	out, err := sh.Output("git", "describe", "--tags", "--dirty")
	fmt.Println("+ build")
	if err != nil {
		return err
	}
	sedExpr := strings.Join([]string{"s/const programVersion = \"master\"/var programVersion = \"", out, "\"/g"}, "")
	err = sh.Run("sed", "-i", sedExpr, versionFileName)
	if err != nil {
		return err
	}
	fmt.Println(out)
	err = sh.Run("go", "build", "-o", binary, "cmd/"+binary+"/main.go")
	if err != nil {
		return err
	}
	err = sh.Run("git", "checkout", versionFileName)
	if err != nil {
		return err
	}
	return err
}

func Clean() {
	fmt.Println("+ clean")
	os.RemoveAll(binary)
}

func Check() {
	fmt.Println("+ check")
	mg.Deps(Fmt)
	mg.Deps(Vet)
	mg.Deps(Lint)
	mg.Deps(Test)
}

// clean, build
func All() {
	// clean build check (fmt lint test vet)
	fmt.Println("+ all")
	mg.Deps(Clean)
	mg.Deps(Build)
	mg.Deps(Check)
}
