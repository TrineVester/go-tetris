//go:build windows

package input

import "os"

func openInputDevice() (*os.File, error) {
	return os.OpenFile("CONIN$", os.O_RDWR, 0)
}
