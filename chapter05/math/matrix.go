package math

type Matrix3 [3][3]float32

func (m Matrix3) Mul(r Matrix3) Matrix3 {
	ret := Matrix3{}
	// row 0
	ret[0][0] = m[0][0]*r[0][0] + m[0][1]*r[1][0] + m[0][2]*r[2][0]
	ret[0][1] = m[0][0]*r[0][1] + m[0][1]*r[1][1] + m[0][2]*r[2][1]
	ret[0][2] = m[0][0]*r[0][2] + m[0][1]*r[1][2] + m[0][2]*r[2][2]

	// row 1
	ret[1][0] = m[1][0]*r[0][0] + m[1][1]*r[1][0] + m[1][2]*r[2][0]
	ret[1][1] = m[1][0]*r[0][1] + m[1][1]*r[1][1] + m[1][2]*r[2][1]
	ret[1][2] = m[1][0]*r[0][2] + m[1][1]*r[1][2] + m[1][2]*r[2][2]

	// row 2
	ret[2][0] = m[2][0]*r[0][0] + m[2][1]*r[1][0] + m[2][2]*r[2][0]
	ret[2][1] = m[2][0]*r[0][1] + m[2][1]*r[1][1] + m[2][2]*r[2][1]
	ret[2][2] = m[2][0]*r[0][2] + m[2][1]*r[1][2] + m[2][2]*r[2][2]

	return ret
}

func Matrix3Identity() Matrix3 {
	return Matrix3{
		{1.0, 0.0, 0.0},
		{0.0, 1.0, 0.0},
		{0.0, 0.0, 1.0},
	}
}

func Matrix3CreateScale(scaleX, scaleY float32) Matrix3 {
	return Matrix3{
		{scaleX, 0.0, 0.0},
		{0.0, scaleY, 0.0},
		{0.0, 0.0, 1.0},
	}
}

// Matrix3CreateUniScale creates a scale matrix with a uniform factor
func Matrix3CreateUniScale(scale float32) Matrix3 {
	return Matrix3CreateScale(scale, scale)
}

func Matrix3CreateRotation(angle Angle) Matrix3 {
	return Matrix3{
		{Cos(angle), Sin(angle), 0.0},
		{-Sin(angle), Cos(angle), 0.0},
		{0.0, 0.0, 1.0},
	}
}

func Matrix3CreateTranslation(trans Vector2) Matrix3 {
	return Matrix3{
		{1.0, 0.0, 0.0},
		{0.0, 1.0, 0.0},
		{trans.X, trans.Y, 1.0},
	}
}

type Matrix4 [4][4]float32

func (m Matrix4) GetAsFloatPtr() *float32 {
	return &m[0][0]
}

func (m Matrix4) Mul(r Matrix4) Matrix4 {
	ret := Matrix4{}
	// row 0
	ret[0][0] = m[0][0]*r[0][0] + m[0][1]*r[1][0] + m[0][2]*r[2][0] + m[0][3]*r[3][0]
	ret[0][1] = m[0][0]*r[0][1] + m[0][1]*r[1][1] + m[0][2]*r[2][1] + m[0][3]*r[3][1]
	ret[0][2] = m[0][0]*r[0][2] + m[0][1]*r[1][2] + m[0][2]*r[2][2] + m[0][3]*r[3][2]
	ret[0][3] = m[0][0]*r[0][3] + m[0][1]*r[1][3] + m[0][2]*r[2][3] + m[0][3]*r[3][3]

	// row 1
	ret[1][0] = m[1][0]*r[0][0] + m[1][1]*r[1][0] + m[1][2]*r[2][0] + m[1][3]*r[3][0]
	ret[1][1] = m[1][0]*r[0][1] + m[1][1]*r[1][1] + m[1][2]*r[2][1] + m[1][3]*r[3][1]
	ret[1][2] = m[1][0]*r[0][2] + m[1][1]*r[1][2] + m[1][2]*r[2][2] + m[1][3]*r[3][2]
	ret[1][3] = m[1][0]*r[0][3] + m[1][1]*r[1][3] + m[1][2]*r[2][3] + m[1][3]*r[3][3]

	// row 2
	ret[2][0] = m[2][0]*r[0][0] + m[2][1]*r[1][0] + m[2][2]*r[2][0] + m[2][3]*r[3][0]
	ret[2][1] = m[2][0]*r[0][1] + m[2][1]*r[1][1] + m[2][2]*r[2][1] + m[2][3]*r[3][1]
	ret[2][2] = m[2][0]*r[0][2] + m[2][1]*r[1][2] + m[2][2]*r[2][2] + m[2][3]*r[3][2]
	ret[2][3] = m[2][0]*r[0][3] + m[2][1]*r[1][3] + m[2][2]*r[2][3] + m[2][3]*r[3][3]

	// row 3
	ret[3][0] = m[3][0]*r[0][0] + m[3][1]*r[1][0] + m[3][2]*r[2][0] + m[3][3]*r[3][0]
	ret[3][1] = m[3][0]*r[0][1] + m[3][1]*r[1][1] + m[3][2]*r[2][1] + m[3][3]*r[3][1]
	ret[3][2] = m[3][0]*r[0][2] + m[3][1]*r[1][2] + m[3][2]*r[2][2] + m[3][3]*r[3][2]
	ret[3][3] = m[3][0]*r[0][3] + m[3][1]*r[1][3] + m[3][2]*r[2][3] + m[3][3]*r[3][3]

	return ret
}

func (m Matrix4) GetTranslation() Vector3 {
	return Vector3{X: m[3][0], Y: m[3][1], Z: m[3][2]}
}

func (m Matrix4) GetXAxis() Vector3 {
	return Vector3{X: m[0][0], Y: m[0][1], Z: m[0][2]}
}

func (m Matrix4) GetYAxis() Vector3 {
	return Vector3{X: m[1][0], Y: m[1][1], Z: m[1][2]}
}

func (m Matrix4) GetZAxis() Vector3 {
	return Vector3{X: m[2][0], Y: m[2][1], Z: m[2][2]}
}

func (m Matrix4) GetScale() Vector3 {
	return Vector3{X: m.GetXAxis().Length(), Y: m.GetYAxis().Length(), Z: m.GetZAxis().Length()}
}

func Matrix4Identity() Matrix4 {
	return Matrix4{
		{1.0, 0.0, 0.0, 0.0},
		{0.0, 1.0, 0.0, 0.0},
		{0.0, 0.0, 1.0, 0.0},
		{0.0, 0.0, 0.0, 1.0},
	}
}

func Matrix4CreateScale(scaleX, scaleY, scaleZ float32) Matrix4 {
	return Matrix4{
		{scaleX, 0.0, 0.0},
		{0.0, scaleY, 0.0},
		{0.0, 0.0, scaleZ, 0.0},
		{0.0, 0.0, 0.0, 1.0},
	}
}

// Matrix4CreateUniScale creates a scale matrix with a uniform factor
func Matrix4CreateUniScale(scale float32) Matrix4 {
	return Matrix4CreateScale(scale, scale, scale)
}

// Matrix4CreateRotationX rotation about x-axis
func Matrix4CreateRotationX(angle Angle) Matrix4 {
	return Matrix4{
		{1.0, 0.0, 0.0, 0.0},
		{0.0, Cos(angle), Sin(angle), 0.0},
		{0.0, -Sin(angle), Cos(angle), 0.0},
		{0.0, 0.0, 0.0, 1.0},
	}
}

// Matrix4CreateRotationY rotation about y-axis
func Matrix4CreateRotationY(angle Angle) Matrix4 {
	return Matrix4{
		{Cos(angle), 0.0, -Sin(angle), 0.0},
		{0.0, 1.0, 0.0, 0.0},
		{Sin(angle), 0.0, Cos(angle), 0.0},
		{0.0, 0.0, 0.0, 1.0},
	}
}

// Matrix4CreateRotationZ rotation about z-axis
func Matrix4CreateRotationZ(angle Angle) Matrix4 {
	return Matrix4{
		{Cos(angle), Sin(angle), 0.0, 0.0},
		{-Sin(angle), Cos(angle), 0.0, 0.0},
		{0.0, 0.0, 1.0, 0.0},
		{0.0, 0.0, 0.0, 1.0},
	}
}

func Matrix4CreateTranslation(trans Vector3) Matrix4 {
	return Matrix4{
		{1.0, 0.0, 0.0, 0.0},
		{0.0, 1.0, 0.0, 0.0},
		{0.0, 0.0, 1.0, 0.0},
		{trans.X, trans.Y, trans.Z, 1.0},
	}
}

func Matrix4CreateLookAt(eye, target, up Vector3) Matrix4 {
	zaxis := target.Sub(eye).Normalize()
	xaxis := up.Cross(zaxis).Normalize()
	yaxis := zaxis.Cross(xaxis).Normalize()
	trans := Vector3{
		X: -xaxis.Dot(eye),
		Y: -yaxis.Dot(eye),
		Z: -zaxis.Dot(eye),
	}

	return Matrix4{
		{xaxis.X, yaxis.X, zaxis.X, 0.0},
		{xaxis.Y, yaxis.Y, zaxis.Y, 0.0},
		{xaxis.Z, yaxis.Z, zaxis.Z, 0.0},
		{trans.X, trans.Y, trans.Z, 1.0},
	}
}

func Matrix4CreateOrtho(width, height, near, far float32) Matrix4 {
	return Matrix4{
		{2.0 / width, 0.0, 0.0, 0.0},
		{0.0, 2.0 / height, 0.0, 0.0},
		{0.0, 0.0, 1.0 / (far - near), 0.0},
		{0.0, 0.0, near / (near - far), 1.0},
	}
}

func Matrix4CreatePerspectiveFOV(fovY, width, height, near, far float32) Matrix4 {
	yScale := Cot(Angle(fovY / 2.0))
	xScale := yScale * height / width
	return Matrix4{
		{xScale, 0.0, 0.0, 0.0},
		{0.0, yScale, 0.0, 0.0},
		{0.0, 0.0, far / (far - near), 1.0},
		{0.0, 0.0, -near * far / (far - near), 0.0},
	}
}

func Matrix4CreateSimpleViewProj(width, height float32) Matrix4 {
	return Matrix4{
		{2.0 / width, 0.0, 0.0, 0.0},
		{0.0, 2.0 / height, 0.0, 0.0},
		{0.0, 0.0, 1.0, 0.0},
		{0.0, 0.0, 1.0, 1.0},
	}
}
