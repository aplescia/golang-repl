package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopulate(t *testing.T) {
	var mapOne = make(map[string]string)
	var mapTwo = make(map[string]string)
	mapOne["yes"] = "no"
	PopulateDb(mapOne, mapTwo)
	assert.Equal(t, mapOne["yes"], mapTwo["yes"])
}

func TestValidateInput(t *testing.T) {
	var cmd = "blob"
	assert.False(t, ValidateCommand(cmd))
	cmd = "WrItE"
	assert.True(t, ValidateCommand(cmd))
	for _, c := range validCommands {
		assert.True(t, ValidateCommand(c))
	}
	var fullString = "read 10 20 30"
	res, err := ValidateInput(fullString)
	assert.NotNil(t, err)
	assert.Nil(t, res)
	fullString = "write"
	res, err = ValidateInput(fullString)
	assert.NotNil(t, err)
	assert.Nil(t, res)
	fullString = "write 10 20"
	res, err = ValidateInput(fullString)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "write", res.Command)
	assert.Equal(t, "10", res.Key)
	assert.Equal(t, "20", res.Value)

}
