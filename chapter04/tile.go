package chapter04

type TileState int

const (
	DefaultTile TileState = iota
	PathTile
	StartTile
	BaseTile
)

func (t TileState) String() string {
	return [...]string{
		"Default",
		"Path",
		"Start",
		"Base",
	}[t]
}

type Tile struct {
	Actor
	// For pathfinding
	adjacent               []*Tile
	parent                 *Tile
	f, g, h                float32
	inOpenSet, inClosedSet bool
	blocked                bool

	sprite   Sprite
	state    TileState
	selected bool
}

func NewTile(game *Game) *Tile {
	t := &Tile{
		Actor: NewActor(game),
	}

	t.sprite = NewSpriteComponent(t, DefaultDrawOrder)
	t.updateTexture()
	t.AddComponent(t.sprite)
	game.AddSprite(t.sprite)

	game.AddActor(t)

	return t
}

func (t *Tile) SetTileState(state TileState) {
	t.state = state
	t.updateTexture()
}

func (t *Tile) GetTileState() TileState {
	return t.state
}

func (t *Tile) ToggleSelect() {
	t.selected = !t.selected
	t.updateTexture()
}

func (t *Tile) GetParent() *Tile {
	return t.parent
}

func (t *Tile) updateTexture() {
	text := ""
	switch t.state {
	case StartTile:
		text = "Assets/TileTan.png"
	case BaseTile:
		text = "Assets/TileGreen.png"
	case PathTile:
		if t.selected {
			text = "Assets/TileGreySelected.png"
		} else {
			text = "Assets/TileGrey.png"
		}
	case DefaultTile:
		fallthrough
	default:
		if t.selected {
			text = "Assets/TileBrownSelected.png"
		} else {
			text = "Assets/TileBrown.png"
		}
	}

	t.sprite.SetTexture(t.GetGame().GetTexture(text))
}

func (t *Tile) Destroy() {
	t.GetGame().RemoveActor(t)
	t.DestroyComponents()
}
