package chapter04

import "log"

func Start() {
	game := NewGame()
	defer func() {
		err := game.Shutdown()
		if err != nil {
			log.Fatalf("unable to initialize game: %s", err)
		}
	}()

	if err := game.Initialize(); err != nil {
		log.Fatalf("unable to initialize game: %s", err)
	}

	game.RunLoop()
}
