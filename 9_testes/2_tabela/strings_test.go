package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte %s) - indices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	t.Parallel()

	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Lucas de Souza", "Lucas", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Lucas", "c", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa de dados: *v", teste) // go test -v (Executar testes com maior verbosidade)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
