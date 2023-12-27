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

const SALT = 4

// returns reversed slice of indexes to tiles of shortest path in level
func (level *Level) GetPath(pX, pY, nX, nY int, random bool, no_walls bool) []int {
	start := &Point{X: pX, Y: pY}
	end := &Point{X: nX, Y: nY}

	var g graph.Graph[int, *Point]

	if random {
		g = level.createGraph(start, end, SALT, no_walls)
	} else {
		g = level.createGraph(start, end, 1, no_walls)
	}

	points, e := graph.ShortestPath(g,
		level.GetIndexFromXY(pX, pY),
		level.GetIndexFromXY(nX, nY))

	if e == graph.ErrTargetNotReachable {
		return nil
	}
	return points
}

func (level *Level) createGraph(p1, p2 *Point, salt int, no_walls bool) graph.Graph[int, *Point] {
	// hash is the tile index
	pointHash := func(p *Point) int {
		return level.GetIndexFromXY(p.X, p.Y)
	}

	g := graph.New(pointHash, graph.Weighted())
	x1, y1 := p1.X, p1.Y
	x2, y2 := p2.X, p2.Y

	graph_nodes := make([]*Point, 0)

	for x := min(x1, x2); x < max(x1, x2)+1; x++ {
		for y := min(y1, y2); y < max(y1, y2)+1; y++ {
			if no_walls {
				idx := level.GetIndexFromXY(x, y)
				if level.Tiles[idx].TileType != WALL {
					p := Point{X: x, Y: y}
					g.AddVertex(&p)
					graph_nodes = append(graph_nodes, &p)
				}
			} else {
				p := Point{X: x, Y: y}
				g.AddVertex(&p)
				graph_nodes = append(graph_nodes, &p)
			}
		}
	}

	for _, p := range graph_nodes {
		x, y := p.X, p.Y
		var e error
		// randomize weights for path gen
		if x != max(x1, x2) && checkPointInList(x+1, y, graph_nodes) {
			e = g.AddEdge(
				level.GetIndexFromXY(x, y),
				level.GetIndexFromXY(x+1, y),
				graph.EdgeWeight(utils.GetDiceRoll(salt)))
			if e != nil && e != graph.ErrEdgeAlreadyExists {
				log.Fatal(e)
			}
		}
		if x != min(x1, x2) && checkPointInList(x-1, y, graph_nodes) {
			e = g.AddEdge(
				level.GetIndexFromXY(x, y),
				level.GetIndexFromXY(x-1, y),
				graph.EdgeWeight(utils.GetDiceRoll(salt)))
			if e != nil && e != graph.ErrEdgeAlreadyExists {
				log.Fatal(e)
			}
		}
		if y != max(y1, y2) && checkPointInList(x, y+1, graph_nodes) {
			e = g.AddEdge(
				level.GetIndexFromXY(x, y),
				level.GetIndexFromXY(x, y+1),
				graph.EdgeWeight(utils.GetDiceRoll(salt)))
			if e != nil && e != graph.ErrEdgeAlreadyExists {
				log.Fatal(e)
			}
		}
		if y != min(y1, y2) && checkPointInList(x, y-1, graph_nodes) {
			e = g.AddEdge(
				level.GetIndexFromXY(x, y),
				level.GetIndexFromXY(x, y-1),
				graph.EdgeWeight(utils.GetDiceRoll(salt)))
			if e != nil && e != graph.ErrEdgeAlreadyExists {
				log.Fatal(e)
			}
		}
	}

	return g
}

func checkPointInList(x, y int, points []*Point) bool {
	in_list := false
	for _, p := range points {
		if p.X == x && p.Y == y {
			in_list = true
		}
	}
	return in_list
}
