// Package cli implements the hello-go CLI commands.
package cli

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// Version is set at build time via -ldflags.
var Version = "dev"

// NewRootCmd builds the root cobra command.
func NewRootCmd() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:     "hello-go",
		Short:   "Prints out Hello World!",
		Version: Version,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return greet(cmd.OutOrStdout(), name)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "World", "name to greet")
	return cmd
}

func greet(w io.Writer, name string) error {
	_, err := fmt.Fprintf(w, "Hello, %s!\n", name)
	return err
}
