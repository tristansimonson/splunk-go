package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/brittonhayes/splunk-go/internal"
	splunk "github.com/brittonhayes/splunk-go/pkg"
	"github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	confirm bool
)

func init() {
	//Commands
	rootCmd.AddCommand(SystemCmd)
	SystemCmd.AddCommand(RestartCmd)
	SystemCmd.AddCommand(ControlsCmd)

	//Flags
	RestartCmd.Flags().BoolVarP(&confirm, "confirm", "c", false, "confirm you would like to restart")

	//Colors
	au = aurora.NewAurora(*colors)
}

// SystemCmd represents the system command
var SystemCmd = &cobra.Command{
	Use:   "system",
	Short: "Perform operations on the Splunk server.",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Help(cmd, args)
	},
}

// RestartCmd is used to restart the Splunk instance
var RestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the Splunk instance",
	Run: func(cmd *cobra.Command, args []string) {

		validate := func(input string) error {
			if len(input) < 1 {
				return errors.New("provide an input to confirm")
			}
			return nil
		}

		if confirm == false {
			confirmPrompt := promptui.Prompt{
				Label:    "Are you sure you'd like to restart Splunk? [Y/n]",
				Validate: validate,
				Default:  "Y",
			}

			result, err := confirmPrompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
			}

			if result == "Y" || result == "y" {
				msg := RestartSplunkInit()
				fmt.Println(msg)
			} else {
				os.Exit(0)
			}

		} else {
			msg := RestartSplunkInit()
			fmt.Println(msg)
		}

	},
}

// ControlsCmd is used to restart the Splunk instance
var ControlsCmd = &cobra.Command{
	Use:   "controls [name]",
	Short: "Lists actions that can be performed at this endpoint.",
	Long: `
	Function: Lists actions that can be performed at this endpoint. 
	Format: JSON
	Tip: Pipe into jq for prettified output`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.Help(cmd, args)
		fmt.Println(InspectControlInit(args[0]))
	},
}

// RestartSplunkInit is used to pass user auth into the RestartServer method from pkg
func RestartSplunkInit() string {
	conn := splunk.Connection{
		Username: viper.GetString("SPLUNK_USERNAME"),
		Password: viper.GetString("SPLUNK_PASSWORD"),
		BaseURL:  viper.GetString("SPLUNK_URL"),
	}
	response, err := conn.RestartServer()
	if err != nil {
		log.Fatal(au.Red("Couldn't restart Splunk. Ensure your configuration is correct."))
	}

	return response
}

// InspectControlInit is used to pass user auth into the InspectControl method from pkg
func InspectControlInit(endpoint string) string {
	conn := splunk.Connection{
		Username: viper.GetString("SPLUNK_USERNAME"),
		Password: viper.GetString("SPLUNK_PASSWORD"),
		BaseURL:  viper.GetString("SPLUNK_URL"),
	}
	response, err := conn.InspectControl(endpoint)
	if err != nil {
		log.Fatal(au.Red("Couldn't list actions for this endpoint. Check your input and try again."))
	}
	return response
}
