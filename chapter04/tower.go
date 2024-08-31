package chapter04

import (
	"github.com/ishtaka/go-game-programming/chapter04/math"
)

type Tower struct {
	Actor
	move        MoveComponent
	nextAttack  float32
	attackTime  float32
	attackRange float32
}

func NewTower(game *Game) *Tower {
	t := &Tower{
		Actor:       NewActor(game),
		attackTime:  2.5,
		attackRange: 100,
	}

	sc := NewSpriteComponent(t, DefaultDrawOrder)
	sc.SetTexture(game.GetTexture("Assets/Tower.png"))
	game.AddSprite(sc)
	t.AddComponent(sc)

	t.move = NewMoveComponent(t, DefaultUpdateOrder)
	t.AddComponent(t.move)

	t.nextAttack = t.attackTime

	game.AddActor(t)

	return t
}

func (t *Tower) Update(deltaTime float32) {
	if t.GetState() == Active {
		t.Actor.Update(deltaTime)
		t.UpdateActor(deltaTime)
	}
}

func (t *Tower) UpdateActor(deltaTime float32) {
	t.Actor.UpdateActor(deltaTime)

	t.nextAttack -= deltaTime
	if t.nextAttack <= 0 {
		e := t.GetGame().GetNearestEnemy(t.GetPosition())
		if e != nil {
			// Vector from me to enemy
			dir := e.GetPosition().Sub(t.GetPosition())
			dist := dir.Length()
			if dist < t.attackRange {
				// Rotate to face enemy
				t.SetRotation(math.Atan2(-dir.Y, dir.X))
				// Spawn bullet at tower position facing enemy
				b := NewBullet(t.GetGame(), DefaultDrawOrder)
				b.SetPosition(t.GetPosition())
				b.SetRotation(t.GetRotation())
			}
		}
		t.nextAttack = t.attackTime
	}
}

func (t *Tower) Destroy() {
	t.GetGame().RemoveActor(t)
	t.DestroyComponents()
}
