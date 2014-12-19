package misc

func Comp(x, y int) string {
	dif := x - y
	switch {
	case dif > 0:
		return "maior"
	case dif == 0:
		return "igual"
	}
	return "menor"
}

func Swap(x, y string) (string, string) {
	return y, x
}

func SomaTudo(x int) int {
	var i, cont = 0, 0
	for i < x {
		cont += i
		i++
	}
	return cont
}
