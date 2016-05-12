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
	level1 := newTestLevel()
	entity := &Entity{}

	entity.SetLevel(level1)

	assert.Equal(t, 1, len(level1.entities))
	assert.Equal(t, entity, level1.entities[0])
	assert.Equal(t, level1, entity.level)

	level2 := newTestLevel()

	entity.SetLevel(level2)

	assert.Equal(t, 0, len(level1.entities))
	assert.Equal(t, 1, len(level2.entities))
	assert.Equal(t, entity, level2.entities[0])
	assert.Equal(t, level2, entity.level)
}

func newTestLevel() *Level {
	buffer := bytes.NewBuffer([]byte("  \n  "))
	level := NewLevel()
	level.Load(buffer)

	return level
}
