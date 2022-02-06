package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line, err := json2test(scanner.Bytes())
		if err != nil {
			panic(err)
		}
		if line != `` {
			fmt.Printf("%s", line)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// Function json2test is a sole kernel of the whole solution.
// I assume that `go tool test2json` produces JSON objects in form
// described in https://github.com/golang/go/blob/master/src/cmd/internal/test2json/test2json.go
func json2test(in []byte) (string, error) {

	var parsed struct {
		Time    *time.Time `json:",omitempty"`
		Action  string
		Package string   `json:",omitempty"`
		Test    string   `json:",omitempty"`
		Elapsed *float64 `json:",omitempty"`
		Output  string   `json:",omitempty"`
	}

	err := json.Unmarshal(in, &parsed)
	if err != nil {
		return ``, err
	}
	if parsed.Action == "output" {
		return parsed.Output, nil
	}
	return ``, nil
}
