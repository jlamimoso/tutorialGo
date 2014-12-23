package main

import (
	"fmt"
	"github.com/jlamimoso/tutorialGo/funcionario"
	"github.com/jlamimoso/tutorialGo/misc"
	"github.com/jlamimoso/tutorialGo/stringutil"
	"time"
)

const (
	p1 = 3
	p2 = 4
)

func main() {
	emp := func(id int) funcionario.Dados {
		return funcionario.Pessoal[id]
	}

	fmt.Printf(stringutil.Reverse("987654321") + "\n")
	fmt.Println("comparar", p1, " > ", p2, " -> ", misc.Comp(p1, p2))
	x, y := misc.Swap("teste", "mimoso")
	fmt.Println(x, y)
	fmt.Println(misc.SomaTudo(10))
	dia := time.Tuesday
	fmt.Printf("Quando e %s ? %s\n", dia.String(), misc.WhenIs(dia))
	id := 4
	fmt.Println("Quem e o ID ", id, "? ", funcionario.GetFunc(id))
	fmt.Println("Quem e o ID ", id, "? ", emp(id))

	D := &funcionario.Dados{}
	fmt.Println("Quem e o ID(1) ", id, "? ", D.GetFunc1(id))

	fmt.Println("Qual o depto do ", id, "? ", funcionario.GetDepto(id))

	i := 0
	for _, v := range funcionario.Depto {
		i += 1
		fmt.Printf("valor do depto do id %d=%s\n", (i), v)
	}
}
