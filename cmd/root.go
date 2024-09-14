package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gofe",
	Version: "0.0.1",
	Short:   "gofe is a simple password generator, which is reliable and secure.",
	Long:    "gofe is a simple password generator, which is reliable and secure.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
