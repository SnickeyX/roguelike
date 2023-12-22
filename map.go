package main

// container for dungeons (i.e. the entire game world)
type GameMap struct {
	Dungeons     []Dungeon
	CurrentLevel Level
}

func NewGameMap() GameMap {
	// single level game map for now
	l := NewLevel()
	levels := make([]Level, 0)
	levels = append(levels, l)
	// single dungeon with the previous levels
	d := Dungeon{Name: "First", Levels: levels}
	dungeons := make([]Dungeon, 0)
	dungeons = append(dungeons, d)
	gm := GameMap{Dungeons: dungeons, CurrentLevel: l}
	return gm
}
