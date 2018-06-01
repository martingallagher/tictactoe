package tictactoe

import (
	"strings"

	"github.com/pkg/errors"
)

// TicTacToe game errors.
var (
	ErrInvalidPlayer = errors.New("invalid player")
	ErrInvalidPos    = errors.New("invalid move position")
	ErrInvalidTurn   = errors.New("invalid turn")
)

// TicTacToe represents a Tic Tac Toe board and game state data.
type TicTacToe struct {
	board  [3][3]byte
	last   byte
	winner byte
	moves  int
}

// TicTacToe game & player constants.
const (
	None    byte = ' '
	Naughts byte = 'O'
	Crosses byte = 'X'
	Draw    byte = '/'
)

// NewGame creates a new Tic Tac Toe game instance.
func NewGame() *TicTacToe {
	return &TicTacToe{
		last:   None,
		winner: None,
		board: [3][3]byte{
			{None, None, None},
			{None, None, None},
			{None, None, None},
		},
	}
}

var winningPositions = [...][3]struct{ x, y int }{
	// Vertical
	{{0, 0}, {0, 1}, {0, 2}},
	{{1, 0}, {1, 1}, {1, 2}},
	{{2, 0}, {2, 1}, {2, 2}},

	// Horizontal
	{{0, 0}, {1, 0}, {2, 0}},
	{{1, 0}, {1, 1}, {2, 1}},
	{{2, 0}, {2, 1}, {2, 2}},

	// Diagonal right
	{{0, 0}, {1, 1}, {2, 2}},
	// Diagonal left
	{{0, 2}, {1, 1}, {2, 0}},
}

// Winner determines the winning state.
func (t *TicTacToe) Winner() byte {
	if t.winner != None {
		return t.winner
	}

	for _, v := range winningPositions {
		winner := None

		for i, c := range v {
			player := t.board[c.y][c.x]

			// 'None' value; can't possibly be a winning position
			if player == None {
				break
			}

			// First value, set active player
			if i == 0 {
				winner = player

				continue
			}

			// Winning streak over
			if player != winner {
				break
			}

			// 3 winning positions results in a victory
			if i == 2 {
				t.winner = winner

				return t.winner
			}
		}
	}

	if t.moves == 9 {
		return Draw
	}

	return None
}

// Move attempts to perform a game move for the given player and position.
func (t *TicTacToe) Move(player byte, x, y int) error {
	if player != Naughts && player != Crosses {
		return ErrInvalidPlayer
	}

	if t.last == player {
		return ErrInvalidTurn
	}

	if x < 0 || x > 2 || y < 0 || y > 2 {
		return errors.Wrapf(ErrInvalidPos, "position overflow - x: %d, y: %d", x, y)
	}

	v := t.board[y][x]

	if v != None {
		return errors.Wrap(ErrInvalidPos, "position taken")
	}

	t.board[y][x] = player
	t.last = player
	t.moves++

	return nil
}

// String fulfils the fmt.Stringer interface.
func (t *TicTacToe) String() string {
	var sb strings.Builder

	for i, v := range t.board {
		for _, c := range v {
			sb.WriteByte('|')
			sb.WriteByte(c)
		}

		sb.WriteByte('|')

		if i < 2 {
			sb.WriteByte('\n')
		}
	}

	return sb.String()
}
