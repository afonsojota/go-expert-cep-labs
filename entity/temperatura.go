package entity

type Temperatura struct {
	Current *Current `json:"current"`
}

type Current struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
}
