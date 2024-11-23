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
	return Vector3{X: m[0][0], Y: m[0][1], Z: m[0][2]}.Normalize()
}

func (m Matrix4) GetYAxis() Vector3 {
	return Vector3{X: m[1][0], Y: m[1][1], Z: m[1][2]}.Normalize()
}

func (m Matrix4) GetZAxis() Vector3 {
	return Vector3{X: m[2][0], Y: m[2][1], Z: m[2][2]}.Normalize()
}

func (m Matrix4) GetScale() Vector3 {
	return Vector3{
		X: Vector3{X: m[0][0], Y: m[0][1], Z: m[0][2]}.Length(),
		Y: Vector3{X: m[1][0], Y: m[1][1], Z: m[1][2]}.Length(),
		Z: Vector3{X: m[2][0], Y: m[2][1], Z: m[2][2]}.Length(),
	}
}

func (m Matrix4) Invert() Matrix4 {
	// Thanks slow math
	// This is a really janky way to unroll everything...
	var tmp [12]float32
	var src, dst [16]float32
	var det float32

	// Transpose matrix
	// row 1 to col 1
	src[0] = m[0][0]
	src[4] = m[0][1]
	src[8] = m[0][2]
	src[12] = m[0][3]

	// row 2 to col 2
	src[1] = m[1][0]
	src[5] = m[1][1]
	src[9] = m[1][2]
	src[13] = m[1][3]

	// row 3 to col 3
	src[2] = m[2][0]
	src[6] = m[2][1]
	src[10] = m[2][2]
	src[14] = m[2][3]

	// row 4 to col 4
	src[3] = m[3][0]
	src[7] = m[3][1]
	src[11] = m[3][2]
	src[15] = m[3][3]

	// Calculate cofactors
	tmp[0] = src[10] * src[15]
	tmp[1] = src[11] * src[14]
	tmp[2] = src[9] * src[15]
	tmp[3] = src[11] * src[13]
	tmp[4] = src[9] * src[14]
	tmp[5] = src[10] * src[13]
	tmp[6] = src[8] * src[15]
	tmp[7] = src[11] * src[12]
	tmp[8] = src[8] * src[14]
	tmp[9] = src[10] * src[12]
	tmp[10] = src[8] * src[13]
	tmp[11] = src[9] * src[12]

	dst[0] = tmp[0]*src[5] + tmp[3]*src[6] + tmp[4]*src[7]
	dst[0] -= tmp[1]*src[5] + tmp[2]*src[6] + tmp[5]*src[7]
	dst[1] = tmp[1]*src[4] + tmp[6]*src[6] + tmp[9]*src[7]
	dst[1] -= tmp[0]*src[4] + tmp[7]*src[6] + tmp[8]*src[7]
	dst[2] = tmp[2]*src[4] + tmp[7]*src[5] + tmp[10]*src[7]
	dst[2] -= tmp[3]*src[4] + tmp[6]*src[5] + tmp[11]*src[7]
	dst[3] = tmp[5]*src[4] + tmp[8]*src[5] + tmp[11]*src[6]
	dst[3] -= tmp[4]*src[4] + tmp[9]*src[5] + tmp[10]*src[6]
	dst[4] = tmp[1]*src[1] + tmp[2]*src[2] + tmp[5]*src[3]
	dst[4] -= tmp[0]*src[1] + tmp[3]*src[2] + tmp[4]*src[3]
	dst[5] = tmp[0]*src[0] + tmp[7]*src[2] + tmp[8]*src[3]
	dst[5] -= tmp[1]*src[0] + tmp[6]*src[2] + tmp[9]*src[3]
	dst[6] = tmp[3]*src[0] + tmp[6]*src[1] + tmp[11]*src[3]
	dst[6] -= tmp[2]*src[0] + tmp[7]*src[1] + tmp[10]*src[3]
	dst[7] = tmp[4]*src[0] + tmp[9]*src[1] + tmp[10]*src[2]
	dst[7] -= tmp[5]*src[0] + tmp[8]*src[1] + tmp[11]*src[2]

	tmp[0] = src[2] * src[7]
	tmp[1] = src[3] * src[6]
	tmp[2] = src[1] * src[7]
	tmp[3] = src[3] * src[5]
	tmp[4] = src[1] * src[6]
	tmp[5] = src[2] * src[5]
	tmp[6] = src[0] * src[7]
	tmp[7] = src[3] * src[4]
	tmp[8] = src[0] * src[6]
	tmp[9] = src[2] * src[4]
	tmp[10] = src[0] * src[5]
	tmp[11] = src[1] * src[4]

	dst[8] = tmp[0]*src[13] + tmp[3]*src[14] + tmp[4]*src[15]
	dst[8] -= tmp[1]*src[13] + tmp[2]*src[14] + tmp[5]*src[15]
	dst[9] = tmp[1]*src[12] + tmp[6]*src[14] + tmp[9]*src[15]
	dst[9] -= tmp[0]*src[12] + tmp[7]*src[14] + tmp[8]*src[15]
	dst[10] = tmp[2]*src[12] + tmp[7]*src[13] + tmp[10]*src[15]
	dst[10] -= tmp[3]*src[12] + tmp[6]*src[13] + tmp[11]*src[15]
	dst[11] = tmp[5]*src[12] + tmp[8]*src[13] + tmp[11]*src[14]
	dst[11] -= tmp[4]*src[12] + tmp[9]*src[13] + tmp[10]*src[14]
	dst[12] = tmp[2]*src[10] + tmp[5]*src[11] + tmp[1]*src[9]
	dst[12] -= tmp[4]*src[11] + tmp[0]*src[9] + tmp[3]*src[10]
	dst[13] = tmp[8]*src[11] + tmp[0]*src[8] + tmp[7]*src[10]
	dst[13] -= tmp[6]*src[10] + tmp[9]*src[11] + tmp[1]*src[8]
	dst[14] = tmp[6]*src[9] + tmp[11]*src[11] + tmp[3]*src[8]
	dst[14] -= tmp[10]*src[11] + tmp[2]*src[8] + tmp[7]*src[9]
	dst[15] = tmp[10]*src[10] + tmp[4]*src[8] + tmp[9]*src[9]
	dst[15] -= tmp[8]*src[9] + tmp[11]*src[10] + tmp[5]*src[8]

	// Calculate determinant
	det = src[0]*dst[0] + src[1]*dst[1] + src[2]*dst[2] + src[3]*dst[3]

	// Inverse of matrix is divided by determinant
	det = 1 / det
	for j := 0; j < 16; j++ {
		dst[j] *= det
	}

	// Set it back
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m[i][j] = dst[i*4+j]
		}
	}

	return m
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

func Matrix4CreateFromQuaternion(q *Quaternion) Matrix4 {
	mat := Matrix4{}

	mat[0][0] = 1.0 - 2.0*q.y*q.y - 2.0*q.z*q.z
	mat[0][1] = 2.0*q.x*q.y + 2.0*q.w*q.z
	mat[0][2] = 2.0*q.x*q.z - 2.0*q.w*q.y
	mat[0][3] = 0.0

	mat[1][0] = 2.0*q.x*q.y - 2.0*q.w*q.z
	mat[1][1] = 1.0 - 2.0*q.x*q.x - 2.0*q.z*q.z
	mat[1][2] = 2.0*q.y*q.z + 2.0*q.w*q.x
	mat[1][3] = 0.0

	mat[2][0] = 2.0*q.x*q.z + 2.0*q.w*q.y
	mat[2][1] = 2.0*q.y*q.z - 2.0*q.w*q.x
	mat[2][2] = 1.0 - 2.0*q.x*q.x - 2.0*q.y*q.y
	mat[2][3] = 0.0

	mat[3][0] = 0.0
	mat[3][1] = 0.0
	mat[3][2] = 0.0
	mat[3][3] = 1.0

	return mat
}
