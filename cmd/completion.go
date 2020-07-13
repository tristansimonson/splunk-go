package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completionCmd)
}

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell completions",
	Long: `To load completions:

	Bash:
	
	$ source <(splunk-golang completion bash)
	
	# To load completions for each session, execute once:
	Linux:
		$ splunk-golang completion bash > /etc/bash_completion.d/splunk-golang
	MacOS:
		$ splunk-golang completion bash > /usr/local/etc/bash_completion.d/splunk-golang
	
	Zsh:
	
	$ source <(splunk-golang completion zsh)
	
	# To load completions for each session, execute once:
	$ splunk-golang completion zsh > "${fpath[1]}/_splunk-golang"
	`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "powershell":
			cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}
