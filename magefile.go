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
	fmt.Println("+ updatevendor")
	err := sh.Run("go", "get", "-u", "./...")
	if err != nil {
		return err
	}
	err = tidy()
	if err != nil {
		return err
	}
	return nil
}

func getLint() error {
	return sh.Run("go", "get", "-u", "golang.org/x/lint/golint")
}

func Lint() error {
	// we have to go mod tidy after linting to clean up go.mod and go.sum
	defer tidy()
	mg.Deps(getLint)
	fmt.Println("+ lint")
	if err := sh.RunV("golint", "./..."); err != nil {
		return fmt.Errorf("error running lint: %v", err)
	}
	return nil
}

func Test() error {
	fmt.Println("+ test")
	return sh.RunV("go", "test", "./...")
}

func TestAll() error {
	fmt.Println("+ testAll")
	return sh.Run("go", "test", "all")
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
	err = regularBuild()
	if err != nil {
		return err
	}
	fmt.Println(out)
	err = sh.Run("git", "checkout", versionFileName)
	if err != nil {
		return err
	}
	return err
}

func regularBuild() error {
	return sh.Run("go", "build", "-o", binary, "cmd/"+binary+"/main.go")
}

// only used in travis
func CiBuild() error {
	mg.Deps(tidy)
	mg.Deps(verify)
	return regularBuild()
}

// only used in travis
func CiInstall() error {
	mg.Deps(CiBuild)
	fmt.Println("+ install")
	return sh.Run("go", "install", ".")
}

func Install() error {
	mg.Deps(Build)
	fmt.Println("+ install")
	return sh.Run("go", "install", ".")
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

// clean build check (fmt lint test vet)
func All() {
	fmt.Println("+ all")
	mg.Deps(Clean)
	mg.Deps(Build)
	mg.Deps(Check)
}
