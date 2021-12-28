package models

import (
	"encoding/json"
	"log"
)

type Directory struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Path  string `json:"path"`
	Files []File `json:"files"`
}

type File struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type ThetvdbLogin struct {
	Status string `json:"status"`
	Data   data   `json:"data"`
}

type data struct {
	Token string `json:"token"`
}

func (d *ThetvdbLogin) UnMarshal(data string) ThetvdbLogin {
	res := ThetvdbLogin{}
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
