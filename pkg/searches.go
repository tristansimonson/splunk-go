package splunk

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

// SearchSync performs a search job on splunk with the provided search string
func (conn Connection) SearchSync(searchString string, outputMode string) (string, error) {
	data := make(url.Values)
	data.Add("search", searchString)
	data.Add("max_time", "30")
	data.Add("earliest_time", "-10m")
	data.Add("latest_time", "now")
	data.Add("output_mode", outputMode)

	response, err := conn.httpPost(fmt.Sprintf("%s/services/search/jobs/export", conn.BaseURL), &data)
	return response, err
}

// SearchInteractive runs the interactive variant of SearchSync
func SearchInteractive() string {

	var files []string

	root := GetSearchDir()
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), "spl") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	prompt := promptui.Select{
		Label: "Select search file",
		Items: files,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return result
}

// GetSearchDir runs a prompt for the use to check where their spl files are for Splunk searching
func GetSearchDir() string {

	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("Directory must have more than 3 characters")
		}
		return nil
	}

	d, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	dirPrompt := promptui.Prompt{
		Label:    "Where are your Splunk search files?",
		Validate: validate,
		Default:  d,
	}

	result, err := dirPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return result

}
