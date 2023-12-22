package world

import (
	"log"

	"github.com/SnickeyX/roguelike/utils"
	"github.com/dominikbraun/graph"
)

type Point struct {
	X int
	Y int
}

func (level *Level) CreateGraph(p1, p2 *Point) graph.Graph[int, *Point] {
	// hash is the tile index
	pointHash := func(p *Point) int {
		return level.GetIndexFromXY(p.X, p.Y)
	}

	g := graph.New(pointHash, graph.Weighted())
	x1, y1 := p1.X, p1.Y
	x2, y2 := p2.X, p2.Y

	for x := min(x1, x2); x < max(x1, x2)+1; x++ {
		for y := min(y1, y2); y < max(y1, y2)+1; y++ {
			g.AddVertex(&Point{X: x, Y: y})
		}
	}

	for x := min(x1, x2); x < max(x1, x2)+1; x++ {
		for y := min(y1, y2); y < max(y1, y2)+1; y++ {
			var e error
			// randomize weights for path gen
			if x != max(x1, x2) {
				e = g.AddEdge(
					level.GetIndexFromXY(x, y),
					level.GetIndexFromXY(x+1, y),
					graph.EdgeWeight(utils.GetDiceRoll(4)))
				if e != nil && e != graph.ErrEdgeAlreadyExists {
					log.Fatal(e)
				}
			}
			if x != min(x1, x2) {
				e = g.AddEdge(
					level.GetIndexFromXY(x, y),
					level.GetIndexFromXY(x-1, y),
					graph.EdgeWeight(utils.GetDiceRoll(4)))
				if e != nil && e != graph.ErrEdgeAlreadyExists {
					log.Fatal(e)
				}
			}
			if y != max(y1, y2) {
				e = g.AddEdge(
					level.GetIndexFromXY(x, y),
					level.GetIndexFromXY(x, y+1),
					graph.EdgeWeight(utils.GetDiceRoll(4)))
				if e != nil && e != graph.ErrEdgeAlreadyExists {
					log.Fatal(e)
				}
			}
			if y != min(y1, y2) {
				e = g.AddEdge(
					level.GetIndexFromXY(x, y),
					level.GetIndexFromXY(x, y-1),
					graph.EdgeWeight(utils.GetDiceRoll(4)))
				if e != nil && e != graph.ErrEdgeAlreadyExists {
					log.Fatal(e)
				}
			}
		}
	}
	return g
}

// returns slice of indexes to tiles of shortest path in level
func (level *Level) GetShortestPath(pX, pY, nX, nY int) []int {
	start := &Point{X: pX, Y: pY}
	end := &Point{X: nX, Y: nY}

	g := level.CreateGraph(start, end)

	points, e := graph.ShortestPath(g,
		level.GetIndexFromXY(pX, pY),
		level.GetIndexFromXY(nX, nY))

	if e == graph.ErrTargetNotReachable {
		log.Fatal(e)
	}
	return points
}
