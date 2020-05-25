package tracer

import "math"

// Vector - represents a vector
type Vector struct {
	X float64
	Y float64
	Z float64
}

// Normalized - returns the vector but normalized
func (v *Vector) Normalized() *Vector {
	magitude := v.Magitude()
	return &Vector{
		X: v.X / magitude,
		Y: v.Y / magitude,
		Z: v.Z / magitude,
	}
}

// Dot - the dot pruduct
func (v *Vector) Dot(other *Vector) float64 {
	return (v.X * other.X) + (v.Y * other.Y) + (v.Z * other.Z)
}

func (v *Vector) Scalar(n float64) *Vector {
	return &Vector{
		X: n * v.X,
		Y: n * v.Y,
		Z: n * v.Z,
	}
}

func (v *Vector) Add(other *Vector) *Vector {
	return &Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

func (v *Vector) Sub(other *Vector) *Vector {
	return v.Add(other.Scalar(-1))
}

func (v *Vector) Magitude() float64 {
	return math.Sqrt(v.Dot(v))
}
