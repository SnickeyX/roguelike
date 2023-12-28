package utils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Logger struct {
	Log  map[int][]string
	Col  int
	Tics int
}

const NUM_MESSAGES_PER_COL = 5

func PrintInUI(screen *ebiten.Image, message []string, col int) {
	m := ""
	for _, str := range message {
		m += str
	}
	ebitenutil.DebugPrintAt(screen, m,
		16+(col-1)*(GameConstants.ScreenWidth*GameConstants.TileWidth)/3,
		(GameConstants.ScreenHeight*GameConstants.TileHeight)-
			((GameConstants.UiHeight*GameConstants.TileHeight)/2))
}

func CreateLogger() *Logger {
	logger := Logger{Log: make(map[int][]string), Col: 1, Tics: 0}
	return &logger
}

func (logger *Logger) Display(msg string) {
	if len(logger.Log[logger.Col]) == NUM_MESSAGES_PER_COL {
		logger.Col++
	}
	logger.Log[logger.Col] = append(logger.Log[logger.Col], msg)
}

func (logger *Logger) Clear() {
	for k := range logger.Log {
		delete(logger.Log, k)
	}
	logger.Tics = 0
	logger.Col = 1
}

func (logger *Logger) DrawColumnWise(screen *ebiten.Image) {
	has_drawn := false
	for k, v := range logger.Log {
		has_drawn = true
		PrintInUI(screen, v, k)
	}
	if has_drawn {
		logger.Tics++
	}
}
