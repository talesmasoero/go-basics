package structsaux

type Pessoa struct {
	Nome     string
	Idade    uint8
	Endereco Endereco
}

type Endereco struct {
	Cidade string
	Numero uint8
}
