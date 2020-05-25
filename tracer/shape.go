package tracer

type Shape interface {
	Intersection(r *Ray) float64
	Normal(interceptionPoint, eyeRay *Vector)
}
