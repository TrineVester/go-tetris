//go:build windows

package termui

import "golang.org/x/sys/windows"

func EnableANSI() {
	hOut, err := windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
	if err != nil {
		return
	}

	var mode uint32
	if err := windows.GetConsoleMode(hOut, &mode); err != nil {
		return
	}

	mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	_ = windows.SetConsoleMode(hOut, mode)

	hIn, err := windows.GetStdHandle(windows.STD_INPUT_HANDLE)
	if err != nil {
		return
	}
	var inMode uint32
	if err := windows.GetConsoleMode(hIn, &inMode); err != nil {
		return
	}
	inMode |= windows.ENABLE_VIRTUAL_TERMINAL_INPUT
	_ = windows.SetConsoleMode(hIn, inMode)
}
