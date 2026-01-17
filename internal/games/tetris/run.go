package tetris

import (
	"fmt"
	"os"
	"time"

	"go-tetris/internal/input"
	"go-tetris/internal/termui"
)

func Run() error {
	g := New()

	inputCh, restore, err := input.StartKeyInput()
	if err != nil {
		return err
	}
	defer restore()
	defer termui.ShowCursor(os.Stdout)

	renderer := NewRenderer()
	ticker := time.NewTicker(g.TickDuration())
	defer ticker.Stop()

	renderer.Render(os.Stdout, g)

	for !g.Over {
		select {
		case ev := <-inputCh:
			switch ev.Type {
			case input.KeyArrowLeft:
				g.Move(-1)
			case input.KeyArrowRight:
				g.Move(1)
			case input.KeyArrowDown:
				g.SoftDrop()
			case input.KeyArrowUp:
				g.Rotate()
			case input.KeyRune:
				switch ev.Rune {
				case 'a', 'A':
					g.Move(-1)
				case 'd', 'D':
					g.Move(1)
				case 's', 'S':
					g.SoftDrop()
				case 'w', 'W':
					g.Rotate()
				case ' ':
					g.HardDrop()
				case 'q', 'Q':
					return nil
				}
			case input.KeyEscape:
				return nil
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
	for ev := range inputCh {
		if ev.Type == input.KeyRune && (ev.Rune == 'q' || ev.Rune == 'Q') {
			return nil
		}
		if ev.Type == input.KeyEscape {
			return nil
		}
	}
	return nil
}
