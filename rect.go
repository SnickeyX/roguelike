package main

// rect, (x1,y1) : top left point, (x2,y2): bottom right point
type Rect struct {
	X1 int
	X2 int
	Y1 int
	Y2 int
}

func NewRect(x int, y int, width int, height int) Rect {
	return Rect{
		X1: x,
		X2: x + width,
		Y1: y,
		Y2: y + height,
	}
}

// center of rect
func (r *Rect) Center() (int, int) {
	cX := (r.X1 + r.X2) / 2
	cY := (r.Y1 + r.Y2) / 2
	return cX, cY
}

// whether one rect intersects another
func (r *Rect) Intersect(other Rect) bool {
	return (r.X1 <= other.X2 && r.X2 >= other.X1 && r.Y1 <= other.Y1 && r.Y2 >= other.Y1)
}
