package main

import (
	"encoding/json"
	"os"
)

func jsonFromFile(filename string) (any, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data any
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
