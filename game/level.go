package game

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

type Level struct {
	tiles     [][]Tile
	SymbolMap map[string]string
	entities  []*Entity
}

func NewLevel() *Level {
	return &Level{
		SymbolMap: map[string]string{
			"*": "wall",
			" ": "empty",
		},
	}
}

func (l *Level) GetTile(x, y int) Tile {
	return l.tiles[y][x]
}

func (l *Level) Load(input io.ReadWriter) error {
	scanner := bufio.NewScanner(input)

	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		l.tiles = append(l.tiles, make([]Tile, len(line)))

		for col := 0; col < len(line); col++ {
			char := string(line[col])

			kind, exists := l.SymbolMap[char]
			if !exists {
				return errors.New(fmt.Sprintf("level: '%s' not found in SymbolMap", char))
			}

			l.tiles[row][col] = Tile{Kind: kind}
		}

		row++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (l *Level) addEntity(e *Entity) {
	l.entities = append(l.entities, e)
}

func (l *Level) removeEntity(e *Entity) {
	for i := 0; i < len(l.entities); i++ {
		if l.entities[i] == e {
			l.entities = append(l.entities[:i], l.entities[i+1:]...)
		}
	}
}
