package chapter02

import (
	"slices"
)

type State int

const (
	Active State = iota
	Paused
	Dead
)

func (s State) String() string {
	return [...]string{
		"Active",
		"Paused",
		"Dead",
	}[s]
}

type Actor interface {
	Update(deltaTime float32)
	UpdateComponents(deltaTime float32)
	UpdateActor(deltaTime float32)
	GetPosition() Vector2
	SetPosition(v Vector2)
	GetScale() float32
	SetScale(s float32)
	GetRotation() Angle
	SetRotation(r float32)
	GetState() State
	SetState(s State)
	GetGame() *Game
	AddComponent(c Component)
	RemoveComponent(c Component)
	Destroy()
	DestroyComponents()
}

type actor struct {
	state      State
	position   Vector2
	scale      float32
	rotation   Angle
	components []Component
	game       *Game
}

func NewActor(game *Game) Actor {
	a := &actor{
		state: Active,
		position: Vector2{
			X: 0,
			Y: 0,
		},
		scale:    1,
		rotation: 0,
		game:     game,
	}

	return a
}

func (a *actor) Update(deltaTime float32) {
	if a.state == Active {
		a.UpdateComponents(deltaTime)
		a.UpdateActor(deltaTime)
	}
}

func (a *actor) UpdateComponents(deltaTime float32) {
	for _, c := range a.components {
		c.Update(deltaTime)
	}
}

func (a *actor) UpdateActor(deltaTime float32) {}

func (a *actor) GetPosition() Vector2 {
	return a.position
}

func (a *actor) SetPosition(v Vector2) {
	a.position = v
}

func (a *actor) GetScale() float32 {
	return a.scale
}

func (a *actor) SetScale(s float32) {
	a.scale = s
}

func (a *actor) GetRotation() Angle {
	return a.rotation
}

func (a *actor) SetRotation(r float32) {
	a.rotation = Angle(r)
}

func (a *actor) GetState() State {
	return a.state
}

func (a *actor) SetState(s State) {
	a.state = s
}

func (a *actor) GetGame() *Game {
	return a.game
}

func (a *actor) AddComponent(c Component) {
	order := c.GetUpdateOrder()
	insertIndex := 0
	for insertIndex < len(a.components) {
		if order < a.components[insertIndex].GetUpdateOrder() {
			break
		}
		insertIndex++
	}

	// Insert at position
	a.components = slices.Insert(a.components, insertIndex, c)
}

func (a *actor) RemoveComponent(c Component) {
	a.components = slices.DeleteFunc(a.components, func(c2 Component) bool {
		return c == c2
	})
}

func (a *actor) Destroy() {
	a.game.RemoveActor(a)
	a.DestroyComponents()
}

func (a *actor) DestroyComponents() {
	for len(a.components) > 0 {
		a.components[0].Destroy()
	}
}
