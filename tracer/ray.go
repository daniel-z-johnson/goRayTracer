package tracer

import "fmt"

type Ray struct {
	Point     Vector
	Direction Vector
}

func (r *Ray) String() string {
	p := r.Point
	d := r.Direction
	return fmt.Sprintf("point(%f, %f, %f) Direction(%f, %f, %f)",
		p.X, p.Y, p.Z,
		d.X, d.Y, d.Z)
}
