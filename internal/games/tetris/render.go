package tetris

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

	b.WriteString("TETRIS\n")
	b.WriteString(fmt.Sprintf("Score: %d  Lines: %d  Level: %d\n\n", g.Score, g.Lines, g.Level))

	b.WriteString("+")
	for i := 0; i < BoardWidth*2; i++ {
		b.WriteString("-")
	}
	b.WriteString("+\n")

	for y := 0; y < BoardHeight; y++ {
		b.WriteString("|")
		for x := 0; x < BoardWidth; x++ {
			val := g.CellValue(x, y)
			if val != 0 {
				b.WriteString(colorFor(val))
				b.WriteString("[]")
				b.WriteString("\x1b[0m")
			} else if g.GhostCell(x, y) {
				b.WriteString("\x1b[2m")
				b.WriteString(colorFor(g.Current.Shape + 1))
				b.WriteString("[]")
				b.WriteString("\x1b[0m")
			} else {
				b.WriteString("  ")
			}
		}
		b.WriteString("|\n")
	}

	b.WriteString("+")
	for i := 0; i < BoardWidth*2; i++ {
		b.WriteString("-")
	}
	b.WriteString("+\n")

	b.WriteString("Controls: ←/→ move  ↓ soft drop  ↑ rotate  Space hard drop  Q quit\n")

	if g.Over {
		b.WriteString("\nGAME OVER\n")
	}

	fmt.Fprint(w, b.String())
}

func colorFor(v int) string {
	switch v {
	case 1: // I
		return "\x1b[36m"
	case 2: // O
		return "\x1b[33m"
	case 3: // T
		return "\x1b[35m"
	case 4: // S
		return "\x1b[32m"
	case 5: // Z
		return "\x1b[31m"
	case 6: // J
		return "\x1b[34m"
	case 7: // L
		return "\x1b[93m"
	default:
		return "\x1b[37m"
	}
}
