package main

import (
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

func cats() ([]string, error) {
	data, err := ioutil.ReadFile("./assets/data/cats.json")
	if err != nil {
		return nil, err
	}

	var cats []string
	err = jsoniter.Unmarshal(data, &cats)
	if err != nil {
		return nil, err
	}

	return cats, nil
}
