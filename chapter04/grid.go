package chapter04

import (
	"cmp"
	"slices"

	"github.com/ishtaka/go-game-programming/chapter04/math"
)

type Grid struct {
	Actor
	// Currently selected tile
	selectedTile *Tile
	// 2D vector of tiles in grid
	tiles [][]*Tile
	// Time until next enemy
	nextEnemy float32
	// Rows/columns in grid
	numRows, numCols int
	// Start y position of top left corner
	startY float32
	// Width/height of each tile
	tileSize float32
	// Time between enemies
	enemyTime float32
}

func NewGrid(game *Game) *Grid {
	g := &Grid{
		Actor:     NewActor(game),
		numRows:   7,
		numCols:   16,
		startY:    192,
		tileSize:  64,
		enemyTime: 1.5,
	}

	// 7 rows, 16 columns
	g.tiles = make([][]*Tile, g.numRows)
	for i := 0; i < g.numRows; i++ {
		g.tiles[i] = make([]*Tile, g.numCols)
	}

	// Create tiles
	for i := 0; i < g.numRows; i++ {
		for j := 0; j < g.numCols; j++ {
			g.tiles[i][j] = NewTile(game)
			g.tiles[i][j].SetPosition(math.Vector2{
				X: g.tileSize/2.0 + float32(j)*g.tileSize,
				Y: g.startY + float32(i)*g.tileSize,
			})
		}
	}

	// Set start/end tiles
	g.GetStartTile().SetTileState(StartTile)
	g.GetEndTile().SetTileState(BaseTile)

	// Set up adjacency tiles
	for i := 0; i < g.numRows; i++ {
		for j := 0; j < g.numCols; j++ {
			t := g.tiles[i][j]
			if i > 0 {
				t.adjacent = append(t.adjacent, g.tiles[i-1][j])
			}
			if i < g.numRows-1 {
				t.adjacent = append(t.adjacent, g.tiles[i+1][j])
			}
			if j > 0 {
				t.adjacent = append(t.adjacent, g.tiles[i][j-1])
			}
			if j < g.numCols-1 {
				t.adjacent = append(t.adjacent, g.tiles[i][j+1])
			}
		}
	}

	// Find path (in reverse)
	g.FindPath(g.GetEndTile(), g.GetStartTile())
	g.updatePathTile(g.GetStartTile())

	g.nextEnemy = g.enemyTime

	game.AddActor(g)

	return g
}

func (g *Grid) Update(deltaTime float32) {
	if g.GetState() == Active {
		g.Actor.Update(deltaTime)
		g.UpdateActor(deltaTime)
	}
}

func (g *Grid) UpdateActor(deltaTime float32) {
	g.Actor.UpdateActor(deltaTime)

	// Is it time to spawn a new enemy?
	g.nextEnemy -= deltaTime
	if g.nextEnemy <= 0.0 {
		NewEnemy(g.GetGame(), DefaultDrawOrder)
		g.nextEnemy += g.enemyTime
	}
}

// ProcessClick handles a mouse click at the x/y screen locations.
func (g *Grid) ProcessClick(x, y int) {
	y -= int(g.startY - g.tileSize/2)
	if y >= 0 {
		x /= int(g.tileSize)
		y /= int(g.tileSize)
		if x >= 0 && x < g.numCols && y >= 0 && y < g.numRows {
			g.selectTile(y, x)
		}
	}
}

// FindPath uses A* to find a path.
func (g *Grid) FindPath(start, goal *Tile) bool {
	for i := 0; i < g.numRows; i++ {
		for j := 0; j < g.numCols; j++ {
			g.tiles[i][j].inOpenSet = false
			g.tiles[i][j].inClosedSet = false
		}
	}

	openSet := make([]*Tile, 0, g.numRows*g.numCols)

	// Set current node to start, and add to closed set
	current := start
	current.inClosedSet = true

	for {
		// Add adjacent nodes to open set
		for _, neighbor := range current.adjacent {
			if neighbor.blocked {
				continue
			}

			// Only check nodes that aren't in the closed set
			if !neighbor.inClosedSet {
				if !neighbor.inOpenSet {
					// Not in the open set, so set parent
					neighbor.parent = current
					neighbor.h = neighbor.GetPosition().Sub(goal.GetPosition()).Length()
					// g(x) is the parent's g plus cost of traversing edge
					neighbor.g = current.g + g.tileSize
					neighbor.f = neighbor.g + neighbor.h
					openSet = append(openSet, neighbor)
					neighbor.inOpenSet = true
				} else {
					// Compute g(x) cost if current becomes the parent
					newG := current.g + g.tileSize
					if newG < neighbor.g {
						// Adopt this node
						neighbor.parent = current
						neighbor.g = newG
						// f(x) changes because g(x) changes
						neighbor.f = neighbor.g + neighbor.h
					}
				}
			}
		}

		// If open set is empty, all possible paths are exhausted
		if len(openSet) == 0 {
			break
		}

		// Find the lowest cost node in open set
		lowest := slices.MinFunc(openSet, func(a, b *Tile) int {
			return cmp.Compare(a.f, b.f)
		})

		// Set to current and move from open to closed
		current = lowest
		openSet = slices.DeleteFunc(openSet, func(t *Tile) bool {
			return t == lowest
		})
		current.inOpenSet = false
		current.inClosedSet = true

		if current == goal {
			break
		}
	}

	// Did we find a path?
	return current == goal
}

// BuildTower tries to build a tower.
func (g *Grid) BuildTower() {
	if g.selectedTile != nil && !g.selectedTile.blocked {
		g.selectedTile.blocked = true
		if g.FindPath(g.GetEndTile(), g.GetStartTile()) {
			t := NewTower(g.GetGame())
			t.SetPosition(g.selectedTile.GetPosition())
		} else {
			// This tower would block the path, so don't allow build
			g.selectedTile.blocked = false
			g.FindPath(g.GetEndTile(), g.GetStartTile())
		}
		g.updatePathTile(g.GetStartTile())
	}
}

// GetStartTile return start tile.
func (g *Grid) GetStartTile() *Tile {
	return g.tiles[3][0]
}

// GetEndTile returns end tile.
func (g *Grid) GetEndTile() *Tile {
	return g.tiles[3][15]
}

// selectTile selects a specific tile.
func (g *Grid) selectTile(row, col int) {
	state := g.tiles[row][col].GetTileState()
	if state != StartTile && state != BaseTile {
		// Deselect previous one
		if g.selectedTile != nil {
			g.selectedTile.ToggleSelect()
		}
		g.selectedTile = g.tiles[row][col]
		g.selectedTile.ToggleSelect()
	}
}

// updatePathTile updates textures for tiles on path.
func (g *Grid) updatePathTile(start *Tile) {
	// Reset all tiles to normal (except for start/end)
	for i := 0; i < g.numRows; i++ {
		for j := 0; j < g.numCols; j++ {
			if !(i == 3 && j == 0) && !(i == 3 && j == 15) {
				g.tiles[i][j].SetTileState(DefaultTile)
			}
		}
	}

	t := start.parent
	for t != g.GetEndTile() {
		t.SetTileState(PathTile)
		t = t.parent
	}
}

func (g *Grid) Destroy() {
	g.GetGame().RemoveActor(g)
	g.DestroyComponents()
}
