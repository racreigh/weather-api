package main

import (
	//"encoding/json"
)

type WeatherResponse struct {
	Coords Coords `json:"coord"`
	Weather []Weather `json:"weather"`
	Temp Temperature `json:"main"`
	City string `json:"name"`
}

type Coords struct {
	Latitude float32 `json:"lat"`
	Longitude float32 `json:"lon"`
}

type Weather struct {
	Conditions string `json:"main"`
	Description string `json:"description"`
}

type Temperature struct {
	Unit string 
	TempCurrent float32 `json:"temp"`
	TempHigh float32 `json:"temp_max"`
	TempLow float32 `json:"temp_min"`
}

type WeatherApp struct {
	BaseURL string `json:"url"`
	Apikey string `json:"apikey"`
	Units []string `json:"units"`
	Port int `json:"port"`
}
