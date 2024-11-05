package math

import "math"

var ZeroVector2 = Vector2{0, 0}

// Vector2 is a 2D vector
type Vector2 struct {
	X, Y float32
}

// Add returns Vector2 addition (a + b)
func (a Vector2) Add(b Vector2) Vector2 {
	return Vector2{a.X + b.X, a.Y + b.Y}
}

// Sub returns Vector2 subtraction (a - b)
func (a Vector2) Sub(b Vector2) Vector2 {
	return Vector2{a.X - b.X, a.Y - b.Y}
}

// Mul returns Vector2 multiplication (a.X * b.X, ...)
func (a Vector2) Mul(b Vector2) Vector2 {
	return Vector2{a.X * b.X, a.Y * b.Y}
}

// MulScalar returns Vector2 multiplied by a scalar (a * s)
func (a Vector2) MulScalar(s float32) Vector2 {
	return Vector2{a.X * s, a.Y * s}
}

// LengthSq returns squared of Vector2
func (a Vector2) LengthSq() float32 {
	return (a.X * a.X) + (a.Y * a.Y)
}

// Length returns length of Vector2
func (a Vector2) Length() float32 {
	return float32(math.Sqrt(float64(a.LengthSq())))
}

// Normalize returns normalized Vector2
func (a Vector2) Normalize() Vector2 {
	length := a.Length()
	if length != 0 {
		return Vector2{a.X / length, a.Y / length}
	}

	return Vector2{0, 0}
}

// Dot returns dot product of Vector2
func (a Vector2) Dot(b Vector2) float32 {
	return (a.X * b.X) + (a.Y * b.Y)
}

// Lerp returns linear interpolation from a to b by f
func (a Vector2) Lerp(b Vector2, f float32) Vector2 {
	return a.Add(b.Sub(a).MulScalar(f))
}

// Reflect returns Reflect a about (normalized) b
func (a Vector2) Reflect(b Vector2) Vector2 {
	return a.Sub(b.MulScalar(2.0 * a.Dot(b)))
}

// Vector3 is a 3D vector
type Vector3 struct {
	X, Y, Z float32
}

// Add returns Vector3 addition (a + b)
func (a Vector3) Add(b Vector3) Vector3 {
	return Vector3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Sub returns Vector3 subtraction (a - b)
func (a Vector3) Sub(b Vector3) Vector3 {
	return Vector3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Mul returns Vector3 multiplication (a.X * b.X, ...)
func (a Vector3) Mul(b Vector3) Vector3 {
	return Vector3{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

// MulScalar returns Vector3 multiplied by a scalar (a * s)
func (a Vector3) MulScalar(s float32) Vector3 {
	return Vector3{a.X * s, a.Y * s, a.Z * s}
}

// LengthSq returns squared of Vector3
func (a Vector3) LengthSq() float32 {
	return (a.X * a.X) + (a.Y * a.Y) + (a.Z * a.Z)
}

// Length returns length of Vector3
func (a Vector3) Length() float32 {
	return float32(math.Sqrt(float64(a.LengthSq())))
}

// Normalize returns normalized Vector3
func (a Vector3) Normalize() Vector3 {
	length := a.Length()
	if length != 0 {
		return Vector3{a.X / length, a.Y / length, a.Z / length}
	}

	return Vector3{0, 0, 0}
}

// Dot product between two vectors (a dot b)
func (a Vector3) Dot(b Vector3) float32 {
	return (a.X * b.X) + (a.Y * b.Y) + (a.Z * b.Z)
}

// Cross product between two vectors (a cross b)
func (a Vector3) Cross(b Vector3) Vector3 {
	return Vector3{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

// Lerp returns linear interpolation from a to b by f
func (a Vector3) Lerp(b Vector3, f float32) Vector3 {
	return a.Add(b.Sub(a).MulScalar(f))
}

// Reflect returns Reflect a about (normalized) b
func (a Vector3) Reflect(b Vector3) Vector3 {
	return a.Sub(b.MulScalar(2.0 * a.Dot(b)))
}
