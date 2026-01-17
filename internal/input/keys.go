package input

import (
	"bufio"
	"os"

	"golang.org/x/term"
)

type KeyType int

const (
	KeyNone KeyType = iota
	KeyArrowUp
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
	KeyBackspace
	KeyEnter
	KeyEscape
	KeyRune
)

type KeyEvent struct {
	Type KeyType
	Rune rune
}

func StartKeyInput() (<-chan KeyEvent, func(), error) {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return nil, func() {}, err
	}

	restore := func() {
		_ = term.Restore(fd, oldState)
	}

	ch := make(chan KeyEvent, 32)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			b, err := reader.ReadByte()
			if err != nil {
				close(ch)
				return
			}
			switch b {
			case 27: // ESC
				b2, _ := reader.ReadByte()
				b3, _ := reader.ReadByte()
				if b2 == '[' {
					switch b3 {
					case 'A':
						ch <- KeyEvent{Type: KeyArrowUp}
					case 'B':
						ch <- KeyEvent{Type: KeyArrowDown}
					case 'C':
						ch <- KeyEvent{Type: KeyArrowRight}
					case 'D':
						ch <- KeyEvent{Type: KeyArrowLeft}
					}
				} else {
					ch <- KeyEvent{Type: KeyEscape}
				}
			case 0, 224: // Windows legacy arrows
				b2, _ := reader.ReadByte()
				switch b2 {
				case 72:
					ch <- KeyEvent{Type: KeyArrowUp}
				case 80:
					ch <- KeyEvent{Type: KeyArrowDown}
				case 77:
					ch <- KeyEvent{Type: KeyArrowRight}
				case 75:
					ch <- KeyEvent{Type: KeyArrowLeft}
				}
			case 8, 127:
				ch <- KeyEvent{Type: KeyBackspace}
			case 13:
				ch <- KeyEvent{Type: KeyEnter}
			default:
				ch <- KeyEvent{Type: KeyRune, Rune: rune(b)}
			}
		}
	}()

	return ch, restore, nil
}
