package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gofee",
	Version: "0.0.1",
	Short:   "Gofee is a simple password generator, which is reliable and secure.",
	Long:    "Gofee is a simple password generator, which is reliable and secure.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
