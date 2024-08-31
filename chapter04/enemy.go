package chapter04

type Enemy struct {
	Actor
	circle CircleComponent
}

func NewEnemy(game *Game, drawOrder int) *Enemy {
	e := &Enemy{
		Actor: NewActor(game),
	}

	sc := NewSpriteComponent(e, drawOrder)
	sc.SetTexture(game.GetTexture("Assets/Airplane.png"))
	game.AddSprite(sc)
	e.AddComponent(sc)

	// Set position at start tile
	e.SetPosition(game.GetGrid().GetStartTile().GetPosition())

	nc := NewNavComponent(e, DefaultUpdateOrder)
	nc.SetForwardSpeed(150)
	nc.StartPath(game.GetGrid().GetStartTile())
	e.AddComponent(nc)

	// Set up the circle for collision
	e.circle = NewCircleComponent(e, DefaultUpdateOrder)
	e.circle.SetRadius(25.0)
	e.AddComponent(e.circle)

	game.AddEnemy(e)
	game.AddActor(e)

	return e
}

func (s *Enemy) Update(deltaTime float32) {
	if s.GetState() == Active {
		s.Actor.Update(deltaTime)
		s.UpdateActor(deltaTime)
	}
}

func (s *Enemy) UpdateActor(deltaTime float32) {
	s.Actor.UpdateActor(deltaTime)

	// Am I near the end tile?
	diff := s.GetPosition().Sub(s.GetGame().GetGrid().GetEndTile().GetPosition())
	if diff.Length() <= 10 {
		s.SetState(Dead)
	}
}

func (s *Enemy) GetCircle() CircleComponent {
	return s.circle
}

func (s *Enemy) Destroy() {
	s.GetGame().RemoveActor(s)
	s.GetGame().RemoveEnemy(s)
	s.DestroyComponents()
}
