package model

import (
	"math"
)

// Vertex is a point in 2D space
type Vertex struct {
	X, Y float64
}

// Abs returns the absolute value of the vertex
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale scales the vertex by some factor
// It mutates the vertex
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
