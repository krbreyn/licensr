package main

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

//go:embed licenses/*.txt
var licenseFiles embed.FS

var rootCmd = &cobra.Command{
	Use:   "licensr",
	Short: "Create license files from the CLI",
	Long: `Licensr is a tool to list, print, or create license files.

Usage:
	licensr list
	licensr print <license-name>
	licensr make <icense-name>`,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available licenses",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		files, err := licenseFiles.ReadDir("licenses")
		if err != nil {
			return fmt.Errorf("error reading licenses: %w", err)
		}

		for _, file := range files {
			name := strings.TrimSuffix(file.Name(), ".txt")
			cmd.Println(name)
		}
		return nil
	},
}

var printCmd = &cobra.Command{
	Use:   "print <license-name>",
	Short: "Print a specific license",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		want := strings.ToLower(args[0])

		files, err := licenseFiles.ReadDir("licenses")
		if err != nil {
			return fmt.Errorf("error reading licenses: %w", err)
		}

		for _, file := range files {
			name := strings.ToLower(strings.TrimSuffix(file.Name(), ".txt"))
			if name == want {
				contents, err := licenseFiles.ReadFile("licenses/" + file.Name())
				if err != nil {
					return fmt.Errorf("error reading license file: %w", err)
				}
				cmd.Println(string(contents))
				return nil
			}
		}

		return fmt.Errorf("no license with name %s found", args[0])
	},
}

var makeCmd = &cobra.Command{
	Use:   "make <license-name>",
	Short: "Create a LICENSE file with the specified license",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		want := strings.ToLower(args[0])

		files, err := licenseFiles.ReadDir("licenses")
		if err != nil {
			return fmt.Errorf("error reading licenses: %w", err)
		}

		var licenseTxt string
		var licenseName string
		for _, file := range files {
			name := strings.TrimSuffix(file.Name(), ".txt")
			if strings.ToLower(name) == want {
				licenseName = name

				contents, err := licenseFiles.ReadFile("licenses/" + file.Name())
				if err != nil {
					return fmt.Errorf("error reading license file: %w")
				}

				licenseTxt = string(contents)
				break
			}
		}

		if licenseTxt == "" {
			return fmt.Errorf("no license with name %s found", args[0])
		}

		dst, err := os.Create("LICENSE")
		if err != nil {
			return fmt.Errorf("error creating LICENSE file: %w", err)
		}
		defer dst.Close()

		_, err = dst.WriteString(licenseTxt)
		if err != nil {
			return fmt.Errorf("error writing to LICENSE file: %w", err)
		}

		cmd.Printf("Success! You rae now licensed under %s.\n", licenseName)
		return nil
	},
}

func main() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(printCmd)
	rootCmd.AddCommand(makeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
