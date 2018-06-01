package tictactoe_test

import (
	"testing"

	. "github.com/martingallagher/tictactoe"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestMove(t *testing.T) {
	game := NewGame()

	// Make a move to test invalid turn
	game.Move(Crosses, 0, 0)

	tests := []struct {
		name   string
		player byte
		x, y   int
		err    error
	}{
		{"Invalid player", 'T', 0, 0, ErrInvalidPlayer},
		{"Invalid turn", Crosses, 1, 0, ErrInvalidTurn},
		{"Invalid position (overflow)", Naughts, 3, 0, ErrInvalidPos},
		{"Invalid position (taken)", Naughts, 0, 0, ErrInvalidPos},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			require.Equal(t, v.err, errors.Cause(game.Move(v.player, v.x, v.y)))
		})
	}
}

func TestStringer(t *testing.T) {
	game := NewGame()

	game.Move(Crosses, 0, 0)
	game.Move(Naughts, 1, 2)
	game.Move(Crosses, 1, 1)
	game.Move(Naughts, 0, 2)
	game.Move(Crosses, 2, 2)

	require.Equal(t, `|X| | |
| |X| |
|O|O|X|`, game.String())
	require.Equal(t, byte('X'), game.Winner())
}
