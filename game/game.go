package game

type Cell *bool
type Board [6][6]Cell
type Move struct {
	X int
	Y int
}
type State int

const (
	Unknown State = iota
	WinTrue
	WinFalse
	Draw
)

func (b *Board) Apply(m Move) (e error) {
	e = MoveError{"not on board"}
	if m.X > 5 {
		return
	}
	if m.X < 0 {
		return
	}
	if m.Y > 5 {
		return
	}
	if m.Y < 0 {
		return
	}
	e = MoveError{"already placed"}
	place := &(b[m.Y][m.X])
	if *place != nil {
		return
	}
	e = nil
	player := b.GetPlayer()
	*place = &player
	return
}

func (b *Board) GetPlayer() bool {
	return b.GetTurn()%2 == 0
}

func (b *Board) GetTurn() int {
	res := 0
	for _, row := range b {
		for _, cell := range row {
			if cell != nil {
				res++
			}
		}
	}
	return res
}
func (b *Board) GetState() State {
	return rateBoard(b)
}

func (s *State) Running() bool {
	return *s == Unknown
}

type MoveError struct {
	err string
}

func (m MoveError) Error() string {
	return m.err
}

func PlayCustom(init Board, a func(b Board) Move, b func(b Board) Move, print func(b Board), finished func(s State), printe func(e error)) {
	board := init
	state := board.GetState()
	getMove := a

	for state.Running() {

		print(board)

		//select getMove function
		switch board.GetPlayer() {
		case true:
			getMove = a
		case false:
			getMove = b
		default:
			return
		}

		//apply (get) move until move is legal
		applyUntilLegal(&board, getMove, printe)

		//get new board rating
		state = board.GetState()
	}

	//final board state
	print(board)
	//report result
	finished(state)
}

func applyUntilLegal(board *Board, getMove func(b Board) Move, printe func(e error)) {
	move := getMove(*board)
	err := board.Apply(move)
	for err != nil {
		printe(err)
		move = getMove(*board)
		err = board.Apply(move)
	}
}

func PlayCustomMirror(init Board, getMove func(b Board) Move, print func(b Board), finished func(s State), printe func(e error)) {
	PlayCustom(init, getMove, getMove, print, finished, printe)
}

type Game struct {
	Board      Board
	GetMoveA   func(b Board) Move
	GetMoveB   func(b Board) Move
	PrintBoard func(b Board)
	Finished   func(s State)
	PrintError func(e error)
}

func (g Game) Play() {
	PlayCustom(g.Board, g.GetMoveA, g.GetMoveB, g.PrintBoard, g.Finished, g.PrintError)
}

func PlayChan(board *Board, print chan Board, getmove chan Move, abort chan bool, report chan error, done chan Board) {
	state := board.GetState()
	for state.Running() {
		print <- *board
		select {
		case move := <-getmove:
			err := board.Apply(move)
			if err != nil {
				report <- err
			}
		case <-abort:
			done <- *board
			return
		}
		state = board.GetState()
	}
	done <- *board
}
