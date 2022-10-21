package util

import (
	"errors"
	"strings"

	"golang.org/x/exp/slices"
)

var validCommands = []string{
	"read", "write", "delete", "start", "commit", "abort", "quit",
}

type CommandInput struct {
	Command string
	Key     string
	Value   string
}

// Copy values from inputDb to targetDb. Potential overwrites for same keys here.
func PopulateDb(inputDb map[string]string, targetDb map[string]string) {
	if inputDb == nil || targetDb == nil {
		return
	}
	for k, _ := range inputDb {
		targetDb[k] = inputDb[k]
	}
}

// Simple validation, trim whitespace and end chars for input command and validate it's one of our "valid" inputs.
func ValidateCommand(command string) bool {
	command = strings.TrimSpace(command)
	command = strings.Trim(command, "\r")
	command = strings.Trim(command, "\n")
	return slices.Contains(validCommands, strings.ToLower(command))
}

// validate input string, assert len of 3 and return any errors
func ValidateInput(input string) (*CommandInput, error) {
	var res = strings.Fields(input)
	if len(res) > 3 {
		return nil, errors.New("Command is not valid. Please enter a valid command.")
	}
	if len(res) < 1 {
		return nil, errors.New("Command is not valid. Please enter a valid command.")
	}
	if !ValidateCommand(res[0]) {
		return nil, errors.New("Command is not valid. Please enter a valid command.")
	}
	if strings.ToLower(res[0]) == "write" && len(res) != 3 {
		return nil, errors.New("Command is not valid. Please enter a valid command.")
	}
	if (strings.ToLower(res[0]) == "read" || strings.ToLower(res[0]) == "delete") && len(res) != 2 {
		return nil, errors.New("Command is not valid. Please enter a valid command.")
	}
	if strings.ToLower(res[0]) == "write" {
		return &CommandInput{Command: res[0], Key: res[1], Value: res[2]}, nil
	} else if strings.ToLower(res[0]) == "read" || strings.ToLower(res[0]) == "delete" {
		return &CommandInput{Command: res[0], Key: res[1]}, nil
	}
	return &CommandInput{Command: res[0]}, nil
}
