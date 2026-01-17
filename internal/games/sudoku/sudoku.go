package sudoku

import "fmt"

type Game struct {
	Board    [9][9]int
	Fixed    [9][9]bool
	Solution [9][9]int
	CursorX  int
	CursorY  int
	Solved   bool
	Message  string
}

func New() *Game {
	puzzle := [9][9]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	solution := [9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	g := &Game{Board: puzzle, Solution: solution}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if puzzle[y][x] != 0 {
				g.Fixed[y][x] = true
			}
		}
	}
	return g
}

func (g *Game) MoveCursor(dx, dy int) {
	g.CursorX += dx
	g.CursorY += dy
	if g.CursorX < 0 {
		g.CursorX = 0
	}
	if g.CursorX > 8 {
		g.CursorX = 8
	}
	if g.CursorY < 0 {
		g.CursorY = 0
	}
	if g.CursorY > 8 {
		g.CursorY = 8
	}
}

func (g *Game) SetValue(v int) {
	if g.Solved {
		return
	}
	if g.Fixed[g.CursorY][g.CursorX] {
		g.Message = "Cell is fixed"
		return
	}
	if v < 0 || v > 9 {
		return
	}
	g.Board[g.CursorY][g.CursorX] = v
	g.checkSolved()
}

func (g *Game) checkSolved() {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if g.Board[y][x] == 0 {
				g.Message = ""
				return
			}
			if g.Board[y][x] != g.Solution[y][x] {
				g.Message = ""
				return
			}
		}
	}
	g.Solved = true
	g.Message = "Solved!"
}

func (g *Game) CellDisplay(x, y int) string {
	v := g.Board[y][x]
	if v == 0 {
		return "."
	}
	return fmt.Sprintf("%d", v)
}
