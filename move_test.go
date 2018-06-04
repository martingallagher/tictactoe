package tictactoe_test

import (
	"testing"

	. "github.com/martingallagher/tictactoe"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestParseMove(t *testing.T) {
	tests := []struct {
		input  string
		player byte
		x, y   int
		err    error
	}{
		{"x 0 0", Crosses, 0, 0, nil},
		{"0 1 2", Naughts, 1, 2, nil},
		{"t 1 1", None, 0, 0, ErrInvalidMove},
		{"x    ", None, 0, 0, ErrInvalidMove},
		{"x 9 9", Crosses, 9, 9, nil},
	}

	for _, v := range tests {
		m, err := ParseMove(v.input)

		if v.err != nil {
			require.Nil(t, m)
			require.Equal(t, v.err, errors.Cause(err))

			continue
		}

		require.NoError(t, err)
		require.Equal(t, v.player, m.Player)
		require.Equal(t, v.x, m.X)
		require.Equal(t, v.y, m.Y)
	}

}
