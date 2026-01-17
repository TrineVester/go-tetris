//go:build !windows

package input

import "os"

func openInputDevice() (*os.File, error) {
	return os.OpenFile("/dev/tty", os.O_RDWR, 0)
}
