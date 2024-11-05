package chapter05

import (
	"slices"

	"github.com/ishtaka/go-game-programming/chapter05/math"
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

	ComputeWorldTransform()
	GetWorldTransform() math.Matrix4

	GetState() State
	SetState(s State)

	GetGame() *Game

	AddComponent(c Component)
	RemoveComponent(c Component)

	Destroy()           // must override if embedded in a struct
	DestroyComponents() // must be called in Destroy if embedded in a struct
}

type actor struct {
	// Actor's state
	state State

	// Transform
	worldTransform          math.Matrix4
	position                math.Vector2
	scale                   float32
	rotation                math.Angle
	recomputeWorldTransform bool

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
		scale:                   1,
		rotation:                0,
		recomputeWorldTransform: true,
		game:                    game,
	}

	return a
}

func (a *actor) Update(deltaTime float32) {
	if a.state == Active {
		a.ComputeWorldTransform()

		a.UpdateComponents(deltaTime)
		a.UpdateActor(deltaTime)

		a.ComputeWorldTransform()
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
	a.recomputeWorldTransform = true
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
	a.recomputeWorldTransform = true
}

func (a *actor) GetForward() math.Vector2 {
	return math.Vector2{
		X: math.Cos(a.rotation),
		Y: math.Sin(a.rotation),
	}
}

func (a *actor) ComputeWorldTransform() {
	if a.recomputeWorldTransform {
		a.recomputeWorldTransform = false
		// Scale, then rotate, then translate
		scale := math.Matrix4CreateScale(a.scale, a.scale, 1)
		rotation := math.Matrix4CreateRotationZ(a.rotation)
		translation := math.Matrix4CreateTranslation(math.Vector3{X: a.position.X, Y: a.position.Y, Z: 0.0})
		a.worldTransform = scale.Mul(rotation).Mul(translation)

		// Inform components world transform updated
		for _, c := range a.components {
			c.OnUpdateWorldTransform()
		}
	}
}

func (a *actor) GetWorldTransform() math.Matrix4 {
	return a.worldTransform
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
