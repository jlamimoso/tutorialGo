package main

import (
	"fmt"
	"github.com/jlamimoso/tutorialGo/misc"
	"github.com/jlamimoso/tutorialGo/stringutil"
)

const (
	p1 = 3
	p2 = 4
)

func main() {
	fmt.Printf(stringutil.Reverse("987654321") + "\n")
	fmt.Println("comparar", p1, " > ", p2, " -> ", misc.Comp(p1, p2))
	x, y := misc.Swap("teste", "mimoso")
	fmt.Println(x, y)
	fmt.Println(misc.SomaTudo(10))
}
