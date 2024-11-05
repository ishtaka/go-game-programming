package chapter05

import (
	"fmt"
	"slices"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/ishtaka/go-game-programming/chapter05/math"
)

type Game struct {
	window     *sdl.Window
	glContext  sdl.GLContext
	ticksCount uint64
	isRunning  bool

	// Map of textures loaded
	textures map[string]*Texture

	// All the actors in the game
	actors []Actor
	// Any pending actors
	pendingActors []Actor
	// Track if we're updating actors right now
	updatingActors bool

	// All the sprite components drawn
	sprites []Sprite
	// Sprite shader
	spriteShader *Shader
	// Sprite vertex array
	spriteVerts *VertexArray

	ship      *Ship
	asteroids []*Asteroid
}

func NewGame() *Game {
	return &Game{
		ticksCount: 0,
		textures:   make(map[string]*Texture),
		isRunning:  true,
	}
}

func (g *Game) Initialize() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		sdl.Log("unable to initialize SDL: %s\n", err)
		return err
	}

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
	// Enable double buffering
	_ = sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)
	// Force OpenGL to use hardware acceleration
	_ = sdl.GLSetAttribute(sdl.GL_ACCELERATED_VISUAL, 1)

	var err error

	g.window, err = sdl.CreateWindow("Chapter 5", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 1024, 768, sdl.WINDOW_OPENGL)
	if err != nil {
		sdl.Log("failed to create window: %s\n", err)
		return err
	}

	// Create an OpenGL context
	g.glContext, err = g.window.GLCreateContext()
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

	if !g.loadShaders() {
		sdl.Log("Failed to load shaders.")
		return fmt.Errorf("failed to load shaders")
	}

	// Create quad for drawing sprites
	g.createSpriteVerts()

	g.loadData()

	g.ticksCount = sdl.GetTicks64()

	return nil
}

func (g *Game) RunLoop() {
	for g.isRunning {
		g.processInput()
		g.update()
		g.generateOutput()
	}
}

func (g *Game) processInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			g.isRunning = false
			return
		}

		state := sdl.GetKeyboardState()
		if state[sdl.SCANCODE_ESCAPE] != 0 {
			g.isRunning = false
			return
		}

		g.updatingActors = true
		for _, a := range g.actors {
			a.ProcessInput(state)
		}
		g.updatingActors = false
	}
}

func (g *Game) update() {
	// SDL_TICKS_PASSED is not available in go-sdl2
	for {
		if sdl.GetTicks64() > g.ticksCount+16 {
			break
		}
	}

	deltaTime := float32(sdl.GetTicks64()-g.ticksCount) / 1000.0
	if deltaTime > 0.05 {
		deltaTime = 0.05
	}

	g.ticksCount = sdl.GetTicks64()

	// Update all actors
	g.updatingActors = true
	for _, a := range g.actors {
		a.Update(deltaTime)
	}
	g.updatingActors = false

	// Move any pending actors to actors
	for _, pending := range g.pendingActors {
		pending.ComputeWorldTransform()
		g.actors = append(g.actors, pending)
	}
	g.pendingActors = nil

	// Add any dead actors to a temp slice
	deadActors := make([]Actor, 0, len(g.actors))
	for _, a := range g.actors {
		if a.GetState() == Dead {
			deadActors = append(deadActors, a)
		}
	}

	// Delete dead actors (which removes them from actors)
	for _, deadActor := range deadActors {
		deadActor.Destroy()
	}
}

func (g *Game) generateOutput() {
	// Set the clear color
	gl.ClearColor(0.86, 0.86, 0.86, 1.0)
	// Clear the color buffer
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// Draw all sprite components
	// Enable alpha blending on the color buffer
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	g.spriteShader.SetActive()
	g.spriteVerts.SetActive()

	for _, s := range g.sprites {
		s.Draw(g.spriteShader)
	}

	// Swap the buffers
	g.window.GLSwap()
}

func (g *Game) loadShaders() bool {
	g.spriteShader = NewShader()
	if !g.spriteShader.Load("sprite.vert", "sprite.frag") {
		return false
	}

	g.spriteShader.SetActive()
	// Set the view-projection matrix
	viewProj := math.Matrix4CreateSimpleViewProj(1024, 768)
	g.spriteShader.SetMatrixUniform("uViewProj", &viewProj)

	return true
}

func (g *Game) createSpriteVerts() {
	vertices := []float32{
		-0.5, 0.5, 0.0, 0.0, 0.0, // top left
		0.5, 0.5, 0.0, 1.0, 0.0, // top right
		0.5, -0.5, 0.0, 1.0, 1.0, // bottom right
		-0.5, -0.5, 0.0, 0.0, 1.0, // bottom left
	}

	indices := []uint32{
		0, 1, 2,
		2, 3, 0,
	}

	g.spriteVerts = NewVertexArray(vertices, 4, indices, 6)
}

func (g *Game) loadData() {
	// Create player's ship
	g.ship = NewShip(g, DefaultDrawOrder)
	g.ship.SetPosition(math.Vector2{X: 0, Y: 0})
	g.ship.SetRotation(math.Angle(math.PiOver2))

	// Create asteroids
	const numAsteroids = 20
	for range numAsteroids {
		NewAsteroid(g, DefaultDrawOrder)
	}
}

func (g *Game) unloadData() {
	// Delete actors
	for len(g.actors) > 0 {
		g.actors[0].Destroy()
	}

	// Destroy textures
	for k, tex := range g.textures {
		tex.Unload()
		delete(g.textures, k)
	}
}

func (g *Game) GetTexture(fileName string) *Texture {
	// Is the texture already in the map?
	if tex, ok := g.textures[fileName]; ok {
		return tex
	}

	tex := NewTexture()
	if tex.Load(fileName) {
		g.textures[fileName] = tex
		return tex
	}

	return nil
}

func (g *Game) Shutdown() (err error) {
	defer sdl.Quit()
	defer func() {
		if g.window != nil {
			err = g.window.Destroy()
		}
	}()
	defer func() {
		if g.glContext != nil {
			sdl.GLDeleteContext(g.glContext)
		}
	}()
	defer func() {
		g.spriteShader.Unload()
	}()
	defer func() {
		g.spriteVerts.Destroy()
	}()

	g.unloadData()

	return
}

func (g *Game) GetAsteroids() []*Asteroid {
	return g.asteroids
}

func (g *Game) AddAsteroid(ast *Asteroid) {
	g.asteroids = append(g.asteroids, ast)
}

func (g *Game) RemoveAsteroid(ast *Asteroid) {
	g.asteroids = slices.DeleteFunc(g.asteroids, func(ast2 *Asteroid) bool {
		return ast == ast2
	})
}

func (g *Game) AddActor(actor Actor) {
	// If we're updating actors, need to add to pending
	if g.updatingActors {
		g.pendingActors = append(g.pendingActors, actor)
	} else {
		g.actors = append(g.actors, actor)
	}
}

func (g *Game) RemoveActor(actor Actor) {
	g.pendingActors = slices.DeleteFunc(g.pendingActors, func(a Actor) bool {
		return a == actor
	})

	g.actors = slices.DeleteFunc(g.actors, func(a Actor) bool {
		return a == actor
	})
}

func (g *Game) AddSprite(s Sprite) {
	order := s.GetDrawOrder()
	insertIndex := 0
	for insertIndex < len(g.sprites) {
		if order < g.sprites[insertIndex].GetDrawOrder() {
			break
		}
		insertIndex++
	}

	// Insert at position
	g.sprites = slices.Insert(g.sprites, insertIndex, s)
}

func (g *Game) RemoveSprite(s Sprite) {
	g.sprites = slices.DeleteFunc(g.sprites, func(s2 Sprite) bool {
		return s == s2
	})
}
