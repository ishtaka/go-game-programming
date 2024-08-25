package chapter03

type Laser struct {
	Actor
	circle     CircleComponent
	deathTimer float32
}

func NewLaser(game *Game, drawOrder int) *Laser {
	l := &Laser{
		Actor: NewActor(game),
	}

	// create a sprite component
	sc := NewSpriteComponent(l, drawOrder)
	sc.SetTexture(game.GetTexture("Assets/Laser.png"))
	game.AddSprite(sc)
	l.AddComponent(sc)

	// create a move component, and set a forward speed
	mc := NewMoveComponent(l, DefaultUpdateOrder)
	mc.SetForwardSpeed(800)
	l.AddComponent(mc)

	// create a circle component
	l.circle = NewCircleComponent(l, DefaultUpdateOrder)
	l.circle.SetRadius(11)
	l.AddComponent(l.circle)

	game.AddActor(l)

	return l
}

func (l *Laser) Update(deltaTime float32) {
	if l.GetState() == Active {
		l.Actor.Update(deltaTime)
		l.UpdateActor(deltaTime)
	}

}

func (l *Laser) UpdateActor(deltaTime float32) {
	l.deathTimer += deltaTime
	if l.deathTimer <= 0.0 {
		l.SetState(Dead)
		return
	}

	// Do we intersect with an asteroid?
	for _, ast := range l.GetGame().GetAsteroids() {
		// The first asteroid we intersect with,
		// set ourselves and the asteroid to dead
		if Intersect(l.circle, ast.GetCircle()) {
			l.SetState(Dead)
			ast.SetState(Dead)
			break
		}
	}
}

func (l *Laser) Destroy() {
	l.GetGame().RemoveActor(l)
	l.DestroyComponents()
}
