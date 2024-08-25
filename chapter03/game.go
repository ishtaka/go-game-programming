package chapter03

import (
	"slices"

	"github.com/ishtaka/go-game-programming/chapter03/math"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	window     *sdl.Window
	renderer   *sdl.Renderer
	ticksCount uint64
	isRunning  bool

	textures map[string]*sdl.Texture
	sprites  []Sprite

	actors         []Actor
	pendingActors  []Actor
	updatingActors bool

	ship      *Ship
	asteroids []*Asteroid
}

func NewGame() *Game {
	return &Game{
		ticksCount: 0,
		textures:   make(map[string]*sdl.Texture),
		isRunning:  true,
	}
}

func (g *Game) Initialize() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		sdl.Log("unable to initialize SDL: %s\n", err)
		return err
	}

	var err error

	g.window, err = sdl.CreateWindow("Chapter 3", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 1024, 768, sdl.WINDOW_SHOWN)
	if err != nil {
		sdl.Log("failed to create window: %s\n", err)
		return err
	}

	g.renderer, err = sdl.CreateRenderer(g.window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		sdl.Log("failed to create renderer: %s\n", err)
		return err
	}

	err = img.Init(img.INIT_PNG)
	if err != nil {
		sdl.Log("unable to initialize SDL_image: %s\n", err)
		return err
	}

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
	_ = g.renderer.SetDrawColor(220, 220, 220, 255)

	// Clear the back buffer
	_ = g.renderer.Clear()

	// Draw all sprite components
	for _, sprite := range g.sprites {
		sprite.Draw(g.renderer)
	}

	g.renderer.Present()
}

func (g *Game) loadData() {
	// Create player's ship
	g.ship = NewShip(g, DefaultDrawOrder)
	g.ship.SetPosition(math.Vector2{X: 512, Y: 384})
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
	for _, tex := range g.textures {
		_ = tex.Destroy()
	}
}

func (g *Game) GetTexture(fileName string) *sdl.Texture {
	// Is the texture already in the map?
	if tex, ok := g.textures[fileName]; ok {
		return tex
	}

	// Load from file
	surf, err := img.Load("chapter03/" + fileName)
	if err != nil {
		sdl.Log("filed to load texture file %s: SDL Error: %s\n", fileName, err)
		return nil
	}
	defer surf.Free()

	// Create texture from surface
	tex, err := g.renderer.CreateTextureFromSurface(surf)
	if err != nil {
		sdl.Log("failed to convert surface to texture for %s: SDL Error: %s\n", fileName, err)
		return nil
	}

	g.textures[fileName] = tex

	return tex
}

func (g *Game) Shutdown() (err error) {
	defer sdl.Quit()
	defer func() {
		if g.window != nil {
			err = g.window.Destroy()
		}
	}()
	defer func() {
		if g.renderer != nil {
			err = g.renderer.Destroy()
		}
	}()
	defer img.Quit()

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
