package chapter03

import (
	"github.com/ishtaka/go-game-programming/chapter03/math"
	"github.com/ishtaka/go-game-programming/chapter03/math/rand"
)

type Asteroid struct {
	Actor
	circle CircleComponent
}

func NewAsteroid(game *Game, drawOrder int) *Asteroid {
	s := &Asteroid{
		Actor: NewActor(game),
	}

	randPos := rand.GetVector2(math.ZeroVector2, math.Vector2{X: 1024, Y: 768})
	s.SetPosition(randPos)

	randAngle := math.Angle(rand.GetFloatRange(0, math.TwoPi))
	s.SetRotation(randAngle)

	// create a sprite component
	sc := NewSpriteComponent(s, drawOrder)
	sc.SetTexture(game.GetTexture("Assets/Asteroid.png"))
	game.AddSprite(sc)
	s.AddComponent(sc)

	// create a move component
	mc := NewMoveComponent(s, DefaultUpdateOrder)
	mc.SetForwardSpeed(150)
	s.AddComponent(mc)

	// create a circle component
	s.circle = NewCircleComponent(s, DefaultUpdateOrder)
	s.circle.SetRadius(40)
	s.AddComponent(s.circle)

	game.AddActor(s)
	game.AddAsteroid(s)

	return s
}

func (a *Asteroid) GetCircle() CircleComponent {
	return a.circle
}

func (a *Asteroid) Destroy() {
	a.GetGame().RemoveActor(a)
	a.GetGame().RemoveAsteroid(a)
	a.DestroyComponents()
}
