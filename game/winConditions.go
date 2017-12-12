package game

func WinConditions() [][]Move {
	cur := 0
	max := 2*12 + 25 + 3*4
	res := make([][]Move, max)
	getSquare := func(x, y int) []Move {
		s := make([]Move, 4)
		s[0] = Move{x, y}
		s[1] = Move{x, y + 1}
		s[2] = Move{x + 1, y}
		s[3] = Move{x + 1, y + 1}
		return s
	}
	getHorLine := func(x, y int) []Move {
		s := make([]Move, 5)
		for i := 0; i < 5; i++ {
			s[i] = Move{x + i, y}
		}
		return s
	}
	getVerLine := func(x, y int) []Move {
		s := make([]Move, 5)
		for i := 0; i < 5; i++ {
			s[i] = Move{x, y + i}
		}
		return s
	}
	type getfunc func(int, int) []Move
	getThings := func(f getfunc, xb int, yb int) {
		for x := 0; x < xb; x++ {
			for y := 0; y < yb; y++ {
				res[cur] = f(x, y)
				cur++
			}
		}
	}
	symfunc := func(sym, x, y int) Move {
		if sym/2 == 1 {
			x = 5 - x
		}
		if sym%2 == 1 {
			y = 5 - y
		}
		return Move{x, y}
	}

	getOuter := func(sym int) {
		s := make([]Move, 3)
		s[0] = symfunc(sym, 0, 0)
		s[1] = symfunc(sym, 0, 1)
		s[2] = symfunc(sym, 1, 0)
		res[cur] = s
		cur++
	}
	getInner := func(sym int) {
		s := make([]Move, 3)
		s[0] = symfunc(sym, 2, 2)
		s[1] = symfunc(sym, 2, 1)
		s[2] = symfunc(sym, 1, 2)
		res[cur] = s
		cur++
	}
	getDiag := func(sym int) {
		s := make([]Move, 3)
		s[0] = symfunc(sym, 1, 1)
		s[1] = symfunc(sym, 2, 0)
		s[2] = symfunc(sym, 0, 2)
		res[cur] = s
		cur++
	}

	getSpecial := func(sym int) {
		getOuter(sym)
		getInner(sym)
		getDiag(sym)
	}

	for i := 0; i < 4; i++ {
		getSpecial(i)
	}

	getThings(getSquare, 5, 5)
	getThings(getVerLine, 6, 2)
	getThings(getHorLine, 2, 6)

	return res
}
