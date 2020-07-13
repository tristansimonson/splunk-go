package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	splunk "github.com/brittonhayes/splunk-golang/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var mode string
var interactive string

func init() {
	rootCmd.AddCommand(SearchCmd)
	SearchCmd.Flags().StringVarP(&mode, "mode", "m", "json", "The output mode of the search: csv, json, json_cols, json_rows, raw, xml")
	SearchCmd.Flags().StringVarP(&interactive, "interactive", "i", "false", "Runs the search command in interactive mode.")
}

// SearchCmd is used to search splunk events
var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search Splunk for events.",
	Long: `The search command is used to perform search queries via the Splunk REST API. 
Searching splunk using the CLI requires one argument of a SPL file containing your search.

e.g. splunk-go search ~/.splunk-go/searches/my-search.spl

`,
	Run: func(cmd *cobra.Command, args []string) {

		conn := splunk.Connection{
			Username: viper.GetString("SPLUNK_USERNAME"),
			Password: viper.GetString("SPLUNK_PASSWORD"),
			BaseURL:  viper.GetString("SPLUNK_URL"),
		}

		// TODO Refactor this so it's not duplicated
		if interactive == "true" {
			file := splunk.SearchInteractive()
			content, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}
			searchString := string(content)

			response, err := conn.SearchSync(searchString, mode)
			if err != nil {
				log.Fatal(au.Red("Couldn't search Splunk. Ensure your configuration is correct."))
			}

			fmt.Printf(response)
		} else {
			content, err := ioutil.ReadFile(args[0])
			if err != nil {
				log.Fatal(err)
			}
			searchString := string(content)

			response, err := conn.SearchSync(searchString, mode)
			if err != nil {
				log.Fatal(au.Red("Couldn't search Splunk. Ensure your configuration is correct."))
			}

			fmt.Printf(response)
		}
	},
}
