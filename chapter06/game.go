package chapter06

import (
	"slices"

	"github.com/ishtaka/go-game-programming/chapter06/math"
	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	renderer *Renderer

	ticksCount uint64
	isRunning  bool

	// All the actors in the game
	actors []Actor
	// Any pending actors
	pendingActors []Actor
	// Track if we're updating actors right now
	updatingActors bool
}

func NewGame() *Game {
	return &Game{
		ticksCount: 0,
		isRunning:  true,
	}
}

func (g *Game) Initialize() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		sdl.Log("unable to initialize SDL: %s\n", err)
		return err
	}

	g.renderer = NewRenderer(g)
	if err := g.renderer.Initialize(1024.0, 768.0); err != nil {
		sdl.Log("failed to initialize renderer: %s\n", err)
		_ = g.renderer.Shutdown()
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
	g.renderer.Draw()
}

func (g *Game) loadData() {
	a := NewActor(g)
	a.SetPosition(math.Vector3{X: 200.0, Y: 75.0, Z: 0.0})
	a.SetScale(100)
	q := math.NewQuaternionFromVec(math.Vector3UnitY, -math.PiOver2)
	q = q.Concatenate(math.NewQuaternionFromVec(math.Vector3UnitZ, math.Pi+math.Pi/4.0))
	a.SetRotation(q)

	mc := NewMeshComponent(a, DefaultUpdateOrder)
	m := g.GetRenderer().GetMesh("Assets/Cube.gpmesh")
	mc.SetMesh(m)
	a.AddComponent(mc)
	g.AddActor(a)

	a = NewActor(g)
	a.SetPosition(math.Vector3{X: 200.0, Y: -75.0, Z: 0.0})
	a.SetScale(3)

	mc = NewMeshComponent(a, DefaultUpdateOrder)
	m = g.GetRenderer().GetMesh("Assets/Sphere.gpmesh")
	mc.SetMesh(m)
	a.AddComponent(mc)
	g.AddActor(a)

	// Setup floor
	const start float32 = -1250.0
	const size float32 = 250.0
	for i := range 10 {
		fi := float32(i)
		for j := range 10 {
			fj := float32(j)
			p := NewPlaneActor(g)
			p.SetPosition(math.Vector3{
				X: start + fi*size,
				Y: start + fj*size,
				Z: -100,
			})
			g.AddActor(p)
		}
	}

	// Left/right walls
	q = math.NewQuaternionFromVec(math.Vector3UnitX, math.PiOver2)
	for i := range 10 {
		fi := float32(i)

		a = NewPlaneActor(g)
		a.SetPosition(math.Vector3{
			X: start + fi*size,
			Y: start - size,
			Z: 0.0,
		})
		a.SetRotation(q)
		g.AddActor(a)

		a = NewPlaneActor(g)
		a.SetPosition(math.Vector3{
			X: start + fi*size,
			Y: -start + size,
			Z: 0.0,
		})
		a.SetRotation(q)
		g.AddActor(a)
	}

	// Forward/back walls
	q = q.Concatenate(math.NewQuaternionFromVec(math.Vector3UnitZ, math.PiOver2))
	for i := range 10 {
		fi := float32(i)

		a = NewPlaneActor(g)
		a.SetPosition(math.Vector3{
			X: start - size,
			Y: start + fi*size,
			Z: 0.0,
		})
		a.SetRotation(q)
		g.AddActor(a)

		a = NewPlaneActor(g)
		a.SetPosition(math.Vector3{
			X: -start + size,
			Y: start + fi*size,
			Z: 0.0,
		})
		a.SetRotation(q)
		g.AddActor(a)
	}

	// Setup lights
	g.renderer.SetAmbientLight(math.Vector3{X: 0.2, Y: 0.2, Z: 0.2})
	dir := g.renderer.dirLight
	dir.Direction = math.Vector3{X: 0.0, Y: -0.7, Z: -0.7}
	dir.DiffuseColor = math.Vector3{X: 0.78, Y: 0.88, Z: 1.0}
	dir.SpecColor = math.Vector3{X: 0.8, Y: 0.8, Z: 0.8}

	// Camera actor
	camera := NewCameraActor(g)
	g.AddActor(camera)

}

func (g *Game) unloadData() {
	// Delete actors
	for len(g.actors) > 0 {
		g.actors[0].Destroy()
	}

	if g.renderer != nil {
		g.renderer.UnloadData()
	}
}

func (g *Game) Shutdown() (err error) {
	defer sdl.Quit()
	defer func() {
		if g.renderer != nil {
			err = g.renderer.Shutdown()
		}
	}()

	g.unloadData()

	return
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

func (g *Game) GetRenderer() *Renderer {
	return g.renderer
}
