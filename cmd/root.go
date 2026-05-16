package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mostafizemon/goku-cli.git/internal/converter"
	"github.com/spf13/cobra"
)

var (
	inputFile  string
	outputFmt  string
)

var rootCmd = &cobra.Command{
	Use:   "goku",
	Short: "Goku CLI - A powerful file conversion tool",
	Long:  `Goku CLI converts configuration files between JSON and YAML formats.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if inputFile == "" {
			return fmt.Errorf("input file is required. Use -i <file_path>")
		}
		if outputFmt == "" {
			return fmt.Errorf("output format is required. Use -o <json|yaml>")
		}

		// Normalize
		outputFmt = strings.ToLower(strings.TrimPrefix(outputFmt, "."))
		if outputFmt != "json" && outputFmt != "yaml" && outputFmt != "yml" {
			return fmt.Errorf("invalid output format %q. Must be json or yaml", outputFmt)
		}

		// Detect input format from extension
		inputFmt := converter.DetectFormat(inputFile)
		if inputFmt == "" {
			return fmt.Errorf("cannot detect input format. File must be .json, .yaml, or .yml")
		}

		// Normalize yml -> yaml for comparison
		normalizedInput  := strings.ReplaceAll(inputFmt, "yml", "yaml")
		normalizedOutput := strings.ReplaceAll(outputFmt, "yml", "yaml")

		// Edge case: same format
		if normalizedInput == normalizedOutput {
			return fmt.Errorf("requested output must be in a different format than the input file (both are %s)", inputFmt)
		}

		// Read input file
		data, err := os.ReadFile(inputFile)
		if err != nil {
			return fmt.Errorf("failed to read input file: %w", err)
		}

		// Convert
		result, err := converter.Convert(data, inputFmt, outputFmt)
		if err != nil {
			return fmt.Errorf("conversion failed: %w", err)
		}

		fmt.Println(string(result))
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file path (.json or .yaml/.yml)")
	rootCmd.Flags().StringVarP(&outputFmt, "output", "o", "", "Output format: json or yaml")
}