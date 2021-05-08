package main

import (
	"github.com/magiconair/properties/assert"
	"mycli/cmd"
	"testing"
)

func TestCliCommand(t *testing.T) {
	command := cmd.GetRootCmd()
	command.SetArgs([]string{"hello", "even", "1", "2", "3"})
	exitCode := command.Execute()

	assert.Equal(t, exitCode, nil)

}
