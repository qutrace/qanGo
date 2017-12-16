package main

import (
	"fmt"

	"github.com/qutrace/qanGo/ai"
	"github.com/qutrace/qanGo/game"
)

func main() {
	moves := make(chan game.Move)
	board := make(chan game.Board)
	abort := make(chan bool)
	report := make(chan error)
	done := make(chan game.Board)

	cur := game.Board{}
	go game.PlayChan(&game.Board{}, board, moves, abort, report, done)
	for {
		select {
		case err := <-report:
			fmt.Println(err)
		case cur = <-board:
			fmt.Println(cur)
			if !cur.GetPlayer() {
				moves <- ai.GetMove(cur)
			} else {
				scan(moves, abort)
			}
		case cur = <-done:
			fmt.Println(cur)
			fmt.Println(cur.GetState())
			fmt.Println("finished")
			return
		}
	}
}

func scan(c chan game.Move, abort chan bool) {
	m := game.Move{}
	done := false
	for !done {
		_, err := fmt.Scan(&(m.X), &(m.Y))
		if err != nil {
			abort <- true
			return
			fmt.Println(err)
		} else {
			done = true
		}
	}
	c <- m
}
