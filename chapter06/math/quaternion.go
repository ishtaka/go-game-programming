package math

type Quaternion struct {
	x, y, z, w float32
}

func QuaternionIdentity() *Quaternion {
	return &Quaternion{
		x: 0.0,
		y: 0.0,
		z: 0.0,
		w: 1.0,
	}
}

func NewQuaternion(x, y, z, w float32) *Quaternion {
	return &Quaternion{
		x: x,
		y: y,
		z: z,
		w: w,
	}
}

func NewQuaternionFromVec(axis Vector3, angle float32) *Quaternion {
	scalar := Sin(Angle(angle / 2.0))
	return &Quaternion{
		x: axis.X * scalar,
		y: axis.Y * scalar,
		z: axis.Z * scalar,
		w: Cos(Angle(angle / 2.0)),
	}
}

func (q *Quaternion) Set(x, y, z, w float32) {
	q.x = x
	q.y = y
	q.z = z
	q.w = w
}

func (q *Quaternion) Conjugate() {
	q.x *= -1.0
	q.y *= -1.0
	q.z *= -1.0
}

func (q *Quaternion) LengthSq() float32 {
	return q.x*q.x + q.y*q.y + q.z*q.z + q.w*q.w
}

func (q *Quaternion) Length() float32 {
	return Sqrt(q.LengthSq())
}

func (q *Quaternion) Normalize() {
	l := q.Length()
	q.x /= l
	q.y /= l
	q.z /= l
	q.w /= l
}

func (q *Quaternion) Dot(b *Quaternion) float32 {
	return q.x*b.x + q.y*b.y + q.z*b.z + q.w*b.w
}

func (q *Quaternion) Lerp(b *Quaternion, f float32) *Quaternion {
	ret := &Quaternion{
		x: Lerp(q.x, b.x, f),
		y: Lerp(q.y, b.y, f),
		z: Lerp(q.z, b.z, f),
		w: Lerp(q.w, b.w, f),
	}
	ret.Normalize()

	return ret
}

func (q *Quaternion) Slerp(b *Quaternion, f float32) *Quaternion {
	rawCosm := q.Dot(b)

	cosom := -rawCosm
	if rawCosm >= 0.0 {
		cosom = rawCosm
	}

	var scale0, scale1 float32

	if cosom < 0.9999 {
		omega := Acos(Angle(cosom))
		invSin := 1.0 / Sin(omega)
		scale0 = Sin(Angle((1.0-f)*float32(omega))) * invSin
		scale1 = Sin(Angle(f*float32(omega))) * invSin
	} else {
		// Use linear interpolation if the quaternions
		// are collinear
		scale0 = 1.0 - f
		scale1 = f
	}

	if rawCosm < 0.0 {
		scale1 = -scale1
	}

	ret := &Quaternion{
		x: scale0*q.x + scale1*b.x,
		y: scale0*q.y + scale1*b.y,
		z: scale0*q.z + scale1*b.z,
		w: scale0*q.w + scale1*b.w,
	}
	ret.Normalize()

	return ret
}

// Concatenate rotate by q FOLLOWED BY p
func (q *Quaternion) Concatenate(p *Quaternion) *Quaternion {
	// Vector component is:
	// ps * qv + qs * pv + pv x qv
	qv := Vector3{q.x, q.y, q.z}
	pv := Vector3{p.x, p.y, p.z}
	newVec := qv.MulScalar(p.w).Add(pv.MulScalar(q.w)).Add(pv.Cross(qv))

	ret := &Quaternion{
		x: newVec.X,
		y: newVec.Y,
		z: newVec.Z,
	}

	// Scalar component is:
	// ps * qs - pv . qv
	ret.w = p.w*q.w - pv.Dot(qv)

	return ret
}
