// Package main is the entrypoint for the hello-go CLI.
package main

import (
	"fmt"
	"os"

	"github.com/John15321/hello-go/internal/cli"
)

func main() {
	if err := cli.NewRootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
