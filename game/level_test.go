package game

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadingFromReader(t *testing.T) {
	buffer := bytes.NewBuffer([]byte("* \n *"))
	level := &Level{
		SymbolMap: map[string]string{
			"*": "wall",
			" ": "empty",
		},
	}

	err := level.Load(buffer)

	assert.Nil(t, err)
	assert.Equal(t, "wall", level.GetTile(0, 0).Kind)
	assert.Equal(t, "empty", level.GetTile(0, 1).Kind)
	assert.Equal(t, "empty", level.GetTile(1, 0).Kind)
	assert.Equal(t, "wall", level.GetTile(1, 1).Kind)
}

func TestLoadingWithMissingSymbol(t *testing.T) {
	buffer := bytes.NewBuffer([]byte(" "))
	level := &Level{}

	err := level.Load(buffer)
	assert.NotNil(t, err)
	assert.Equal(t, "level: ' ' not found in SymbolMap", err.Error())
}
