package models

type Players struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	History     string `json:"history"`
	ShirtNumber int    `json:"shirt_number"`
}

var Player []Players
