package goshellshocked

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"

	"gopkg.in/yaml.v2"
)

// Write will write the provided commands. Depends on the value set in the output flag.
func Write(commands Commands) error {
	switch getOutput() {
	case "json":
		return writeToJSONFile(commands)
	case "yaml":
		return writeToYAMLFile(commands)
	default:
		return logToStdOut(commands)
	}
}

// logToStdOut will log the provided commands to stdout.
func logToStdOut(commands Commands) error {
	for _, c := range commands.GetData() {
		log.Printf("Frequency: %v, Command: %v", commands.GetFrequency(c), c)
	}

	return nil
}

// writeCmd is a representation of a printable command.
type writeCmd struct {
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

	dir, err := getOutputDir()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path.Join(dir, jsonFile), b, 0400)
}

func writeToYAMLFile(commands Commands) error {
	cl := getCommandList(commands)

	b, err := yaml.Marshal(&cl)
	if err != nil {
		return err
	}

	dir, err := getOutputDir()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path.Join(dir, yamlFile), b, 0600)
}

// getCommandList is a helper function to build a command slice
func getCommandList(commands Commands) []writeCmd {
	var cl []writeCmd
	for _, c := range commands.GetData() {
		cl = append(cl, writeCmd{
			Command:   c,
			Frequency: commands.GetFrequency(c),
		})
	}
	return cl
}
