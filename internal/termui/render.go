package termui

import (
	"fmt"
	"io"
	"strings"

	"go-tetris/internal/game"
)

type Renderer struct {}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) Render(w io.Writer, g *game.Game) {
	var b strings.Builder
	b.WriteString("\x1b[H\x1b[2J")
	b.WriteString("\x1b[?25l")

	b.WriteString("TETRIS\n")
	b.WriteString(fmt.Sprintf("Score: %d  Lines: %d  Level: %d\n\n", g.Score, g.Lines, g.Level))

	b.WriteString("+")
	for i := 0; i < game.BoardWidth*2; i++ {
		b.WriteString("-")
	}
	b.WriteString("+\n")

	for y := 0; y < game.BoardHeight; y++ {
		b.WriteString("|")
		for x := 0; x < game.BoardWidth; x++ {
			if g.Occupied(x, y) {
				b.WriteString("[]")
			} else {
				b.WriteString("  ")
			}
		}
		b.WriteString("|\n")
	}

	b.WriteString("+")
	for i := 0; i < game.BoardWidth*2; i++ {
		b.WriteString("-")
	}
	b.WriteString("+\n")

	b.WriteString("Controls: ←/→ move  ↓ soft drop  ↑ rotate  Space hard drop  Q quit\n")

	if g.Over {
		b.WriteString("\nGAME OVER\n")
	}

	fmt.Fprint(w, b.String())
}

func ShowCursor(w io.Writer) {
	fmt.Fprint(w, "\x1b[?25h")
}
