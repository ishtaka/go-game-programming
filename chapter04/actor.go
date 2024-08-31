package chapter04

import (
	"slices"

	"github.com/ishtaka/go-game-programming/chapter04/math"
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
	// Update function called from Game
	Update(deltaTime float32)
	// UpdateComponents updates all components attached to the actor
	UpdateComponents(deltaTime float32)
	// UpdateActor any actor-specific update code (overridable)
	UpdateActor(deltaTime float32)

	// ProcessInput function called from Game (not overridable)
	ProcessInput(keyState []uint8)
	// ActorInput any actor-specific input code (overridable)
	ActorInput(keyState []uint8)

	GetPosition() math.Vector2
	SetPosition(v math.Vector2)
	GetScale() float32
	SetScale(s float32)
	GetRotation() math.Angle
	SetRotation(r math.Angle)
	GetForward() math.Vector2

	GetState() State
	SetState(s State)

	GetGame() *Game

	AddComponent(c Component)
	RemoveComponent(c Component)

	Destroy()           // must override if embedded in a struct
	DestroyComponents() // must be called in Destroy if embedded in a struct
}

type actor struct {
	state      State
	position   math.Vector2
	scale      float32
	rotation   math.Angle
	components []Component
	game       *Game
}

func NewActor(game *Game) Actor {
	a := &actor{
		state: Active,
		position: math.Vector2{
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

func (a *actor) ProcessInput(keyState []uint8) {
	if a.state == Active {
		// first process input for components
		for _, c := range a.components {
			c.ProcessInput(keyState)
		}

		a.ActorInput(keyState)
	}
}

func (a *actor) ActorInput(keyState []uint8) {}

func (a *actor) GetPosition() math.Vector2 {
	return a.position
}

func (a *actor) SetPosition(v math.Vector2) {
	a.position = v
}

func (a *actor) GetScale() float32 {
	return a.scale
}

func (a *actor) SetScale(s float32) {
	a.scale = s
}

func (a *actor) GetRotation() math.Angle {
	return a.rotation
}

func (a *actor) SetRotation(r math.Angle) {
	a.rotation = r
}

func (a *actor) GetForward() math.Vector2 {
	return math.Vector2{
		X: math.Cos(a.rotation),
		Y: -math.Sin(a.rotation), // invert Y because SDL uses top-left origin.
	}
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
