package funcionario

type Dados struct {
	Nome, Endereco, Depto string
}

var (
	Depto   = []string{"depto1", "depto2", "depto3", "depto4"}
	pessoal = map[int]Dados{
		1: {"func1", "End1", "Depto1"},
		2: {"func2", "End2", "Depto2"},
		3: {"func3", "End3", "Depto3"},
		4: {"func4", "End4", "Depto4"},
		5: {"func5", "End5", "Depto5"},
	}
)

func GetFunc(id int) Dados {
	return pessoal[id]
}

func GetDepto(id int) string {
	return Depto[id-1]
}
