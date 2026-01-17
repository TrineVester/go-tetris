package hub

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ShowMenu() string {
	fmt.Print("\x1b[H\x1b[2J")
	fmt.Println("GAME HUB")
	fmt.Println("1) Tetris")
	fmt.Println("2) Sudoku")
	fmt.Println("Q) Quit")
	fmt.Print("Choose a game: ")

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(strings.ToLower(line))
	return line
}
