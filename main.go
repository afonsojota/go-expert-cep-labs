package main

import (
	"afonsojota/go-expert-cep-labs/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/api/temperatura", controllers.GetWeatherHandle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
