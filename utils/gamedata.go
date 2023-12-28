package utils

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
	UiHeight     int
}

func NewGameData() GameData {
	g := GameData{
		ScreenWidth:  80,
		ScreenHeight: 60,
		TileWidth:    16,
		TileHeight:   16,
		UiHeight:     10,
	}
	return g
}

var GameConstants GameData = NewGameData()
