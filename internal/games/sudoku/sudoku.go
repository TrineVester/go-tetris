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
	index    int
}

type Puzzle struct {
	Grid     [9][9]int
	Solution [9][9]int
}

var puzzles = []Puzzle{
	{
		Grid: [9][9]int{
			{5, 3, 0, 0, 7, 0, 0, 0, 0},
			{6, 0, 0, 1, 9, 5, 0, 0, 0},
			{0, 9, 8, 0, 0, 0, 0, 6, 0},
			{8, 0, 0, 0, 6, 0, 0, 0, 3},
			{4, 0, 0, 8, 0, 3, 0, 0, 1},
			{7, 0, 0, 0, 2, 0, 0, 0, 6},
			{0, 6, 0, 0, 0, 0, 2, 8, 0},
			{0, 0, 0, 4, 1, 9, 0, 0, 5},
			{0, 0, 0, 0, 8, 0, 0, 7, 9},
		},
		Solution: [9][9]int{
			{5, 3, 4, 6, 7, 8, 9, 1, 2},
			{6, 7, 2, 1, 9, 5, 3, 4, 8},
			{1, 9, 8, 3, 4, 2, 5, 6, 7},
			{8, 5, 9, 7, 6, 1, 4, 2, 3},
			{4, 2, 6, 8, 5, 3, 7, 9, 1},
			{7, 1, 3, 9, 2, 4, 8, 5, 6},
			{9, 6, 1, 5, 3, 7, 2, 8, 4},
			{2, 8, 7, 4, 1, 9, 6, 3, 5},
			{3, 4, 5, 2, 8, 6, 1, 7, 9},
		},
	},
	{
		Grid: [9][9]int{
			{0, 2, 0, 6, 0, 8, 0, 0, 0},
			{5, 8, 0, 0, 0, 9, 7, 0, 0},
			{0, 0, 0, 0, 4, 0, 0, 0, 0},
			{3, 7, 0, 0, 0, 0, 5, 0, 0},
			{6, 0, 0, 0, 0, 0, 0, 0, 4},
			{0, 0, 8, 0, 0, 0, 0, 1, 3},
			{0, 0, 0, 0, 2, 0, 0, 0, 0},
			{0, 0, 9, 8, 0, 0, 0, 3, 6},
			{0, 0, 0, 3, 0, 6, 0, 9, 0},
		},
		Solution: [9][9]int{
			{1, 2, 3, 6, 7, 8, 9, 4, 5},
			{5, 8, 4, 2, 3, 9, 7, 6, 1},
			{9, 6, 7, 1, 4, 5, 3, 2, 8},
			{3, 7, 2, 4, 6, 1, 5, 8, 9},
			{6, 9, 1, 5, 8, 3, 2, 7, 4},
			{4, 5, 8, 7, 9, 2, 6, 1, 3},
			{8, 3, 6, 9, 2, 4, 1, 5, 7},
			{2, 1, 9, 8, 5, 7, 4, 3, 6},
			{7, 4, 5, 3, 1, 6, 8, 9, 2},
		},
	},
}

func New() *Game {
	g := &Game{}
	g.Reset(0)
	return g
}

func (g *Game) Reset(index int) {
	if len(puzzles) == 0 {
		return
	}
	if index < 0 {
		index = 0
	}
	if index >= len(puzzles) {
		index = 0
	}

	g.index = index
	g.Board = puzzles[index].Grid
	g.Solution = puzzles[index].Solution
	g.Fixed = [9][9]bool{}
	g.CursorX = 0
	g.CursorY = 0
	g.Solved = false
	g.Message = ""

	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if g.Board[y][x] != 0 {
				g.Fixed[y][x] = true
			}
		}
	}
}

func (g *Game) NextPuzzle() {
	if len(puzzles) == 0 {
		return
	}
	idx := g.index + 1
	if idx >= len(puzzles) {
		idx = 0
	}
	g.Reset(idx)
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
