package main

import (
	"fmt"
	"os"
	"time"

	"go-tetris/internal/game"
	"go-tetris/internal/termui"
)

func main() {
	g := game.New()

	input, restore, err := termui.StartInput()
	if err != nil {
		fmt.Println("Failed to init input:", err)
		return
	}
	defer restore()
	defer termui.ShowCursor(os.Stdout)

	renderer := termui.NewRenderer()
	ticker := time.NewTicker(g.TickDuration())
	defer ticker.Stop()

	renderer.Render(os.Stdout, g)

	for !g.Over {
		select {
		case act := <-input:
			if act == termui.ActionQuit {
				return
			}
			switch act {
			case termui.ActionLeft:
				g.Move(-1)
			case termui.ActionRight:
				g.Move(1)
			case termui.ActionDown:
				g.SoftDrop()
			case termui.ActionRotate:
				g.Rotate()
			case termui.ActionDrop:
				g.HardDrop()
			}
		case <-ticker.C:
			prevLevel := g.Level
			g.Step()
			if g.Level != prevLevel {
				ticker.Reset(g.TickDuration())
			}
		}

		renderer.Render(os.Stdout, g)
	}

	renderer.Render(os.Stdout, g)
	fmt.Println("Press Q to exit.")
	for act := range input {
		if act == termui.ActionQuit {
			return
		}
	}
}
