package models

type Players struct {
	Name        string `json:"name"`
	History     string `json:"history"`
	ShirtNumber int    `json:"shirt_number"`
}

var Player []Players
