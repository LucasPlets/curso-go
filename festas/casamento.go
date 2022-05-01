package festas

//Parametros da funcao CalculaChurrascp
type Parametros struct{
	Adultos int
	Criancas int
	NaoAlcoolicos int
	Duracao int
	Bar bool
}

//Casorio resultado do calculo
type Casorio struct{
	TPessoas int
	TRefri int
	TAgua int
	TSuco int
	TChoop int
}
//Calcula os itens do casorio
func CalculaCasorio(p Parametros) Casorio{
	casamento := Casorio{}
	casamento.TPessoas = p.Adultos + p.Criancas
	casamento.TRefri = (p.Adultos*400 + p.Criancas*200)/1000
	casamento.TAgua = (p.Adultos*700 + p.Criancas*300)/1000
	casamento.TSuco = (p.Adultos*200 + p.Criancas*100)/1000
	casamento.TChoop = (p.Adultos*2000)/1000
	
	return casamento
}