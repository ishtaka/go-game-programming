package chapter03

import (
	"github.com/ishtaka/go-game-programming/chapter03/math"
	"github.com/veandco/go-sdl2/sdl"
)

type Ship struct {
	Actor
	laserCoolDown float32
}

func NewShip(game *Game, drawOrder int) *Ship {
	s := &Ship{
		Actor: NewActor(game),
	}

	// create a sprite component
	sc := NewSpriteComponent(s, drawOrder)
	sc.SetTexture(game.GetTexture("Assets/Ship.png"))
	game.AddSprite(sc)
	s.AddComponent(sc)

	// create an input component and set keys/speed
	ic := NewInputComponent(s, DefaultUpdateOrder)
	ic.SetForwardKey(sdl.SCANCODE_W)
	ic.SetBackKey(sdl.SCANCODE_S)
	ic.SetClockwiseKey(sdl.SCANCODE_A)
	ic.SetCounterClockwiseKey(sdl.SCANCODE_D)
	ic.SetMaxForwardSpeed(300)
	ic.SetMaxAngularSpeed(math.TwoPi)
	s.AddComponent(ic)

	game.AddActor(s)

	return s
}

func (s *Ship) Update(deltaTime float32) {
	if s.GetState() == Active {
		s.Actor.Update(deltaTime)
		s.UpdateActor(deltaTime)
	}
}

func (s *Ship) UpdateActor(deltaTime float32) {
	s.laserCoolDown -= deltaTime
}

func (s *Ship) ProcessInput(keyState []uint8) {
	s.Actor.ProcessInput(keyState)
	if s.GetState() == Active {
		s.ActorInput(keyState)
	}
}

func (s *Ship) ActorInput(keyState []uint8) {
	if keyState[sdl.SCANCODE_SPACE] != 0 && s.laserCoolDown <= 0.0 {
		// Create a laser and set its position/rotation to mine
		laser := NewLaser(s.GetGame(), DefaultUpdateOrder)
		laser.SetPosition(s.GetPosition())
		laser.SetRotation(s.GetRotation())

		// Reset laser cooldown (half second)
		s.laserCoolDown = 0.5
	}
}

func (s *Ship) Destroy() {
	s.GetGame().RemoveActor(s)
	s.DestroyComponents()
}
