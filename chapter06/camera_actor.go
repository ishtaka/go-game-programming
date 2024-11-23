package chapter06

import (
	"github.com/ishtaka/go-game-programming/chapter06/math"
	"github.com/veandco/go-sdl2/sdl"
)

type CameraActor struct {
	Actor
	moveComp MoveComponent
}

func NewCameraActor(game *Game) *CameraActor {
	a := NewActor(game)
	m := NewMoveComponent(a, DefaultUpdateOrder)
	a.AddComponent(m)

	return &CameraActor{
		Actor:    a,
		moveComp: m,
	}
}

func (c *CameraActor) Update(deltaTime float32) {
	if c.GetState() == Active {
		c.ComputeWorldTransform()

		c.UpdateComponents(deltaTime)
		c.UpdateActor(deltaTime)

		c.ComputeWorldTransform()
	}
}

func (c *CameraActor) UpdateActor(deltaTime float32) {
	c.Actor.UpdateActor(deltaTime)

	cameraPos := c.GetPosition()
	target := cameraPos.Add(c.GetForward().MulScalar(100.0))
	up := math.Vector3UnitZ

	view := math.Matrix4CreateLookAt(cameraPos, target, up)
	c.GetGame().GetRenderer().SetViewMatrix(view)
}

func (c *CameraActor) ProcessInput(keyState []uint8) {
	if c.GetState() == Active {
		c.Actor.ProcessInput(keyState)
		c.ActorInput(keyState)
	}
}

func (c *CameraActor) ActorInput(keyState []uint8) {
	var forwardSpeed, angularSpeed float32

	// wasd movement
	if keyState[sdl.SCANCODE_W] != 0 {
		forwardSpeed += 300.0
	}
	if keyState[sdl.SCANCODE_S] != 0 {
		forwardSpeed -= 300.0
	}
	if keyState[sdl.SCANCODE_A] != 0 {
		angularSpeed -= math.TwoPi
	}
	if keyState[sdl.SCANCODE_D] != 0 {
		angularSpeed += math.TwoPi
	}

	c.moveComp.SetForwardSpeed(forwardSpeed)
	c.moveComp.SetAngularSpeed(angularSpeed)
}

func (c *CameraActor) Destroy() {
	c.GetGame().RemoveActor(c)
	c.DestroyComponents()
}
