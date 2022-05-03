package enderecos_test

import (
	"introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenarioDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos imigrantes", "Rodovia"},
		// {"Praça sete", "Tipo inválido"},
		{"RUA DOS BOBOS", "Rua"},
		// {"Alameda", "Tipo inválido"},
		// {"", "Tipo inválido"},
	}

	for _, cenario := range cenarioDeTeste {
		tipoDeEnderecoRecebido := enderecos.TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}
