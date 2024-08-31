package chapter04

type Bullet struct {
	Actor
	circle   CircleComponent
	liveTime float32
}

func NewBullet(game *Game, drawOrder int) *Bullet {
	l := &Bullet{
		Actor: NewActor(game),
	}

	// create a sprite component
	sc := NewSpriteComponent(l, drawOrder)
	sc.SetTexture(game.GetTexture("Assets/Projectile.png"))
	game.AddSprite(sc)
	l.AddComponent(sc)

	// create a move component, and set a forward speed
	mc := NewMoveComponent(l, DefaultUpdateOrder)
	mc.SetForwardSpeed(400)
	l.AddComponent(mc)

	// create a circle component
	l.circle = NewCircleComponent(l, DefaultUpdateOrder)
	l.circle.SetRadius(5)
	l.AddComponent(l.circle)

	l.liveTime = 1.0

	game.AddActor(l)

	return l
}

func (l *Bullet) Update(deltaTime float32) {
	if l.GetState() == Active {
		l.Actor.Update(deltaTime)
		l.UpdateActor(deltaTime)
	}

}

func (l *Bullet) UpdateActor(deltaTime float32) {
	l.Actor.UpdateActor(deltaTime)

	// Check for collision vs enemies
	for _, e := range l.GetGame().GetEnemies() {
		if Intersect(l.circle, e.GetCircle()) {
			// We both die on collision
			e.SetState(Dead)
			l.SetState(Dead)
			break
		}
	}

	l.liveTime -= deltaTime
	if l.liveTime <= 0.0 {
		// Time limit hit, die
		l.SetState(Dead)
	}
}

func (l *Bullet) Destroy() {
	l.GetGame().RemoveActor(l)
	l.DestroyComponents()
}
