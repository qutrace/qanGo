package game

func rateBoard(b *Board) State {
	conditions := winConditions()
	for _, condition := range conditions {
		cells := movesToCells(b, condition)
		res := rateCells(cells)
		if res != Unknown {
			return res
		}
	}	
	if b.GetTurn() > 35 {
		return Draw
	}
	return Unknown
}

func rateCells(cells []*Cell) State {
	if len(cells) == 0 {
		return Unknown
	}
	if *cells[0] == nil {
		return Unknown
	}
	player := **cells[0]

	for _, cell := range cells {
		if *cell == nil {
			return Unknown
		}
		if **cell != player {
			return Unknown
		}
	}
	if player == true {
		return WinTrue
	}
	return WinFalse
}

func movesToCells(b *Board, moves []Move) []*Cell {
	length := len(moves)
	getter := cellGetter(b)
	cells := make([]*Cell, length)
	for i := 0; i < length; i++ {
		m := moves[i]
		cells[i] = getter(m.X,m.Y)
	}
	return cells
}

func cellGetter(b *Board) func(x, y int) *Cell {
	return func(x, y int) *Cell {
		return &(b[y][x])
	}
}

