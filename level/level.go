package level

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

type Tile struct {
	Kind string
}

type Level struct {
	Tiles     [][]Tile
	SymbolMap map[string]string
}

func (l *Level) GetTile(x, y int) Tile {
	return l.Tiles[y][x]
}

func (l *Level) Load(input io.ReadWriter) error {
	scanner := bufio.NewScanner(input)

	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		l.Tiles = append(l.Tiles, make([]Tile, len(line)))

		for i := 0; i < len(line); i++ {
			char := string(line[i])

			kind, exists := l.SymbolMap[char]
			if !exists {
				return errors.New(fmt.Sprintf("level: %s not found in SymbolMap", char))
			}

			l.Tiles[row][i] = Tile{Kind: kind}
		}

		row++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}