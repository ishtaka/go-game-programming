package chapter06

import "github.com/ishtaka/go-game-programming/chapter06/math"

type MoveComponent interface {
	Component
	GetAngularSpeed() float32
	SetAngularSpeed(speed float32)
	GetForwardSpeed() float32
	SetForwardSpeed(speed float32)
}

type moveComponent struct {
	Component
	angularSpeed float32
	forwardSpeed float32
}

func NewMoveComponent(owner Actor, updateOrder int) MoveComponent {
	c := NewComponent(owner, updateOrder)
	mc := &moveComponent{
		Component: c,
	}

	return mc
}

func (m *moveComponent) Update(deltaTime float32) {
	if !math.NearZero(m.angularSpeed) {
		rot := m.GetOwner().GetRotation()
		angle := m.angularSpeed * deltaTime
		// Create quaternion for incremental rotation
		// (Rotate about up axis)
		inc := math.NewQuaternionFromVec(math.Vector3UnitZ, angle)
		rot = rot.Concatenate(inc)
		m.GetOwner().SetRotation(rot)
	}

	if !math.NearZero(m.forwardSpeed) {
		pos := m.GetOwner().GetPosition()
		forward := m.GetOwner().GetForward()
		pos = pos.Add(forward.MulScalar(m.forwardSpeed * deltaTime))
		m.GetOwner().SetPosition(pos)
	}
}

func (m *moveComponent) GetAngularSpeed() float32 {
	return m.angularSpeed
}

func (m *moveComponent) SetAngularSpeed(speed float32) {
	m.angularSpeed = speed
}

func (m *moveComponent) GetForwardSpeed() float32 {
	return m.forwardSpeed
}

func (m *moveComponent) SetForwardSpeed(speed float32) {
	m.forwardSpeed = speed
}

func (m *moveComponent) Destroy() {
	owner := m.GetOwner()
	owner.RemoveComponent(m)
}
