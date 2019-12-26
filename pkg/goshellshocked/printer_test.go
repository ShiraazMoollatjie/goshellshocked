package goshellshocked

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

func setOutputForTesting(t *testing.T, nOutput string) func() {
	return func() {
		old := output
		*output = nOutput
		defer func() {
			*output = *old
		}()
	}
}

func setOutputDirForTesting(t *testing.T, nOutputDir string) func() {
	return func() {
		old := outputDir
		*outputDir = nOutputDir
		defer func() {
			*outputDir = *old
		}()
	}
}

var testCommand = Commands{
	data: []string{
		"ls",
		"git status",
	},
	frequencies: map[string]int{
		"ls":         3,
		"git status": 2,
	},
}

var result = []printCmd{
	{
		Command:   "ls",
		Frequency: 3,
	},
	{
		Command:   "git status",
		Frequency: 2,
	},
}

func TestLogToConsole(t *testing.T) {
	setOutputForTesting(t, "console")()

	// this is a really simple test, and I don't care about correctness here,
	// no error would mean that it works and this is okay.
	err := logToStdOut(testCommand)
	if err != nil {
		t.Errorf("error when logging to console. Error: %+v", err)
	}
}

func TestLogToJSON(t *testing.T) {
	setOutputForTesting(t, "json")()

	tmpDir := os.TempDir()
	f := path.Join(tmpDir, jsonFile)
	setOutputDirForTesting(t, tmpDir)()
	defer func() { os.Remove(f) }()

	err := writeToJSONFile(testCommand)
	if err != nil {
		t.Errorf("error when writing to yaml file Error: %+v", err)
	}

	_, err = os.Stat(f)
	if err != nil {
		t.Errorf("error finding json file Error: %+v", err)
	}

	b, err := ioutil.ReadFile(f)
	if err != nil {
		t.Errorf("error finding json file Error: %+v", err)
	}

	var res []printCmd
	json.Unmarshal(b, &res)
	if !reflect.DeepEqual(res, result) {
		t.Errorf("error comparing content. \nExpected:\t%+v\nActual:\t%+v", res, result)
	}
}

func TestLogToYAML(t *testing.T) {
	setOutputForTesting(t, "yaml")()

	tmpDir := os.TempDir()
	f := path.Join(tmpDir, yamlFile)
	setOutputDirForTesting(t, tmpDir)()
	defer func() { os.Remove(f) }()

	err := writeToYAMLFile(testCommand)
	if err != nil {
		t.Errorf("error when writing to yaml file Error: %+v", err)
	}

	_, err = os.Stat(f)
	if err != nil {
		t.Errorf("error finding yaml file Error: %+v", err)
	}

	b, err := ioutil.ReadFile(f)
	if err != nil {
		t.Errorf("error finding yaml file Error: %+v", err)
	}

	var res []printCmd
	yaml.Unmarshal(b, &res)
	if !reflect.DeepEqual(res, result) {
		t.Errorf("error comparing content. \nExpected:\t%+v\nActual:\t%+v", res, result)
	}
}
