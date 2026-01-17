# Go Game Hub (Terminal)

A small terminal game hub with Tetris and Sudoku.

## Requirements

- Go 1.22+
- A terminal that supports ANSI escape sequences (Windows Terminal, PowerShell, or VS Code terminal)

## Run

In a terminal:

1) Open Windows Terminal or PowerShell (outside VS Code) in the project folder.
2) Run:

go run .

## Game Hub

Pick a game:

- 1) Tetris
- 2) Sudoku
- Q to quit

## Tetris Controls

- Left/Right: move
- Down: soft drop
- Up: rotate
- Space: hard drop
- Q: quit

## Sudoku Controls

- Arrows/WASD: move cursor
- 1-9: set value
- 0 or Backspace: clear
- Q: quit

## Notes

If the screen prints multiple frames instead of updating in-place, use Windows Terminal and ensure ANSI support is enabled. The app enables ANSI automatically on Windows.
