package entity

type Cep struct {
	Cep          string `json:"cep"`
	Log          string `json:"logradouro"`
	Comp         string `json:"complemento"`
	Neighborhood string `json:"bairro"`
	Locale       string `json:"localidade"`
	UF           string `json:"uf"`
	IBGE         string `json:"ibge"`
	DDD          string `json:"ddd"`
}
