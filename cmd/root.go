package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/spf13/cobra"
)

var files []string
var asString bool

var rootCmd = &cobra.Command{
	Use:   "tftarget",
	Short: "Lists Terraform targets from files",
	Run: func(cmd *cobra.Command, args []string) {
		if len(files) == 0 {
			log.Fatal("Please provide Terraform files with -f flag")
		}

		targetSet := make(map[string]struct{})
		for _, filename := range files {
			targets, err := parseTerraformFile(filename)
			if err != nil {
				log.Printf("Failed parsing %s: %v\n", filename, err)
				continue
			}
			for _, target := range targets {
				targetSet[target] = struct{}{}
			}
		}

		for t := range targetSet {
			if asString {
				fmt.Printf("-target=%s ", t)
			} else {
				fmt.Println(t)
			}
		}

		if asString {
			fmt.Println() // just to ensure newline after targets list
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringSliceVarP(&files, "file", "f", []string{}, "Terraform files to parse (multiple allowed)")
	rootCmd.Flags().BoolVarP(&asString, "string", "s", false, "Output as -target=resource strings")
}

func parseTerraformFile(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	file, diags := hclsyntax.ParseConfig(content, filename, hcl.InitialPos)
	if diags.HasErrors() {
		return nil, fmt.Errorf(diags.Error())
	}

	syntaxBody, ok := file.Body.(*hclsyntax.Body)
	if !ok {
		return nil, fmt.Errorf("error casting file body")
	}

	var targets []string
	for _, block := range syntaxBody.Blocks {
		if block.Type == "resource" && len(block.Labels) == 2 {
			resourceType := block.Labels[0]
			resourceName := block.Labels[1]
			target := fmt.Sprintf("%s.%s", resourceType, resourceName)
			targets = append(targets, target)
		}
	}

	return targets, nil
}
