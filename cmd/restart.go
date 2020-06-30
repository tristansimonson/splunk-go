package cmd

import (
	"fmt"
	"log"

	splunk "github.com/brittonhayes/splunk-golang/pkg"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(RestartCmd)
	au = aurora.NewAurora(*colors)
}

// RestartCmd is used to restart the Splunk instance
var RestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the Splunk instance",
	Run: func(cmd *cobra.Command, args []string) {

		conn := splunk.Connection{
			Username: viper.GetString("SPLUNK_USERNAME"),
			Password: viper.GetString("SPLUNK_PASSWORD"),
			BaseURL:  viper.GetString("SPLUNK_URL"),
		}

		response, err := conn.RestartServer()
		if err != nil {
			log.Fatal(au.Red("Couldn't restart Splunk. Ensure your configuration is correct."))
		}

		fmt.Printf(response)
	},
}
