package chapter02

import "github.com/veandco/go-sdl2/sdl"

type AnimSpriteComponent struct {
	Sprite
	textures  []*sdl.Texture
	currFrame float32
	animFPS   float32
}

func NewAnimSpriteComponent(owner Actor, drawOrder int) *AnimSpriteComponent {
	s := NewSpriteComponent(owner, drawOrder)
	sc := &AnimSpriteComponent{
		Sprite:    s,
		currFrame: 0,
		animFPS:   24,
	}

	return sc
}

func (a *AnimSpriteComponent) Update(deltaTime float32) {
	a.Sprite.Update(deltaTime)

	if len(a.textures) > 0 {
		// Update the current frame based on frame rate and delta time
		a.currFrame += a.animFPS * deltaTime
		// Wrap current frame if needed
		for a.currFrame >= float32(len(a.textures)) {
			a.currFrame -= float32(len(a.textures))
		}

		// Set the current texture
		a.SetTexture(a.textures[int(a.currFrame)])
	}
}

func (a *AnimSpriteComponent) SetAnimTextures(textures []*sdl.Texture) {
	a.textures = textures
	if len(textures) > 0 {
		// Set the current texture to first frame
		a.currFrame = 0
		a.SetTexture(textures[0])
	}
}

func (a *AnimSpriteComponent) GetAnimFPS() float32 {
	return a.animFPS
}

func (a *AnimSpriteComponent) SetAnimFPS(fps float32) {
	a.animFPS = fps
}

func (a *AnimSpriteComponent) Destroy() {
	owner := a.GetOwner()
	owner.RemoveComponent(a)
	owner.GetGame().RemoveSprite(a)
}
