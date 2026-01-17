package termui

import (
	"bufio"
	"os"

	"golang.org/x/term"
)

type Action int

const (
	ActionNone Action = iota
	ActionLeft
	ActionRight
	ActionDown
	ActionRotate
	ActionDrop
	ActionQuit
)

func StartInput() (<-chan Action, func(), error) {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return nil, func() {}, err
	}

	restore := func() {
		_ = term.Restore(fd, oldState)
	}

	ch := make(chan Action, 16)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			b, err := reader.ReadByte()
			if err != nil {
				return
			}
			switch b {
			case 27: // ESC
				b2, _ := reader.ReadByte()
				b3, _ := reader.ReadByte()
				if b2 == '[' {
					switch b3 {
					case 'A':
						ch <- ActionRotate
					case 'B':
						ch <- ActionDown
					case 'C':
						ch <- ActionRight
					case 'D':
						ch <- ActionLeft
					}
				}
			case 0, 224: // Windows arrow keys in legacy mode
				b2, _ := reader.ReadByte()
				switch b2 {
				case 72:
					ch <- ActionRotate
				case 80:
					ch <- ActionDown
				case 77:
					ch <- ActionRight
				case 75:
					ch <- ActionLeft
				}
			case 'a', 'A':
				ch <- ActionLeft
			case 'd', 'D':
				ch <- ActionRight
			case 's', 'S':
				ch <- ActionDown
			case 'w', 'W':
				ch <- ActionRotate
			case ' ':
				ch <- ActionDrop
			case 'q', 'Q':
				ch <- ActionQuit
			}
		}
	}()

	return ch, restore, nil
}
