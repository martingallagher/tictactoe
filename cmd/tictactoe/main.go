package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/martingallagher/tictactoe"
	"github.com/pkg/errors"
)

func main() {
	log.SetFlags(0)

	game := tictactoe.NewGame()
	scanner := bufio.NewScanner(os.Stdin)

	log.Println("Tic Tac Toe Game Started!")

	for scanner.Scan() {
		move, err := parseInput(scanner.Text())

		if err != nil {
			log.Printf("Unable to make move: %s", err)

			continue
		}

		err = game.Move(move.player, move.x, move.y)

		if err != nil {
			log.Printf("Move error: %s", err)

			continue
		}

		winner := game.Winner()

		if winner != tictactoe.None {
			log.Printf("%q is the winner!", winner)

			break
		}

		log.Printf("Game status:\n%s", game)
	}
}

type move struct {
	player byte
	x, y   int
}

var errInvalidMove = errors.New("invalid move")

func parseInput(input string) (*move, error) {
	if len(input) != 5 {
		return nil, errInvalidMove
	}

	p := strings.Split(input, " ")

	if len(p) != 3 {
		return nil, errInvalidMove
	}

	var player byte

	switch p[0] {
	case "x", "X":
		player = tictactoe.Crosses
	case "o", "O", "0": // Support zero for usability
		player = tictactoe.Naughts
	default:
		return nil, errors.Wrapf(errInvalidMove, "invalid player %q", p[0])
	}

	x, err := strconv.Atoi(p[1])

	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(p[2])

	if err != nil {
		return nil, err
	}

	return &move{player, x, y}, nil
}
