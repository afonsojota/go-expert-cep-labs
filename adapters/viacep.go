package adapters

import (
	"afonsojota/go-expert-cep-labs/entity"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func SearchZipCode(zipCode string) (*entity.Cep, error) {
	resp, err := http.Get("http://viacep.com.br/ws/" + zipCode + "/json/")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var c entity.Cep
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
