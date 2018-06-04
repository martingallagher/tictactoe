package tictactoe

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// ErrInvalidMove - invalid move.
var ErrInvalidMove = errors.New("invalid move")

// Move represents a Tic Tac Toe move.
type Move struct {
	Player byte
	X, Y   int
}

// ParseMove parses a game move from the given input string.
func ParseMove(s string) (*Move, error) {
	if len(s) != 5 {
		return nil, ErrInvalidMove
	}

	p := strings.Split(s, " ")

	if len(p) != 3 {
		return nil, ErrInvalidMove
	}

	var player byte

	switch p[0] {
	case "x", "X":
		player = Crosses
	case "o", "O", "0": // Support zero for usability
		player = Naughts
	default:
		return nil, errors.Wrapf(ErrInvalidMove, "invalid player %q", p[0])
	}

	x, err := strconv.Atoi(p[1])

	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(p[2])

	if err != nil {
		return nil, err
	}

	return &Move{player, x, y}, nil
}
