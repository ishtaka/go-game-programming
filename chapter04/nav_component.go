package chapter04

import "github.com/ishtaka/go-game-programming/chapter04/math"

type NavComponent interface {
	MoveComponent
	StartPath(start *Tile)
	TurnTo(pos math.Vector2)
}

type navComponent struct {
	MoveComponent
	nextNode *Tile
}

func NewNavComponent(owner Actor, updateOrder int) NavComponent {
	m := NewMoveComponent(owner, updateOrder)
	n := &navComponent{
		MoveComponent: m,
	}

	return n
}

func (m *navComponent) Update(deltaTime float32) {
	if m.nextNode != nil {
		diff := m.GetOwner().GetPosition().Sub(m.nextNode.GetPosition())
		if diff.Length() < 2.0 {
			m.nextNode = m.nextNode.GetParent()
			m.TurnTo(m.nextNode.GetPosition())
		}
	}

	m.MoveComponent.Update(deltaTime)
}

func (m *navComponent) StartPath(start *Tile) {
	m.nextNode = start.GetParent()
	m.TurnTo(m.nextNode.GetPosition())
}

func (m *navComponent) TurnTo(pos math.Vector2) {
	// Vector from me to pos
	dir := pos.Sub(m.GetOwner().GetPosition())
	// New angle is just atan2 of this dir vector
	// (Negate y because +y is down on screen)
	angle := math.Atan2(-dir.Y, dir.X)
	m.GetOwner().SetRotation(angle)
}

func (m *navComponent) Destroy() {
	owner := m.GetOwner()
	owner.RemoveComponent(m)
}
