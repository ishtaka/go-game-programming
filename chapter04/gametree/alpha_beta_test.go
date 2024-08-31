package gametree

import (
	"fmt"
	"testing"
)

func TestAlphaBetaDecide(t *testing.T) {
	state := GameState{}
	state.Board[0][0] = O
	state.Board[0][1] = Empty
	state.Board[0][2] = X
	state.Board[1][0] = X
	state.Board[1][1] = O
	state.Board[1][2] = O
	state.Board[2][0] = X
	state.Board[2][1] = Empty
	state.Board[2][2] = Empty
	t.Log(fmt.Sprintf("\n%s", state))

	root := NewNode(state)
	root.GenState(true)

	choice := AlphaBetaDecide(root)
	if choice == nil {
		t.Error("choice is nil")
		return
	}
	t.Log(fmt.Sprintf("\n%s", choice.State))
}
