package chapter04

import "github.com/veandco/go-sdl2/sdl"

type AIComponent struct {
	Component
	stateMap     map[string]AIState
	currentState AIState
}

func NewAIComponent(owner Actor, updateOrder int) *AIComponent {
	c := NewComponent(owner, updateOrder)
	ac := &AIComponent{
		Component: c,
		stateMap:  make(map[string]AIState),
	}

	return ac
}

func (a *AIComponent) Update(deltaTime float32) {
	if a.currentState != nil {
		a.currentState.Update(deltaTime)
	}
}

func (a *AIComponent) ChangeState(name string) {
	// First exit the current state
	if a.currentState != nil {
		a.currentState.OnExit()
	}

	// Try to find the new state from the map
	if state, ok := a.stateMap[name]; ok {
		a.currentState = state
		// We're entering the new state
		a.currentState.OnEnter()
	} else {
		sdl.Log("Could not find AIState %s in state map\n", name)
		a.currentState = nil
	}
}

func (a *AIComponent) RegisterState(state AIState) {
	a.stateMap[state.GetName()] = state
}
