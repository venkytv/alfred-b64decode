package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

type AlfredResultItem struct {
	Uid          string `json:"uid"`
	Arg          string `json:"arg"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	Icon         string `json:"icon"`
	Valid        bool   `json:"valid"`
	AutoComplete string `json:"autocomplete"`
}

type AlfredResult struct {
	Items []AlfredResultItem `json:"items"`
}

func main() {
	// If no args are passed, print usage
	if len(os.Args) < 2 {
		fmt.Println("Usage: alfred-workflow-go <b64-encoded-string>")
		return
	}

	subtitle := ""
	valid := true

	// Decode the passed string
	decoded, err := base64.StdEncoding.DecodeString(os.Args[1])
	if err != nil {
		subtitle = err.Error()
		valid = false
	}

	// Decode byte array to string
	decoded_str := string(decoded)

	// Create an AlfredResultItem and marshal it to JSON
	i := AlfredResultItem{
		Arg:      decoded_str,
		Title:    fmt.Sprintf("\"%s\"", decoded_str),
		Subtitle: subtitle,
		Valid:    valid,
	}

	jsonData, err := json.Marshal(AlfredResult{Items: []AlfredResultItem{i}})
	if err != nil {
		// Abort with error
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Printf("%s\n", jsonData)
}
