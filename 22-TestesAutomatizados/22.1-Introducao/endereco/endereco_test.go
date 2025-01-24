// O único caso em que podemos usar mais de uma pacote na mesma
// pasta, é quando for um pacote para testes (é opcional)
package endereco_test

import (
	. "basics/22-TestesAutomatizados/22.1-Introducao/endereco"
	"testing"
)

// TestTipoDeEndereco é um teste unitário básico (apenas um cenário)
func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Rua dos Bobos"
	enderecoEsperado := "Rua"
	enderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if enderecoRecebido != enderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperado: \"%s\" Recebido: \"%s\"",
			enderecoEsperado,
			enderecoRecebido,
		)
	}
}

type cenarioDeTeste struct {
	EnderecoInserido string
	EnderecoEsperado string
}

// TestTipoDeEndereco2 é um teste unitário "normal" (multiplos cenários)
func TestTipoDeEndereco2(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Apartamento 706", "Apartamento"},
		{"Casa 4", "Casa"},
		{"Bairro dos Pobres", "Bairro"},
		{"RUA DOS BOBOS", "Rua"},
		{"Lago Norte", "endereco invalido"},
		{"Torre de TV", "endereco invalido"},
		{"", "endereco invalido"},
	}

	for _, cenario := range cenariosDeTeste {
		enderecoRecebido := TipoDeEndereco(cenario.EnderecoInserido)
		if enderecoRecebido != cenario.EnderecoEsperado {
			t.Errorf("\nEsperado: %s\nRecebido: %s\n", cenario.EnderecoEsperado, enderecoRecebido)
		}
	}
}
