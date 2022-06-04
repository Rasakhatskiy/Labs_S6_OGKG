package main

import "math/rand"

func generateRandomPointsBetween(n int, xMin, yMin, xMax, yMax float64) []Point {
	var points []Point
	for i := 0; i < n; i++ {
		points = append(points, Point{
			x: xMin + rand.Float64()*(xMax-xMin),
			y: yMin + rand.Float64()*(yMax-yMin),
		})
	}
	return points
}
