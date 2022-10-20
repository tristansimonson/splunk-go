package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/tristansimonson/splunk-go/internal"
	splunk "github.com/tristansimonson/splunk-go/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	mode        string
	interactive bool
)

func init() {
	rootCmd.AddCommand(SearchCmd)
	SearchCmd.Flags().StringVarP(&mode, "mode", "m", "json", "The output mode of the search: csv, json, json_cols, json_rows, raw, xml")
	SearchCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Runs the search command in interactive mode.")
}

// SearchCmd is used to search splunk events
var SearchCmd = &cobra.Command{
	Use:   "search [string or filepath]",
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
		if interactive == false {
			internal.Help(cmd, args)

			if splunk.ValidSPL(args[0]) {
				content, err := ioutil.ReadFile(args[0])
				if err != nil {
					log.Fatal(err)
				}

				result := searchString(string(content), mode, conn)
				fmt.Println(result)
			} else {
				log.Fatal(au.Red("Please provide a valid .spl file."))
			}

		} else if interactive == true {
			file := splunk.SearchInteractive()
			if !splunk.ValidSPL(file) {
				log.Fatal(au.Red("Please provide a valid .spl file."))
			}
			content, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}
			result := searchString(string(content), mode, conn)
			fmt.Println(result)
		}

	},
}

func searchString(content string, mode string, conn splunk.Connection) string {
	input := string(content)

	response, err := conn.SearchSync(input, mode)
	if err != nil {
		log.Fatal(au.Red("Couldn't search Splunk. Ensure your configuration is correct."))
	}

	return response
}
