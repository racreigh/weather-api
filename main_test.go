package main
import (
	"fmt"
	"testing"
	"gotest.tools/assert"
)
//Get should return a WeatherResponse
func TestGet(t *testing.T) { 
	w := CreateWeatherApp("config.json")
	resp, err := w.GetWeather("27617", "", "imperial")
	assert.NilError(t, err)
	fmt.Printf("%v",resp)
	assert.Equal(t, resp.City, "Raleigh")
}

//Get should return a WeatherResponse even if unit is blank
func TestDefaultUnit(t *testing.T) {
        w := CreateWeatherApp("config.json")
        resp, err := w.GetWeather("27617", "", "")
        assert.NilError(t, err)
        fmt.Printf("%v",resp)
        assert.Equal(t, resp.City, "Raleigh")
}

