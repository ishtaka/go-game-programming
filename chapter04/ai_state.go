package chapter04

import "github.com/veandco/go-sdl2/sdl"

type AIState interface {
	Update(deltaTime float32)
	OnEnter()
	OnExit()
	GetName() string
}

type AIPatrol struct {
	owner AIComponent
}

func (a *AIPatrol) Update(deltaTime float32) {
	sdl.Log("Updating %s state\n", a.GetName())
	dead := true
	if dead {
		a.owner.ChangeState("Death")
	}
}

func (a *AIPatrol) OnEnter() {
	sdl.Log("Entering %s state\n", a.GetName())
}

func (a *AIPatrol) OnExit() {
	sdl.Log("Exiting %s state\n", a.GetName())
}

func (a *AIPatrol) GetName() string {
	return "Patrol"
}

type AIDeath struct{}

func (a *AIDeath) Update(deltaTime float32) {
	sdl.Log("Updating %s state\n", a.GetName())
}

func (a *AIDeath) OnEnter() {
	sdl.Log("Entering %s state\n", a.GetName())
}

func (a *AIDeath) OnExit() {
	sdl.Log("Exiting %s state\n", a.GetName())
}

func (a *AIDeath) GetName() string {
	return "Death"
}

type AIAttack struct{}

func (a *AIAttack) Update(deltaTime float32) {
	sdl.Log("Updating %s state\n", a.GetName())
}

func (a *AIAttack) OnEnter() {
	sdl.Log("Entering %s state\n", a.GetName())
}

func (a *AIAttack) OnExit() {
	sdl.Log("Exiting %s state\n", a.GetName())
}

func (a *AIAttack) GetName() string {
	return "Attack"
}
