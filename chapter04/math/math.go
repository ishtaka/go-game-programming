package math

import "math"

const Pi float32 = 3.1415926535

var TwoPi float32 = Pi * 2
var PiOver2 float32 = Pi / 2
var Infinity = math.Inf(1)
var NegInfinity = math.Inf(-1)

type Angle float64

const (
	Radian Angle = 1
	Degree       = (math.Pi / 180) * Radian
)

func (a Angle) Add(angle float32) Angle {
	return a + Angle(angle)
}

// Radians returns the angle in radians
func (a Angle) Radians() float64 {
	return float64(a)
}

// Degrees returns the angle in degrees
func (a Angle) Degrees() float64 {
	return float64(a / Degree)
}

func NearZero(value float32) bool {
	const epsilon float64 = 0.001
	return math.Abs(float64(value)) <= epsilon
}

func Max(a, b float32) float32 {
	if a < b {
		return b
	}
	return a
}

func Min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func Clamp(value, min, max float32) float32 {
	return Min(max, Max(min, value))
}

func Abs(value float32) float32 {
	return float32(math.Abs(float64(value)))
}

func Cos(angle Angle) float32 {
	return float32(math.Cos(angle.Radians()))
}

func Sin(angle Angle) float32 {
	return float32(math.Sin(angle.Radians()))
}

func Tan(angle Angle) float32 {
	return float32(math.Tan(angle.Radians()))
}

func Acos(value float32) Angle {
	return Angle(math.Acos(float64(value)))
}

func Atan2(y, x float32) Angle {
	return Angle(math.Atan2(float64(y), float64(x)))
}

func Cot(angle Angle) float32 {
	return 1 / Tan(angle)
}

func Lerp(a, b, f float32) float32 {
	return a + f*(b-a)
}

func Sqrt(value float32) float32 {
	return float32(math.Sqrt(float64(value)))
}

func Fmod(a, b float32) float32 {
	return float32(math.Mod(float64(a), float64(b)))
}
