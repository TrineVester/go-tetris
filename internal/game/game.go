package game

import (
	"math/rand"
	"time"
)

const (
	BoardWidth  = 10
	BoardHeight = 20
)

type Game struct {
	Board   [BoardHeight][BoardWidth]int
	Current Piece
	Next    Piece
	Score   int
	Lines   int
	Level   int
	Over    bool

	rng *rand.Rand
}

type Piece struct {
	Shape int
	Rot   int
	X     int
	Y     int
}

type MoveResult int

const (
	MoveOK MoveResult = iota
	MoveBlocked
)

var shapes = [7][4]uint16{
	// I
	{0x00F0, 0x2222, 0x00F0, 0x2222},
	// O
	{0x0066, 0x0066, 0x0066, 0x0066},
	// T
	{0x0072, 0x0262, 0x0270, 0x0232},
	// S
	{0x0036, 0x0462, 0x0036, 0x0462},
	// Z
	{0x0063, 0x0264, 0x0063, 0x0264},
	// J
	{0x0071, 0x0226, 0x0470, 0x0322},
	// L
	{0x0074, 0x0622, 0x0170, 0x0223},
}

func New() *Game {
	g := &Game{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	g.Current = g.randomPiece()
	g.Next = g.randomPiece()
	g.Current.X = 3
	g.Current.Y = 0
	g.Next.X = 3
	g.Next.Y = 0
	return g
}

func (g *Game) TickDuration() time.Duration {
	base := 500 * time.Millisecond
	dec := time.Duration(g.Level) * 30 * time.Millisecond
	if dec > 400*time.Millisecond {
		dec = 400 * time.Millisecond
	}
	return base - dec
}

func (g *Game) Move(dx int) MoveResult {
	if g.Over {
		return MoveBlocked
	}
	p := g.Current
	p.X += dx
	if g.collides(p) {
		return MoveBlocked
	}
	g.Current = p
	return MoveOK
}

func (g *Game) Rotate() MoveResult {
	if g.Over {
		return MoveBlocked
	}
	p := g.Current
	p.Rot = (p.Rot + 1) % 4
	if g.collides(p) {
		return MoveBlocked
	}
	g.Current = p
	return MoveOK
}

func (g *Game) SoftDrop() MoveResult {
	if g.Over {
		return MoveBlocked
	}
	p := g.Current
	p.Y++
	if g.collides(p) {
		return MoveBlocked
	}
	g.Current = p
	return MoveOK
}

func (g *Game) HardDrop() {
	if g.Over {
		return
	}
	for {
		p := g.Current
		p.Y++
		if g.collides(p) {
			g.lockPiece()
			return
		}
		g.Current = p
	}
}

func (g *Game) Step() {
	if g.Over {
		return
	}
	p := g.Current
	p.Y++
	if g.collides(p) {
		g.lockPiece()
		return
	}
	g.Current = p
}

func (g *Game) Occupied(x, y int) bool {
	if x < 0 || x >= BoardWidth || y < 0 || y >= BoardHeight {
		return true
	}
	if g.Board[y][x] != 0 {
		return true
	}
	return g.pieceCell(g.Current, x, y)
}

func (g *Game) CellValue(x, y int) int {
	if x < 0 || x >= BoardWidth || y < 0 || y >= BoardHeight {
		return 0
	}
	if g.Board[y][x] != 0 {
		return g.Board[y][x]
	}
	if g.pieceCell(g.Current, x, y) {
		return g.Current.Shape + 1
	}
	return 0
}

func (g *Game) GhostPiece() Piece {
	if g.Over {
		return g.Current
	}
	p := g.Current
	for {
		next := p
		next.Y++
		if g.collides(next) {
			return p
		}
		p = next
	}
}

func (g *Game) GhostCell(x, y int) bool {
	if x < 0 || x >= BoardWidth || y < 0 || y >= BoardHeight {
		return false
	}
	if g.Board[y][x] != 0 {
		return false
	}
	if g.pieceCell(g.Current, x, y) {
		return false
	}
	ghost := g.GhostPiece()
	if ghost.Y == g.Current.Y {
		return false
	}
	return g.pieceCell(ghost, x, y)
}

func (g *Game) lockPiece() {
	p := g.Current
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if !pieceMask(p.Shape, p.Rot, r, c) {
				continue
			}
			x := p.X + c
			y := p.Y + r
			if y >= 0 && y < BoardHeight && x >= 0 && x < BoardWidth {
				g.Board[y][x] = p.Shape + 1
			}
		}
	}

	lines := g.clearLines()
	if lines > 0 {
		scoreTable := []int{0, 100, 300, 500, 800}
		g.Score += scoreTable[lines] * (g.Level + 1)
		g.Lines += lines
		g.Level = g.Lines / 10
	}

	g.Current = g.Next
	g.Current.X = 3
	g.Current.Y = 0
	g.Next = g.randomPiece()

	if g.collides(g.Current) {
		g.Over = true
	}
}

func (g *Game) clearLines() int {
	cleared := 0
	for y := BoardHeight - 1; y >= 0; y-- {
		full := true
		for x := 0; x < BoardWidth; x++ {
			if g.Board[y][x] == 0 {
				full = false
				break
			}
		}
		if !full {
			continue
		}

		cleared++
		for row := y; row > 0; row-- {
			g.Board[row] = g.Board[row-1]
		}
		g.Board[0] = [BoardWidth]int{}
		y++
	}
	return cleared
}

func (g *Game) collides(p Piece) bool {
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if !pieceMask(p.Shape, p.Rot, r, c) {
				continue
			}
			x := p.X + c
			y := p.Y + r
			if x < 0 || x >= BoardWidth || y >= BoardHeight {
				return true
			}
			if y >= 0 && g.Board[y][x] != 0 {
				return true
			}
		}
	}
	return false
}

func (g *Game) pieceCell(p Piece, x, y int) bool {
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if !pieceMask(p.Shape, p.Rot, r, c) {
				continue
			}
			if p.X+c == x && p.Y+r == y {
				return true
			}
		}
	}
	return false
}

func pieceMask(shape, rot, r, c int) bool {
	mask := shapes[shape][rot]
	idx := r*4 + c
	return mask&(1<<idx) != 0
}

func (g *Game) randomPiece() Piece {
	return Piece{
		Shape: g.rng.Intn(len(shapes)),
		Rot:   0,
		X:     3,
		Y:     0,
	}
}
