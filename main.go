package main

import (
	"fmt"
	"os"

	"github.com/brittonhayes/splunk-golang"
)

func main() {
	conn := splunk.Connection{
		Username: os.Getenv("SPLUNK_USERNAME"),
		Password: os.Getenv("SPLUNK_PASS"),
		BaseURL:  os.Getenv("SPLUNK_URL"),
	}

	key, err := conn.Login()
	if err != nil {
		fmt.Println("Couldn't login to splunk: ", err)
	}

	fmt.Println("Session key: ", key.Value)
}
