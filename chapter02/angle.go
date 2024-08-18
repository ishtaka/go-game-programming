package chapter02

import "math"

type Angle float64

const (
	Radian Angle = 1
	Degree       = (math.Pi / 180) * Radian
)

// Radians returns the angle in radians
func (a Angle) Radians() float64 {
	return float64(a)
}

// Degrees returns the angle in degrees
func (a Angle) Degrees() float64 {
	return float64(a / Degree)
}
