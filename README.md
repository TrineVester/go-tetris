# Go Tetris (Terminal)

A simple Tetris clone you can play directly in a terminal window.

## Requirements

- Go 1.22+
- A terminal that supports ANSI escape sequences (Windows Terminal, PowerShell, or VS Code terminal)

## Run

In a terminal:

1) Open Windows Terminal or PowerShell (outside VS Code) in the project folder.
2) Run:

go run .

## Controls

- Left/Right: move
- Down: soft drop
- Up: rotate
- Space: hard drop
- Q: quit

## Notes

If the screen prints multiple frames instead of updating in-place, use Windows Terminal and ensure ANSI support is enabled. The app enables ANSI automatically on Windows.
