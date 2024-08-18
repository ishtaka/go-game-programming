package chapter02

import "github.com/veandco/go-sdl2/sdl"

type BGTexture struct {
	tex    *sdl.Texture
	offset Vector2
}

type BgSpriteComponent struct {
	Sprite
	textures    []*BGTexture
	screenSize  Vector2
	scrollSpeed float32
}

func NewBgSpriteComponent(owner Actor, drawOrder int) *BgSpriteComponent {
	s := NewSpriteComponent(owner, drawOrder)
	sc := &BgSpriteComponent{
		Sprite:      s,
		screenSize:  Vector2{1024, 768},
		scrollSpeed: 0,
	}

	return sc
}

func (b *BgSpriteComponent) Update(deltaTime float32) {
	b.Sprite.Update(deltaTime)

	for _, bg := range b.textures {
		// Update the x offset
		bg.offset.X += b.scrollSpeed * deltaTime
		// If this is completely off the screen, reset offset to
		// the right of the last bg texture
		if bg.offset.X < -b.screenSize.X {
			bg.offset.X = float32((len(b.textures))-1)*b.screenSize.X - 1
		}
	}
}

func (b *BgSpriteComponent) Draw(renderer *sdl.Renderer) {
	// Draw each background texture
	for _, bg := range b.textures {
		r := sdl.Rect{}
		// Assume screen size dimensions
		r.W = int32(b.screenSize.X)
		r.H = int32(b.screenSize.Y)
		// Center the rectangle around the position of the owner
		pos := b.GetOwner().GetPosition()
		r.X = int32(pos.X - float32(r.W)/2 + bg.offset.X)
		r.Y = int32(pos.Y - float32(r.H)/2 + bg.offset.Y)

		_ = renderer.Copy(bg.tex, nil, &r)
	}
}

func (b *BgSpriteComponent) SetBGTextures(textures []*sdl.Texture) {
	count := 0
	b.textures = make([]*BGTexture, len(textures))
	for i, tex := range textures {
		bg := &BGTexture{
			tex: tex,
			offset: Vector2{
				X: float32(count) * b.screenSize.X,
				Y: 0,
			},
		}
		b.textures[i] = bg
		count++
	}
}

func (b *BgSpriteComponent) SetScreenSize(size Vector2) {
	b.screenSize = size
}

func (b *BgSpriteComponent) SetScrollSpeed(speed float32) {
	b.scrollSpeed = speed
}

func (b *BgSpriteComponent) GetScrollSpeed() float32 {
	return b.scrollSpeed
}

func (b *BgSpriteComponent) Destroy() {
	owner := b.GetOwner()
	owner.RemoveComponent(b)
	owner.GetGame().RemoveSprite(b)
}
