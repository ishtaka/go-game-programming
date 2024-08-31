package chapter04

import "github.com/ishtaka/go-game-programming/chapter04/math"

type CircleComponent interface {
	Component
	GetRadius() float32
	SetRadius(radius float32)
	GetCenter() math.Vector2
}

type circleComponent struct {
	Component
	radius float32
}

func NewCircleComponent(owner Actor, updateOrder int) CircleComponent {
	c := NewComponent(owner, updateOrder)
	cc := &circleComponent{
		Component: c,
	}

	return cc
}

func (c *circleComponent) GetRadius() float32 {
	return c.GetOwner().GetScale() * c.radius
}

func (c *circleComponent) SetRadius(radius float32) {
	c.radius = radius
}

func (c *circleComponent) GetCenter() math.Vector2 {
	return c.GetOwner().GetPosition()
}

func (c *circleComponent) Destroy() {
	c.GetOwner().RemoveComponent(c)
}

func Intersect(a, b CircleComponent) bool {
	// Calculate distance squared
	diff := a.GetCenter().Sub(b.GetCenter())
	distSq := diff.LengthSq()

	// Calculate sum of radii squared
	radiiSq := a.GetRadius() + b.GetRadius()
	radiiSq *= radiiSq

	return distSq <= radiiSq
}
