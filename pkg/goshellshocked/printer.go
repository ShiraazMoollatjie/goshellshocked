package goshellshocked

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// PrintOption defines the different options of printing output.
type PrintOption int

const (
	printOptionUnknown PrintOption = iota

	// PrintOptionStdout writes to stdout.
	PrintOptionStdout

	// PrintOptionJSON writes to a json file.
	PrintOptionJSON

	// PrintOptionYAML writes to a json file.
	PrintOptionYAML
)

// Write will write the provided commands to the provided PrintOption mode.
func Write(commands Commands) error {
	mode := GetPrintOption()
	switch mode {
	case PrintOptionStdout:
		return logToStdOut(commands)
	case PrintOptionJSON:
		return writeToJSONFile(commands)
	case PrintOptionYAML:
		return writeToYAMLFile(commands)
	default:
		return errors.New("unsupported printoption")
	}
}

// logToStdOut will log the provided commands to stdout.
func logToStdOut(commands Commands) error {
	for _, c := range commands.GetData() {
		log.Printf("Frequency: %v, Command: %v", commands.GetFrequency(c), c)
	}

	return nil
}

type command struct {
	Command   string
	Frequency int
}

const (
	jsonFile = "shellshocked.json"
	yamlFile = "shellshocked.yml"
)

func writeToJSONFile(commands Commands) error {
	cl := getCommandList(commands)

	b, err := json.MarshalIndent(&cl, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(jsonFile, b, 0400)
}

func writeToYAMLFile(commands Commands) error {
	cl := getCommandList(commands)

	b, err := yaml.Marshal(&cl)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(yamlFile, b, 0600)
}

// getCommandList is a helper function to build a command slice
func getCommandList(commands Commands) []command {
	var cl []command
	for _, c := range commands.GetData() {
		cl = append(cl, command{
			Command:   c,
			Frequency: commands.GetFrequency(c),
		})
	}
	return cl
}
