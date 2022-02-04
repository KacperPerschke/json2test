package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
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

func json2test(in []byte) (string, error) {

	var parsed struct {
		Time    string
		Action  string
		Package string
		Test    string
		Output  string
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
