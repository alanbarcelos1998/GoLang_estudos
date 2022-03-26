package services

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	if Sum(2, 2) != 4 {
		t.Error(" 2 + 2 must be 4")
	}
}

func TestMult(t *testing.T) {
	require.Equal(t, 4, Multiply(2, 2))
}

// para rodar todos os testes basta fazer na pasta principal do pacote o seguinte comando no terminal:
// go test ./...

// para saber quantos porcento das funções foram cobertas pelo teste, fazemos:
// go test -cover ./...

/*
	para
*/
