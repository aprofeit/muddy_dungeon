package game

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatingAnEntity(t *testing.T) {
	entity := &Entity{
		Position: Position{X: 0.0, Y: 0.0},
	}

	assert.Equal(t, 0.0, entity.X)
	assert.Equal(t, 0.0, entity.Y)
}

func TestAddingEntityToLevel(t *testing.T) {
	level := newTestLevel()
	entity := &Entity{}

	entity.SetLevel(level)

	assert.Equal(t, 1, len(level.entities))
	assert.Equal(t, entity, level.entities[0])
	assert.Equal(t, entity.level, level)
}

func newTestLevel() *Level {
	buffer := bytes.NewBuffer([]byte("  \n  "))
	level := NewLevel()
	level.Load(buffer)

	return level
}
