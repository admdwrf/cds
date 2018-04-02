package main

import (
	"io/ioutil"

	"github.com/ovh/cds/sdk/plugin"
)

//HelloPlugin is a plugin to write Hello World into a text file
type HelloPlugin struct {
	plugin.Common
}

//Name return plugin name. It must me the same as the binary file
func (d HelloPlugin) Name() string {
	return "plugin-hello"
}

//Description explains the purpose of the plugin
func (d HelloPlugin) Description() string {
	return "This is a plugin to write Hello World into a file"
}

//Author of the plugin
func (d HelloPlugin) Author() string {
	return "Frederic ALIX <frederic.alix@fredalix.com>"
}

// Parameters return parameters description
func (d HelloPlugin) Parameters() plugin.Parameters {
	params := plugin.NewParameters()
	params.Add("filepath", plugin.StringParameter, "the name and destination of your hello file", ".")

	return params
}

// Run execute the action
func (d HelloPlugin) Run(a plugin.IJob) plugin.Result {
	filepath := a.Arguments().Get("filepath")

	// Create the file
	file := []byte("Hello World\n")
	err := ioutil.WriteFile(filepath, file, 0644)

	if err != nil {
		plugin.SendLog(a, "Error to create the file %s : %s", filepath, err)
		return plugin.Fail
	}

	return plugin.Success
}

func main() {
	plugin.Main(&HelloPlugin{})
}
