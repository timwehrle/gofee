package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/timwehrle/gofee/pkg/gofee"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	// The default length of the password
	defaultLength int = 16

	MAJOR = 0
	MINOR = 1
	PATCH = 0
)

// Options for the generate command
var options struct {
	length       int
	lowers       bool
	uppers       bool
	digits       bool
	symbols      bool
	passwordType string
}

func init() {
	rootCmd.Flags().BoolVarP(&options.lowers, "exclude-lowers", "w", false, "exclude lowercase letters")
	rootCmd.Flags().BoolVarP(&options.uppers, "exclude-uppers", "u", false, "exclude uppercase letters")
	rootCmd.Flags().BoolVarP(&options.digits, "exclude-digits", "d", false, "exclude digits")
	rootCmd.Flags().BoolVarP(&options.symbols, "exclude-symbols", "s", false, "exclude symbols")
	rootCmd.Flags().IntVarP(&options.length, "length", "l", defaultLength, "length of the password")
	rootCmd.Flags().StringVarP(&options.passwordType, "type", "t", "", "type of password to generate (pin, memorable)")

	// Colorize the usage output
	rootCmd.SetOutput(color.Output)
	cobra.AddTemplateFunc("StyleHeading", color.New(color.FgGreen).SprintFunc())
	usageTemplate := rootCmd.UsageTemplate()
	usageTemplate = strings.NewReplacer(
		`Usage:`, `{{StyleHeading "Usage:"}}`,
		`Aliases:`, `{{StyleHeading "Aliases:"}}`,
		`Examples:`, `{{StyleHeading "Examples:"}}`,
		`Available Commands:`, `{{StyleHeading "Available Commands:"}}`,
		`Flags:`, `{{StyleHeading "Flags:"}}`,
	).Replace(usageTemplate)
	re := regexp.MustCompile(`(?m)^Flags:\s*$`)
	usageTemplate = re.ReplaceAllLiteralString(usageTemplate, `{{StyleHeading "Flags:"}}`)
	rootCmd.SetUsageTemplate(usageTemplate)
}

var example = `
gofee --length 20 --exclude-lowers
gofee --length 12 -u -d 
gofee --type pin --length 4
`

var long = `
Gofee is a simple password generator, which relies on the cryptographic strength of the system's random number generator.
It generates a password of a given length, using a set of characters that can be customized by the user.
`

var rootCmd = &cobra.Command{
	Use:     "gofee",
	Version: fmt.Sprintf("%d.%d.%d", MAJOR, MINOR, PATCH),
	Example: example,
	Short:   "Gofee is a simple password generator, which is reliable and secure.",
	Long:    long,
	Run: func(cmd *cobra.Command, args []string) {
		config := gofee.PasswordConfig{
			IncludeLowers:  !options.lowers,
			IncludeUppers:  !options.uppers,
			IncludeDigits:  !options.digits,
			IncludeSymbols: !options.symbols,
			Type:           options.passwordType,
		}

		pw, err := gofee.Generate(options.length, config)
		if err != nil {
			log.Fatalf("Error generating password: %v", err)
		}

		entropy, err := gofee.CalculateEntropy(len(gofee.Charset), options.length)
		if err != nil {
			log.Fatalf("Error calculating entropy: %v", err)
		}

		fmt.Print("Entropy: ")
		color.Green("%.2f bits", entropy)

		fmt.Printf("Password: %s", color.GreenString(pw))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
