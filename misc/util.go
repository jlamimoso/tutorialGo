package misc

import "time"

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
	for {
		if i += 1; i > x {
			return cont
		}
		cont += i
	}
}

func WhenIs(dia time.Weekday) string {
	hoje := time.Now().Weekday()
	switch dia {
	case hoje + 0:
		return "e hoje !"
	case hoje + 1:
		return "e amanha !"
	case hoje + 2:
		return "depois de amanha !"
	default:
		return "esta longe !"
	}
}
