package fileio

import (
	"encoding/json"
	"os"
)

func GetJson[T any](path string, t T) T {
	fileData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(fileData, &t)

	return t
}

func SaveJson[T any](path string, t T) {
	jsonString, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		panic(err)
	}

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	f.WriteString(string(jsonString))
}
