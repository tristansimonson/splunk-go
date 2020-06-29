# Splunk REST API Client

A Splunk REST API Client written in Golang.

## Table of Contents :notebook:

- [Installation](#installation)
- [Usage](#usage)
- [Support](#support)
- [Contributing](#contributing)

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
    "github.com/brittonhayes/splunk-golang"
)

func main() {
    conn := splunk.Connection {
            Username: os.Getenv("SPLUNK_USERNAME"),
            Password: os.Getenv("SPLUNK_PASS"),
            BaseURL: os.Getenv("SPLUNK_URL"),
    }

    key, err:= conn.Login()
    if err != nil {
            fmt.Println("Couldn't login to splunk: ", err)
    }

    fmt.Println("Session key: ", key.Value)
}
```

### Using the package's CLI

Once the binary is compiled and added to your path, you can utilize the CLI to perform Splunk operations.

```shell
splunk-golang --help
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
