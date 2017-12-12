package ai

import (
	"github.com/qutrace/qanGo/game"
)

func GetRating(b *Board, player bool) BoardState {
	cells := game.WinCondonditions() //array of all winconditions
	const (
		cwin0 = iota
		cwin1
		cdraw
		cforce1
		cforceable1
		cforce0
		cforceable0
		cboth
		cpot0
		cpot1
	)
	getnum := func(p0, p1, l int) int {
		if p1 == l {
			return cwin1
		}
		if p0 == l {
			return cwin0
		}
		if p0 == 0 && p1 == 0 {
			return cboth
		}
		if p1 == 0 {
			if p0 == l-1 {
				return cforce0
			}
			if p0 == l-2 {
				return cforceable0
			}
			return cpot0
		}
		if p0 == 0 {
			if p1 == l-1 {
				return cforce1
			}
			if p1 == l-2 {
				return cforceable1
			}
			return cpot1
		}
		if p0+p1 == 0 {
			return cboth
		}
		if (p0 > 0) && (p1 > 0) {
			return cdraw
		}
		return cboth
	}
	nums := [11]int{} //win0, win1, draw, force1, forcable1, force0, forcable0, both
	rateCond := func(c []*Cell) int {
		l := len(c)
		p1 := 0
		p0 := 0
		for _, cell := range c {
			if cell != nil {
				if cell.Player {
					p1++
				} else {
					p0++
				}
			}
		}
		return getnum(p0, p1, l)
	}
	for _, cond := range cells {
		cells := MovesToCells(b, cond)
		nums[rateCond(cells)]++
	}

	if nums[cwin0] > 0 {
		if !player {
			return BoardState(8000)
		}
		return BoardState(0)
	}
	if nums[cwin1] > 0 {
		if player {
			return BoardState(8000)
		}
		return BoardState(0)
	}
	if nums[cdraw] > 60 {
		return BoardState(1)
	}
	//fmt.Println(nums)
	offset := 4000
	num := nums[cforce0]*8*8 + nums[cforceable0]*8 + nums[cpot0] - (nums[cforce1]*8*8 + nums[cforceable1]*8 + nums[cpot1])
	if player {
		num = num * -1
	}

	return BoardState(num + offset)
}
