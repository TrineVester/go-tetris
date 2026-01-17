package sudoku

import (
	"os"

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
	renderer.Render(os.Stdout, g)

	for ev := range inputCh {
		switch ev.Type {
		case input.KeyArrowUp:
			g.MoveCursor(0, -1)
		case input.KeyArrowDown:
			g.MoveCursor(0, 1)
		case input.KeyArrowLeft:
			g.MoveCursor(-1, 0)
		case input.KeyArrowRight:
			g.MoveCursor(1, 0)
		case input.KeyBackspace:
			g.SetValue(0)
		case input.KeyRune:
			switch ev.Rune {
			case 'q', 'Q':
				return nil
			case 'w', 'W':
				g.MoveCursor(0, -1)
			case 's', 'S':
				g.MoveCursor(0, 1)
			case 'a', 'A':
				g.MoveCursor(-1, 0)
			case 'd', 'D':
				g.MoveCursor(1, 0)
			case '0':
				g.SetValue(0)
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				g.SetValue(int(ev.Rune - '0'))
			}
		case input.KeyEscape:
			return nil
		}

		renderer.Render(os.Stdout, g)
	}

	return nil
}
