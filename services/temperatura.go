package services

import (
	"afonsojota/go-expert-cep-labs/adapters"
	"afonsojota/go-expert-cep-labs/dto"
	"afonsojota/go-expert-cep-labs/errors"
	"net/url"
)

func SearchWeather(zipCode string) (*dto.TemperaturaResponse, error) {
	resViaCep, err := adapters.SearchZipCode(zipCode)
	if resViaCep.Cep == "" {
		return nil, errors.NotFoundZipCode
	}

	if err != nil {
		return nil, errors.UnableToRetrieveZipCode
	}

	city := url.QueryEscape(resViaCep.Locale)
	resWeather, err := adapters.GetWeather(city)
	if err != nil {
		return nil, errors.UnableToRetrieveWeather
	}

	return &dto.TemperaturaResponse{
		TempC: resWeather.Current.TempC,
		TempK: getTemperatureKelvin(resWeather.Current.TempC),
		TempF: getTemperatureFahrenheit(resWeather.Current.TempC),
	}, nil
}

func getTemperatureFahrenheit(celsius float64) float64 {
	return (celsius * 1.8) + 32
}

func getTemperatureKelvin(celsius float64) float64 {
	return celsius + 273
}
