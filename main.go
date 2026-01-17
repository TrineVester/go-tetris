package main

import (
	"go-tetris/internal/games/sudoku"
	"go-tetris/internal/games/tetris"
	"go-tetris/internal/hub"
	"go-tetris/internal/termui"
)

func main() {
	termui.EnableANSI()

	for {
		choice := hub.ShowMenu()
		switch choice {
		case "1":
			_ = tetris.Run()
		case "2":
			_ = sudoku.Run()
		case "q":
			return
		}
	}
}
