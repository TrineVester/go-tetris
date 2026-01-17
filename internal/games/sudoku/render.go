package sudoku

import (
	"fmt"
	"io"
	"strings"
)

type Renderer struct{}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) Render(w io.Writer, g *Game) {
	var b strings.Builder
	b.WriteString("\x1b[H\x1b[2J")
	b.WriteString("\x1b[?25l")

	b.WriteString("SUDOKU\n")
	b.WriteString("Controls: arrows/WASD move  1-9 set  0/Backspace clear  Q quit\n\n")

	b.WriteString("+-------+-------+-------+\n")
	for y := 0; y < 9; y++ {
		b.WriteString("|")
		for x := 0; x < 9; x++ {
			b.WriteString(" ")
			cell := g.CellDisplay(x, y)
			isCursor := g.CursorX == x && g.CursorY == y
			if isCursor {
				b.WriteString("\x1b[7m")
			}
			if g.Fixed[y][x] {
				b.WriteString("\x1b[97m")
			} else if cell != "." {
				b.WriteString("\x1b[36m")
			} else {
				b.WriteString("\x1b[2m")
			}
			b.WriteString(cell)
			b.WriteString("\x1b[0m")
			b.WriteString(" ")

			if (x+1)%3 == 0 {
				b.WriteString("|")
			}
		}
		b.WriteString("\n")
		if (y+1)%3 == 0 {
			b.WriteString("+-------+-------+-------+\n")
		}
	}

	if g.Message != "" {
		b.WriteString("\n")
		b.WriteString(fmt.Sprintf("%s\n", g.Message))
	}

	fmt.Fprint(w, b.String())
}
