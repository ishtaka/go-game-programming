package chapter02

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Ship struct {
	Actor
	rightSpeed float32
	downSpeed  float32
}

func NewShip(game *Game, drawOrder int) *Ship {
	s := &Ship{
		Actor: NewActor(game),
	}

	// Create a sprite component
	asc := NewAnimSpriteComponent(s, drawOrder)
	anims := []*sdl.Texture{
		game.GetTexture("Assets/Ship01.png"),
		game.GetTexture("Assets/Ship02.png"),
		game.GetTexture("Assets/Ship03.png"),
		game.GetTexture("Assets/Ship04.png"),
	}
	asc.SetAnimTextures(anims)
	game.AddSprite(asc)
	s.Actor.AddComponent(asc)

	return s
}

func (s *Ship) Update(deltaTime float32) {
	if s.GetState() == Active {
		s.Actor.UpdateComponents(deltaTime)
		s.UpdateActor(deltaTime)
	}
}

func (s *Ship) UpdateActor(deltaTime float32) {
	s.Actor.UpdateActor(deltaTime)
	// Update position based on speeds and delta time
	pos := s.GetPosition()
	pos.X += s.rightSpeed * deltaTime
	pos.Y += s.downSpeed * deltaTime
	// Restrict position to left half of screen
	if pos.X < 25 {
		pos.X = 25
	} else if pos.X > 500 {
		pos.X = 500
	}

	if pos.Y < 25 {
		pos.Y = 25
	} else if pos.Y > 743 {
		pos.Y = 743
	}
	s.SetPosition(pos)
}

func (s *Ship) ProcessKeyboard(keyState []uint8) {
	s.rightSpeed = 0
	s.downSpeed = 0
	// right/left
	if keyState[sdl.SCANCODE_D] != 0 {
		s.rightSpeed += 250
	}
	if keyState[sdl.SCANCODE_A] != 0 {
		s.rightSpeed -= 250
	}
	// up/down
	if keyState[sdl.SCANCODE_S] != 0 {
		s.downSpeed += 300
	}
	if keyState[sdl.SCANCODE_W] != 0 {
		s.downSpeed -= 300
	}
}

func (s *Ship) GetRightSpeed() float32 {
	return s.rightSpeed
}

func (s *Ship) GetDownSpeed() float32 {
	return s.downSpeed
}

func (s *Ship) Destroy() {
	s.GetGame().RemoveActor(s)
	s.DestroyComponents()
}
