package gametree

import (
	"strings"
)

type SquareState int

const (
	Empty SquareState = iota
	X
	O
)

func (s SquareState) String() string {
	return [...]string{
		" ",
		"X",
		"O",
	}[s]
}

type GameState struct {
	// (For tic-tac-toe, array of board)
	Board [3][3]SquareState
}

func (g GameState) String() string {
	var sb strings.Builder
	for i := 0; i < 3; i++ {
		row := make([]string, 3)
		for j, state := range g.Board[i] {
			row[j] = state.String()
		}

		sb.WriteString(strings.Join(row[:], "|") + "\n")
		if i < 2 {
			sb.WriteString("-+-+-\n")
		}
	}
	return sb.String()
}

func (g GameState) GetScore() float64 {
	// Are any of the rows the same?
	for i := 0; i < 3; i++ {
		same := true
		v := g.Board[i][0]
		for j := 1; j < 3; j++ {
			if g.Board[i][j] != v {
				same = false
			}
		}

		if same {
			if v == X {
				return 1.0
			} else {
				return -1.0
			}
		}
	}

	// Are any of the columns the same?
	for j := 0; j < 3; j++ {
		same := true
		v := g.Board[0][j]
		for i := 1; i < 3; i++ {
			if g.Board[i][j] != v {
				same = false
			}
		}

		if same {
			if v == X {
				return 1.0
			} else {
				return -1.0
			}
		}
	}

	// What about diagonals?
	if g.Board[0][0] == g.Board[1][1] && g.Board[1][1] == g.Board[2][2] ||
		g.Board[2][0] == g.Board[1][1] && g.Board[1][1] == g.Board[0][2] {
		if g.Board[1][1] == X {
			return 1.0
		} else {
			return -1.0
		}
	}

	// We tied
	return 0.0
}

type Node struct {
	Children []*Node
	State    GameState
}

func NewNode(state GameState) *Node {
	return &Node{
		Children: make([]*Node, 0, 3*3),
		State:    state,
	}
}

func (n *Node) GenState(xPlayer bool) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if n.State.Board[i][j] == Empty {
				node := NewNode(n.State)
				n.Children = append(n.Children, node)
				if xPlayer {
					node.State.Board[i][j] = X
				} else {
					node.State.Board[i][j] = O
				}
				node.GenState(!xPlayer)
			}
		}
	}
}
