package chapter06

import (
	"encoding/json"

	"github.com/ishtaka/go-game-programming/chapter06/math"
	"github.com/veandco/go-sdl2/sdl"
)

type meshData struct {
	Version       int          `json:"version"`
	VertexFormat  string       `json:"vertexformat"`
	Shader        string       `json:"shader"`
	Textures      []string     `json:"textures"`
	SpecularPower float32      `json:"specularPower"`
	Vertices      [][8]float32 `json:"vertices"`
	Indices       [][3]uint32  `json:"indices"`
}

type Mesh struct {
	textures    []*Texture
	vertexArray *VertexArray
	shaderName  string
	radius      float32
	specPower   float32
}

func (m *Mesh) Load(fileName string, renderer *Renderer) bool {
	file, err := assets.ReadFile(fileName)
	if err != nil {
		sdl.Log("file not found: Mesh %s %s", fileName, err)
		return false
	}

	var doc meshData
	if err := json.Unmarshal(file, &doc); err != nil {
		sdl.Log("mesh %s is not valid json: %s", fileName, err)
		return false
	}

	if doc.Version != 1 {
		sdl.Log("mesh %s not version 1", fileName)
		return false
	}

	m.shaderName = doc.Shader

	// Skip the vertex format/shader for now
	// (This is changed in a later chapter's code)
	const vertSize int = 8

	textures := doc.Textures
	if len(textures) < 1 {
		sdl.Log("mesh %s has no textures, there should be at least one", fileName)
		return false
	}

	m.specPower = doc.SpecularPower

	for i := 0; i < len(textures); i++ {
		// Is this texture already loaded?
		texName := textures[i]
		tex := renderer.GetTexture(texName)
		if tex == nil {
			// Try loading the texture
			tex = renderer.GetTexture(texName)
			if tex == nil {
				// If it's still null, just use the default texture
				tex = renderer.GetTexture("Assets/Default.png")
			}
		}
		m.textures = append(m.textures, tex)
	}

	// Load in the vertices
	vertsJson := doc.Vertices
	if len(vertsJson) < 1 {
		sdl.Log("mesh %s has no vertices", fileName)
		return false
	}

	vertices := make([]float32, 0, len(vertsJson)*vertSize)
	m.radius = 0.0
	for i := 0; i < len(vertsJson); i++ {
		vert := vertsJson[i]
		if len(vert) != vertSize {
			sdl.Log("unexpected vertex format for %s", fileName)
			return false
		}

		pos := math.Vector3{X: vert[0], Y: vert[1], Z: vert[2]}
		m.radius = math.Max(m.radius, pos.LengthSq())

		// Add the floats
		for j := 0; j < vertSize; j++ {
			vertices = append(vertices, vert[j])
		}
	}

	// We were computing length squared earlier
	m.radius = math.Sqrt(m.radius)

	// Load in the indices
	indJson := doc.Indices
	if len(indJson) < 1 {
		sdl.Log("mesh %s has no indices", fileName)
		return false
	}

	indices := make([]uint32, 0, len(indJson)*3)
	for i := 0; i < len(indJson); i++ {
		ind := indJson[i]
		if len(ind) != 3 {
			sdl.Log("invalid indices for %s", fileName)
			return false
		}

		indices = append(indices, ind[0])
		indices = append(indices, ind[1])
		indices = append(indices, ind[2])
	}

	m.vertexArray = NewVertexArray(vertices, len(vertices), indices, len(indices))

	return true
}

func (m *Mesh) Unload() {
	if m.vertexArray != nil {
		m.vertexArray.Destroy()
	}
}

func (m *Mesh) GetVertexArray() *VertexArray {
	return m.vertexArray
}

func (m *Mesh) GetTexture(index int) *Texture {
	if index < len(m.textures) {
		return m.textures[index]
	}

	return nil
}

func (m *Mesh) ShaderName() string {
	return m.shaderName
}

func (m *Mesh) Radius() float32 {
	return m.radius
}

func (m *Mesh) SpecPower() float32 {
	return m.specPower
}
