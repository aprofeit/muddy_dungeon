package game

type Entity struct {
	Position
	level *Level
}

func (e *Entity) SetLevel(l *Level) {
	if e.level != nil {
		e.level.removeEntity(e)
	}

	l.addEntity(e)
	e.level = l
}
