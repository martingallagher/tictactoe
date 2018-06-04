package main

import (
	"bufio"
	"log"
	"os"

	"github.com/martingallagher/tictactoe"
)

func main() {
	log.SetFlags(0)
	log.Println("Tic Tac Toe Game Started!")

	game := tictactoe.NewGame()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		move, err := tictactoe.ParseMove(scanner.Text())

		if err != nil {
			log.Printf("Unable to make move: %s", err)

			continue
		}

		err = game.Move(move.Player, move.X, move.Y)

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
