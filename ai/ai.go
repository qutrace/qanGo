package ai

import (
	"github.com/qutrace/qanGo/game"
)

type BoardMove struct {
	game.Board
	game.Move
}

func GetMove(b game.Board) game.Move {
	boards := make([]BoardMove, 0)
	for x := 0; x < 6; x++ {
		for y := 0; y < 6; y++ {
			board := b
			err := (&board).Apply(game.Move{x, y})
			if err == nil {
				boards = append(boards, BoardMove{board, game.Move{x, y}})
			}
		}
	}
	player := b.GetPlayer()
	rating := 0
	pos := -1
	for p, bm := range boards {
		r := GetRating(&(bm.Board), player)
		if r > rating {
			rating = r
			pos = p
		}
	}
	return boards[pos].Move
}
