package main

import (
	"fmt"
	"os"

	"github.com/raihanul-2k15/zatca-qr-reader/cli"
)

var VERSION string = "current" // to be replaced by -ldflags="-X main.VERSION=xyz"

func main() {
	versionArgs := []string{"-v", "--version", "version"}
	if len(os.Args) > 1 && contains(versionArgs, os.Args[1]) {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	cli.Run()
}

func contains(args []string, arg string) bool {
	for _, a := range args {
		if a == arg {
			return true
		}
	}
	return false
}
