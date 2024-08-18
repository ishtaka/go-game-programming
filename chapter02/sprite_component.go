package chapter02

import "github.com/veandco/go-sdl2/sdl"

const DefaultDrawOrder = 100

type Sprite interface {
	Component
	Draw(renderer *sdl.Renderer)
	GetDrawOrder() int
	SetTexture(tex *sdl.Texture)
	GetTexWidth() int32
	GetTexHeight() int32
}

type SpriteComponent struct {
	Component
	tex       *sdl.Texture
	drawOrder int
	texWidth  int32
	texHeight int32
}

func NewSpriteComponent(owner Actor, drawOrder int) *SpriteComponent {
	c := NewComponent(owner, DefaultUpdateOrder)
	sc := &SpriteComponent{
		Component: c,
		drawOrder: drawOrder,
	}

	return sc
}

func (s *SpriteComponent) Draw(renderer *sdl.Renderer) {
	if s.tex != nil {
		r := sdl.Rect{}
		owner := s.GetOwner()

		// Scale the width/height by owner's scale
		r.W = int32(float32(s.texWidth) * owner.GetScale())
		r.H = int32(float32(s.texHeight) * owner.GetScale())

		// Center the rectangle around the position of the owner
		pos := owner.GetPosition()
		r.X = int32(pos.X - float32(r.W)/2)
		r.Y = int32(pos.Y - float32(r.H)/2)

		// Draw (have to convert angle from radians to degrees, and clockwise to counter)
		if err := renderer.CopyEx(s.tex, nil, &r, owner.GetRotation().Degrees(), nil, sdl.FLIP_NONE); err != nil {
			sdl.Log("failed to copy texture: %s\n", err)
		}
	}
}

func (s *SpriteComponent) SetTexture(tex *sdl.Texture) {
	s.tex = tex
	_, _, width, height, _ := tex.Query()
	// Set width/height
	s.texWidth = width
	s.texHeight = height
}

func (s *SpriteComponent) GetDrawOrder() int {
	return s.drawOrder
}

func (s *SpriteComponent) GetTexWidth() int32 {
	return s.texWidth
}

func (s *SpriteComponent) GetTexHeight() int32 {
	return s.texHeight
}

func (s *SpriteComponent) Destroy() {
	owner := s.GetOwner()
	owner.RemoveComponent(s)
	owner.GetGame().RemoveSprite(s)
}
