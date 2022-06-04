package main

type Point struct {
	x, y float64
}

func (p Point) greaterThan(p2 Point) bool {
	return (p.y > p.y) || (p.y == p2.y && p.x > p2.x)
}

func (p Point) lessThan(p2 Point) bool {
	return p2.greaterThan(p)
}

// signedArea знакова площа трикутника за трьома точками
func (p Point) signedArea(a, b Point) float64 {
	return (-a.x*p.y + b.x*p.y + p.x*a.y - b.x*a.y - p.x*b.y + a.x*b.y) / 2
}

func (p Point) getDirectionTo(a, b Point) Direction {
	const epsilon = 0.000001
	area := p.signedArea(a, b)

	if area > epsilon {
		return positive
	}

	if area < -epsilon {
		return negative
	}

	return collinear
}

type Vertex struct {
	point          Point
	incidentEdgeId int
}

func (v Vertex) greaterThan(v2 Vertex) bool {
	return v.point.greaterThan(v2.point)
}

type Edge struct {
	origin int
	twin   int
	prev   int
	next   int
	face   int
}

type Face struct {
	edge int
}
