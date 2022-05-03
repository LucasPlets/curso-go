package casamento

import (
	"encoding/json"
	"errors"
	"github.com/eminetto/curso-go/festas"
)

//Casorio resultado do calculo
type Casorio struct {
	TPessoas int `json:"tpessoas"`
	TRefri   int `json:"trefri"`
	TAgua    int `json:"tagua"`
	TSuco    int `json:"tsuco"`
	TChoop   int `json:"tchoop"`
	TChampanhe int `json: "tchampanhe`
}
func(c Casorio) ToJSON() ([]byte, error){
	return json.Marshal(c)
}
//Servide define um serviço de casorio
type Service struct{}

//NewCasamento cria um novo casorio
func NewCasamento() Service{
	return Service{}
}

//Calcula os itens do casorio
func (s Service) Calcula(p festas.Parametros) (festas.Resultado, error) {
	casamento := Casorio{}
	if p.Adultos == 0 {
		return casamento, errors.New("Sua festa não tem convidados")
	}

	casamento.TPessoas = p.Adultos + p.Criancas
	casamento.TRefri = (p.Adultos*400 + p.Criancas*200) / 1000
	casamento.TAgua = (p.Adultos*700 + p.Criancas*300) / 1000
	casamento.TSuco = (p.Adultos*200 + p.Criancas*100) / 1000
	casamento.TChoop = (p.Adultos * 2000) / 1000
	casamento.TChampanhe = ((p.Adultos/8) *750)/1000

	return casamento, nil

}
