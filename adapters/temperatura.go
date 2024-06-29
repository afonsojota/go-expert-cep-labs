package adapters

import (
	"afonsojota/go-expert-cep-labs/entity"
	"afonsojota/go-expert-cep-labs/errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetWeather(locale string) (*entity.Temperatura, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?q=%s&key=0895eea478b24a31892212515242801", locale)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, errors.UnableToRetrieveWeather
	}
	defer resp.Body.Close()

	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var c entity.Temperatura
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
