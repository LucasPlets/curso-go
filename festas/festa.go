package festas

//Parametros da funcao CalculaCasorio
type Parametros struct {
	Adultos       int  `json:"adultos"`
	Criancas      int  `json:"criancas"`
	NaoAlcoolicos int  `json:"naoalcoolicos"`
	Duracao       int  `json:"duracao"`
	Bar           bool `json:"bar"`
}

//Festa define oque é uma festa
type Festa interface {
	Calcula(Parametros) (Resultado,error)
}

//Resultado define oque é um resultado
type Resultado interface{
	ToJSON() ([]byte,error)
}