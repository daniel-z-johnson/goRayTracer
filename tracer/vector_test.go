package tracer

import (
	"math"
	"testing"
)

const (
	espilon = 0.001
)

var (
	a = Vector{X: 1, Y: 2, Z: 3}
	b = Vector{X: 123, Y: 223, Z: 323}
	c = Vector{X: 122, Y: 221, Z: 320}
)

func TestNormalize(t *testing.T) {
	na := a.Normalized()
	t.Logf("Magitude: '%f'", na.Magitude())
	if !compareFloats(1, na.Magitude()) {
		t.Fail()
	}
	na = b.Normalized()
	t.Logf("Magitude: '%f'", na.Magitude())
	if !compareFloats(1, na.Magitude()) {
		t.Fail()
	}
}

func TestSub(t *testing.T) {
	ac := b.Sub(&a)
	if c.X != ac.X || c.Y != ac.Y || c.Z != ac.Z {
		t.Fatalf("vecter(%+v) should equal vecter(%+v)", ac, c)
	}
}

func TestScalar(t *testing.T) {
	da := a.Scalar(2)
	if da.X != 2*a.X || da.Y != 2*a.Y || da.Z != 2*a.Z {
		t.Fail()
	}
}

func compareFloats(a, b float64) bool {
	diff := math.Abs(a - b)
	return diff < espilon
}
