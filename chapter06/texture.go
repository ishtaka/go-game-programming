package chapter06

import (
	"embed"
	"image"
	"image/draw"
	"image/png"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)

//go:embed Assets/*
var assets embed.FS

type Texture struct {
	// OpenGL ID of this texture
	textureID uint32
	// Width/height of the texture
	width  int32
	height int32
}

func NewTexture() *Texture {
	return &Texture{}
}

func (t *Texture) Load(fileName string) bool {
	f, err := assets.Open(fileName)
	if err != nil {
		sdl.Log("failed to load image %s: %s", fileName, err)
		return false
	}
	defer func() { _ = f.Close() }()

	img, err := png.Decode(f)
	if err != nil {
		sdl.Log("failed to decode image %s: %s", fileName, err)
		return false
	}

	format := gl.RGB
	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride == rgba.Rect.Size().X*4 {
		format = gl.RGBA
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)

	t.width = int32(rgba.Rect.Size().X)
	t.height = int32(rgba.Rect.Size().Y)

	gl.GenTextures(1, &t.textureID)
	gl.BindTexture(gl.TEXTURE_2D, t.textureID)

	gl.TexImage2D(gl.TEXTURE_2D, 0, int32(format), t.width, t.height, 0, uint32(format),
		gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))

	// Enable bilinear filtering
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	return true
}

func (t *Texture) Unload() {
	gl.DeleteTextures(1, &t.textureID)
}

func (t *Texture) SetActive() {
	gl.BindTexture(gl.TEXTURE_2D, t.textureID)
}

func (t *Texture) Width() int32 {
	return t.width
}

func (t *Texture) Height() int32 {
	return t.height
}
