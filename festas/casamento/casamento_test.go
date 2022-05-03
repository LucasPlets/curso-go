package casamento

import ("testing"
	"github.com/eminetto/curso-go/festas"
		"github.com/stretchr/testify/assert"
	
)

func TestCasamentoSemPessoa(t *testing.T) {
	s:= NewCasamento()
	_, err :=s.Calcula(festas.Parametros{
		Adultos:       0,
		Criancas:      10,
		NaoAlcoolicos: 0,
		Duracao:       8,
		Bar:           true,
	})
	assert.Equal(t, err.Error(), "Sua festa não tem convidados")
	
}
func TestDadosCasamento(t *testing.T) {
	s := NewCasamento()
	ch , err := s.Calcula(festas.Parametros{
		Adultos:       110,
		Criancas:      30,
		NaoAlcoolicos: 0,
		Duracao:       8,
		Bar:           true,
	})
	assert.Nil(t, err)
	

	esperado := Casorio{
	TPessoas : 140,
	TRefri : 50,
	TAgua : 86,   
	TSuco : 25,   
	TChoop : 220,
	TChampanhe: 9,  
	}
	assert.Equal(t, ch, esperado)
	
}