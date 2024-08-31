package chapter04

import "github.com/ishtaka/go-game-programming/chapter04/math"

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
		rot = rot.Add(m.angularSpeed * deltaTime)
		m.GetOwner().SetRotation(rot)
	}

	if !math.NearZero(m.forwardSpeed) {
		pos := m.GetOwner().GetPosition()
		forward := m.GetOwner().GetForward()
		pos = pos.Add(forward.MulScalar(m.forwardSpeed * deltaTime))

		// (screen wrapping code only asteroids)
		if pos.X < 0 {
			pos.X = 1022
		} else if pos.X > 1024 {
			pos.X = 2
		}

		if pos.Y < 0 {
			pos.Y = 766
		} else if pos.Y > 768 {
			pos.Y = 2
		}

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
