# Splunk REST API Client

A Splunk REST API Client written in Golang.

## Table of Contents :notebook:

- [Installation](#installation)
- [Usage](#usage)
- [Support](#support)
- [Contributing](#contributing)

![Go](https://github.com/brittonhayes/splunk-golang/workflows/Go/badge.svg)

## Installation :cloud:

1. Install the package from Github

```
git clone https://github.com/brittonhayes/splunk-golang
```

## Usage :hammer:

1. Navigate to the project directory

2. Run the entry file

```bash
go run main.go
```

3. Use the Splunk REST API CLI!

## Examples

### Using the package to run without intervention

```go
package main

import (
    "fmt"
    "os"
    "github.com/logrusorgru/aurora"
    "github.com/brittonhayes/splunk-golang"
    "github.com/spf13/viper"
)

func init() {
	au = aurora.NewAurora(*colors)
}

func main() {
    conn := splunk.Connection{
        Username: viper.GetString("SPLUNK_USERNAME"),
        Password: viper.GetString("SPLUNK_PASSWORD"),
        BaseURL:  viper.GetString("SPLUNK_URL"),
    }

    key, err := conn.Login()
    if err != nil {
        log.Fatal(au.Red("Couldn't login to Splunk. Ensure your configuration is correct."))
    }

    fmt.Println("Logged in Successfully. \nSession key:", au.BrightBlue(key.Value))
}
```

### Using the package's CLI

Once the binary is compiled and added to your path, you can utilize the CLI to perform Splunk operations.

1. Build the binary

```shell
go build -o ./bin/splunk-go .

```

2. Run the help command to get a list of possible actions

```shell
splunk-go --help
```

This will output a message like the following:

```
A Splunk REST API client written in GO.

Usage:
  splunk-go [command]

Available Commands:
  help        Help about any command
  login       Authenticate to Splunk and return a session token.
  restart     Restart the Splunk instance
  search      Search Splunk for events.

Flags:
      --config string   config file (default is $HOME/.splunk-go.yaml)
  -h, --help            help for splunk-go

Use "splunk-go [command] --help" for more information about a command.

```

### Adding features to the package

1. Navigate to the `pkg` directory of the repository and add a new \*.go file.

2. Import your function or type from the new file into the `main.go` if you want it to run without interventon.

3. Import your function or type from the new file into the `cmd` directory if you want to add it to the CLI.

---

## Support

Please [open an issue](https://github.com/brittonhayes/splunk-golang/issues/new) for support.

## Contributing

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/). Create a branch, add commits, and [open a pull request](https://github.com/brittonhayes/splunk-golang/compare/).

## Acknowledgements

Social image from [Ashley Mcnamara](https://github.com/ashleymcnamara/gophers)
