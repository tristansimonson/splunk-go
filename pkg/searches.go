package splunk

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
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

	dirPrompt := promptui.Prompt{
		Label:    "Where are your Splunk search files?",
		Validate: validate,
		Default:  viper.GetString("SEARCH_DIR"),
	}

	result, err := dirPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	realpath := AbsHome(result)

	return realpath
}

//ValidSPL checks if the filetype is a valid .spl file
func ValidSPL(path string) bool {
	valid := false
	if filepath.Ext(path) == ".spl" {
		valid = true
	} else {
		valid = false
	}

	return valid
}

//AbsHome converts all tildes to the current user's absolute home path
func AbsHome(path string) string {
	// Get the current user's home dir
	usr, _ := user.Current()
	dir := usr.HomeDir

	// Parse tilde symbols into prepended absolute home path
	if path == "~" {
		path = dir
	} else if strings.HasPrefix(path, "~/") {
		path = filepath.Join(dir, path[2:])
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("Home path conversion failed %v\n", err)
	}

	return abs
}
