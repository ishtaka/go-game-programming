package chapter06

import (
	"fmt"
	"slices"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/ishtaka/go-game-programming/chapter06/math"
)

type DirectionalLight struct {
	// Direction of light
	Direction math.Vector3
	// Diffuse color
	DiffuseColor math.Vector3
	// Specular color
	SpecColor math.Vector3
}

type Renderer struct {
	// Map of textures loaded
	textures map[string]*Texture
	// Map of meshes loaded
	meshes map[string]*Mesh

	// All the sprite components drawn
	sprites []Sprite

	// All mesh components drawn
	meshComps []MeshComponent

	// Game
	game *Game

	// Sprite shader
	spriteShader *Shader
	// Sprite vertex array
	spriteVerts *VertexArray

	// Mesh shader
	meshShader *Shader

	// View/Projection for 3D shaders
	view math.Matrix4
	proj math.Matrix4
	// Width/Height of screen
	screenWidth  float32
	screenHeight float32

	// Lighting data
	ambientLight math.Vector3
	dirLight     *DirectionalLight

	// Window
	window *sdl.Window
	// OpenGL context
	glContext sdl.GLContext
}

func NewRenderer(game *Game) *Renderer {
	return &Renderer{
		textures: make(map[string]*Texture),
		meshes:   make(map[string]*Mesh),
		game:     game,
		dirLight: &DirectionalLight{},
	}
}

func (r *Renderer) Initialize(screenWidth, screenHeight float32) error {
	r.screenWidth = screenWidth
	r.screenHeight = screenHeight

	// Set OpenGL attributes
	// Use the core OpenGL profile
	_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
	// Specify version 3.3
	_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3)
	_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 3)
	// Request a color buffer with 8-bits per RGBA channel
	_ = sdl.GLSetAttribute(sdl.GL_RED_SIZE, 8)
	_ = sdl.GLSetAttribute(sdl.GL_GREEN_SIZE, 8)
	_ = sdl.GLSetAttribute(sdl.GL_BLUE_SIZE, 8)
	_ = sdl.GLSetAttribute(sdl.GL_ALPHA_SIZE, 8)
	_ = sdl.GLSetAttribute(sdl.GL_DEPTH_SIZE, 24)
	// Enable double buffering
	_ = sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)
	// Force OpenGL to use hardware acceleration
	_ = sdl.GLSetAttribute(sdl.GL_ACCELERATED_VISUAL, 1)

	var err error

	r.window, err = sdl.CreateWindow("Chapter 6",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 1024, 768, sdl.WINDOW_OPENGL)
	if err != nil {
		sdl.Log("failed to create window: %s\n", err)
		return err
	}

	// Create an OpenGL context
	r.glContext, err = r.window.GLCreateContext()
	if err != nil {
		sdl.Log("failed to create OpenGL context: %s\n", err)
		return err
	}

	// Initialize OpenGL
	if err = gl.Init(); err != nil {
		sdl.Log("failed to initialize OpenGL: %s\n", err)
		return err
	}

	// On some platforms, GLEW will emit a benign error code,
	// so clear it
	gl.GetError()

	if !r.loadShaders() {
		sdl.Log("Failed to load shaders.")
		return fmt.Errorf("failed to load shaders")
	}

	// Create quad for drawing sprites
	r.createSpriteVerts()

	return nil
}

func (r *Renderer) loadShaders() bool {
	// Create sprite shader
	r.spriteShader = NewShader()
	if !r.spriteShader.Load("sprite.vert", "sprite.frag") {
		return false
	}

	r.spriteShader.SetActive()
	// Set the view-projection matrix
	viewProj := math.Matrix4CreateSimpleViewProj(r.screenWidth, r.screenHeight)
	r.spriteShader.SetMatrixUniform("uViewProj", &viewProj)

	// Create basic mesh shader
	r.meshShader = NewShader()
	if !r.meshShader.Load("phong.vert", "phong.frag") {
		return false
	}

	r.meshShader.SetActive()
	// Set the view-projection matrix
	r.view = math.Matrix4CreateLookAt(math.Vector3Zero, math.Vector3UnitX, math.Vector3UnitZ)
	r.proj = math.Matrix4CreatePerspectiveFOV(70, r.screenWidth, r.screenHeight, 25.0, 10000.0)
	viewProj2 := r.view.Mul(r.proj)
	r.meshShader.SetMatrixUniform("uViewProj", &viewProj2)

	return true
}

func (r *Renderer) createSpriteVerts() {
	vertices := []float32{
		-0.5, 0.5, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, // top left
		0.5, 0.5, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, // top right
		0.5, -0.5, 0.0, 0.0, 0.0, 0.0, 1.0, 1.0, // bottom right
		-0.5, -0.5, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, // bottom left
	}

	indices := []uint32{
		0, 1, 2,
		2, 3, 0,
	}

	r.spriteVerts = NewVertexArray(vertices, 4, indices, 6)
}

func (r *Renderer) Draw() {
	// Set the clear color
	gl.ClearColor(0.86, 0.86, 0.86, 1.0)
	// Clear the color buffer
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	// Draw mesh components
	// Enable depth buffering/disable alpha blend
	gl.Enable(gl.DEPTH_TEST)
	gl.Disable(gl.BLEND)
	// Set the mesh shader active
	r.meshShader.SetActive()
	// Update view-projection matrix
	viewProj := r.view.Mul(r.proj)
	r.meshShader.SetMatrixUniform("uViewProj", &viewProj)
	// Update lighting uniforms
	r.SetLightUniforms(r.meshShader)

	for _, mc := range r.meshComps {
		mc.Draw(r.meshShader)
	}

	// Draw all sprite components
	// Disable depth buffering
	gl.Disable(gl.DEPTH_TEST)
	// Enable alpha blending on the color buffer
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	// Set shader/vao as active
	r.spriteShader.SetActive()
	r.spriteVerts.SetActive()

	for _, s := range r.sprites {
		s.Draw(r.spriteShader)
	}

	// Swap the buffers
	r.window.GLSwap()
}

func (r *Renderer) AddSprite(s Sprite) {
	order := s.GetDrawOrder()
	insertIndex := 0
	for insertIndex < len(r.sprites) {
		if order < r.sprites[insertIndex].GetDrawOrder() {
			break
		}
		insertIndex++
	}

	// Insert at position
	r.sprites = slices.Insert(r.sprites, insertIndex, s)
}

func (r *Renderer) RemoveSprite(s Sprite) {
	r.sprites = slices.DeleteFunc(r.sprites, func(s2 Sprite) bool {
		return s == s2
	})
}

func (r *Renderer) AddMeshComp(mc MeshComponent) {
	r.meshComps = append(r.meshComps, mc)
}

func (r *Renderer) RemoveMeshComp(mc MeshComponent) {
	r.meshComps = slices.DeleteFunc(r.meshComps, func(mc2 MeshComponent) bool {
		return mc == mc2
	})
}

func (r *Renderer) GetTexture(fileName string) *Texture {
	// Is the texture already in the map?
	if tex, ok := r.textures[fileName]; ok {
		return tex
	}

	tex := NewTexture()
	if tex.Load(fileName) {
		r.textures[fileName] = tex
		return tex
	}

	return nil
}

func (r *Renderer) GetMesh(fileName string) *Mesh {
	// Is the mesh already in the map?
	if mesh, ok := r.meshes[fileName]; ok {
		return mesh
	}

	mesh := &Mesh{}
	if mesh.Load(fileName, r) {
		r.meshes[fileName] = mesh
		return mesh
	}

	return nil
}

func (r *Renderer) SetViewMatrix(view math.Matrix4) {
	r.view = view
}

func (r *Renderer) SetAmbientLight(ambient math.Vector3) {
	r.ambientLight = ambient
}

func (r *Renderer) GetDirectionalLight() *DirectionalLight {
	return r.dirLight
}

func (r *Renderer) SetLightUniforms(shader *Shader) {
	// Camera position is from inverted view
	invView := r.view.Invert()
	shader.SetVectorUniform("uCameraPos", invView.GetTranslation())
	// Ambient light
	shader.SetVectorUniform("uAmbientLight", r.ambientLight)
	// Directional light
	shader.SetVectorUniform("uDirLight.mDirection", r.dirLight.Direction)
	shader.SetVectorUniform("uDirLight.mDiffuseColor", r.dirLight.DiffuseColor)
	shader.SetVectorUniform("uDirLight.mSpecColor", r.dirLight.SpecColor)
}

func (r *Renderer) UnloadData() {
	// Destroy textures
	for k, tex := range r.textures {
		tex.Unload()
		delete(r.textures, k)
	}

	// Destroy meshes
	for k, mesh := range r.meshes {
		mesh.Unload()
		delete(r.meshes, k)
	}
}

func (r *Renderer) Shutdown() (err error) {
	defer func() {
		if r.window != nil {
			err = r.window.Destroy()
		}
	}()
	defer func() {
		if r.glContext != nil {
			sdl.GLDeleteContext(r.glContext)
		}
	}()
	defer func() {
		if r.meshShader != nil {
			r.meshShader.Unload()
		}
	}()
	defer func() {
		if r.spriteShader != nil {
			r.spriteShader.Unload()
		}
	}()
	defer func() {
		if r.spriteVerts != nil {
			r.spriteVerts.Destroy()
		}
	}()

	return nil
}
