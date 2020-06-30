package cmd

import (
	"log"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func init() {
	rootCmd.AddCommand(GenDocs)
	au = aurora.NewAurora(*colors)
}

// GenDocs is used to generate documentation for the CLI
var GenDocs = &cobra.Command{
	Use:   "docs",
	Short: "Automatically generate documentation for the CLI.",
	Run: func(cmd *cobra.Command, args []string) {
		err := doc.GenMarkdownTree(rootCmd, "./docs")
		if err != nil {
			log.Fatal(err)
		}
	},
}
