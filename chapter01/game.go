package chapter01

import (
	"github.com/veandco/go-sdl2/sdl"
)

const thickness = 15
const paddleHeight = 100.0

type Game struct {
	window     *sdl.Window
	renderer   *sdl.Renderer
	paddlePos  Vector2
	paddleDir  int
	ballPos    Vector2
	ballVel    Vector2
	ticksCount uint64
	isRunning  bool
}

type Vector2 struct {
	X, Y float32
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

	var err error

	g.window, err = sdl.CreateWindow("Chapter 1", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 1024, 768, sdl.WINDOW_SHOWN)
	if err != nil {
		sdl.Log("failed to create window: %s\n", err)
		return err
	}

	g.renderer, err = sdl.CreateRenderer(g.window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		sdl.Log("failed to create renderer: %s\n", err)
		return err
	}

	g.paddlePos.X = 10.0
	g.paddlePos.Y = 768.0 / 2.0

	g.ballPos.X = 1024.0 / 2.0
	g.ballPos.Y = 768.0 / 2.0
	g.ballVel.X = -200.0
	g.ballVel.Y = 235.0

	return nil
}

func (g *Game) RunLoop() {
	for g.isRunning {
		g.processInput()
		g.update()
		g.generateOutput()
	}
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

	return
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

		g.paddleDir = 0
		if state[sdl.SCANCODE_W] != 0 {
			g.paddleDir -= 1
		}
		if state[sdl.SCANCODE_S] != 0 {
			g.paddleDir += 1
		}
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

	// update paddle position
	if g.paddleDir != 0 {
		g.paddlePos.Y += float32(g.paddleDir) * 300.0 * deltaTime

		if g.paddlePos.Y < paddleHeight/2.0+thickness {
			g.paddlePos.Y = paddleHeight/2.0 + thickness
		} else if g.paddlePos.Y > 768.0-paddleHeight/2.0-thickness {
			g.paddlePos.Y = 768.0 - paddleHeight/2.0 - thickness
		}
	}

	// update ball position
	g.ballPos.X += g.ballVel.X * deltaTime
	g.ballPos.Y += g.ballVel.Y * deltaTime

	// Bounce if needed
	diff := g.paddlePos.Y - g.ballPos.Y
	if diff <= 0.0 {
		diff *= -1
	}

	// Did the ball hit the paddle?
	if diff <= paddleHeight/2.0 && g.ballPos.X <= 25.0 && g.ballPos.X >= 20.0 && g.ballVel.X < 0.0 {
		g.ballVel.X *= -1.0
	} else if g.ballPos.X <= 0.0 { // Did the ball go off the screen?
		g.isRunning = false
	} else if g.ballPos.X >= 1024.0-thickness && g.ballVel.X > 0.0 { // Did the ball hit the right wall?
		g.ballVel.X *= -1.0
	}

	if g.ballPos.Y <= thickness && g.ballVel.Y < 0.0 { // Did the ball hit the top wall?
		g.ballVel.Y *= -1.0
	} else if g.ballPos.Y >= 768.0-thickness && g.ballVel.Y > 0.0 { // Did the ball hit the bottom wall?
		g.ballVel.Y *= -1.0
	}
}

func (g *Game) generateOutput() {
	_ = g.renderer.SetDrawColor(0, 0, 255, 255)

	// Clear the back buffer
	_ = g.renderer.Clear()

	// Draw walls
	_ = g.renderer.SetDrawColor(255, 255, 255, 255)

	wall := sdl.Rect{X: 0, Y: 0, W: 1024, H: thickness}

	// Draw top wall
	_ = g.renderer.FillRect(&wall)

	// Draw bottom wall
	wall.Y = 768 - thickness
	_ = g.renderer.FillRect(&wall)

	// Draw right wall
	wall.X = 1024 - thickness
	wall.Y = 0
	wall.W = thickness
	wall.H = 1024
	_ = g.renderer.FillRect(&wall)

	// Draw paddle
	paddle := sdl.Rect{
		X: int32(g.paddlePos.X),
		Y: int32(g.paddlePos.Y - paddleHeight/2),
		W: thickness,
		H: paddleHeight,
	}
	_ = g.renderer.FillRect(&paddle)

	// Draw ball
	ball := sdl.Rect{
		X: int32(g.ballPos.X - thickness/2),
		Y: int32(g.ballPos.Y - thickness/2),
		W: thickness,
		H: thickness,
	}
	_ = g.renderer.FillRect(&ball)

	g.renderer.Present()
}
