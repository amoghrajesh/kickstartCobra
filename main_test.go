package main

import (
	"github.com/magiconair/properties/assert"
	"mycli/cmd"
	"testing"
)

func TestCliCommand(t *testing.T) {
	command := cmd.GetRootCmd()
	command.SetArgs([]string{"hello", "1", "2"})
	exitCode := command.Execute()

	assert.Equal(t, exitCode, nil)

}
