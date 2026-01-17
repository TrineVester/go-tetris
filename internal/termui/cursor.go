package termui

import (
	"fmt"
	"io"
	"os"
)

func ShowCursor(w io.Writer) {
	fmt.Fprint(w, "\x1b[?25h")
}

func ShowCursorStdout() {
	ShowCursor(os.Stdout)
}
