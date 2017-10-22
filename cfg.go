package main

import (
	"encoding/json"
	"os"
)

// Configuration represent conf.json
type Configuration struct {
	LoginURL string
	Username string
	Password string
	Serie    []string
}

func leggiCFG(cfgName string) (*Configuration, error) {

	file, err := os.Open(cfgName)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(file)
	Conf := Configuration{}
	err = decoder.Decode(&Conf)
	if err != nil {
		return nil, err
	}

	return &Conf, err
}
