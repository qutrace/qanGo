package game



func (b Board) String() string {
	res := ""
	for _, line := range b {
		for _, char := range line {
			if char == nil {
				res += "#"
			} else {
				switch *char {
					case true: res += "o"
					case false: res += "x"
				}
			}
			res += " "
		}
		res += "\n"
	}
	return res
}

func (s State) String() string {
	switch s {
		case Unknown:
		return "Unknown"
		case WinTrue:
		return "WinTrue"
		case WinFalse:
		return "WinFalse"
		case Draw:
		return "Draw"
	}
	return "state"
}
