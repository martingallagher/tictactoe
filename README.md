# Tic Tac Toe

A simple Tic Tac Toe game implementation.

## Build Instructions

1. Ensure depdencies (`go get -u github.com/golang/dep/cmd/dep`) via `dep ensure -vendor-only`
2. Run tests `go test -v -cover`
3. Run the game CLI in `cmd/tictactoe` via `go run main.go`

## Game Instructions

The game takes a string input of player (X or O [letter O OR zero]) and a board position (0-2).

Example inputs:

- `x 0 0`
- `O 1 0`
- `X 2 2`
- `0 1 1`
