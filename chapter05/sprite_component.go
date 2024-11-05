package chapter05

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ishtaka/go-game-programming/chapter05/math"
)

const DefaultDrawOrder = 100

type Sprite interface {
	Component
	Draw(shader *Shader)
	SetTexture(tex *Texture)
	GetDrawOrder() int
	GetTexWidth() int32
	GetTexHeight() int32
}

type SpriteComponent struct {
	Component
	texture   *Texture
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

func (s *SpriteComponent) Draw(shader *Shader) {
	if s.texture != nil {
		// Scale the quad by the width/height of texture
		scaleMat := math.Matrix4CreateScale(float32(s.texWidth), float32(s.texHeight), 1.0)
		world := scaleMat.Mul(s.GetOwner().GetWorldTransform())

		// Set world transform
		shader.SetMatrixUniform("uWorldTransform", &world)

		// Set current texture
		s.texture.SetActive()

		// Draw quad
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
	}
}

func (s *SpriteComponent) SetTexture(tex *Texture) {
	s.texture = tex
	s.texWidth = tex.Width()
	s.texHeight = tex.Height()
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
