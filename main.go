package main
import (
	"net/http"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"
	"errors"
	"os"

)

func main() {

	app := CreateWeatherApp("config.json")

	if app.Port == 0 {
		fmt.Println("port is a required parameter in config.json")
		os.Exit(-1)
	}
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/weather",app.handleGetWeather)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", app.Port)))
}


func (w WeatherApp) handleGetWeather(c echo.Context) error {
	params := c.QueryParams()
	unit := ""
	if len(params["unit"]) == 1 {
		unit = params["unit"][0]
	}
	zipcode := ""
	if len(params["zipcode"]) == 1 {
		zipcode = params["zipcode"][0]
	}
	city := ""
	if len(params["city"]) == 1 {
		city = params["city"][0]
	}
	if zipcode == "" && city == "" {
		c.String(http.StatusBadRequest, "zipcode or city required")
		return nil
	}
	if zipcode != "" && city != "" {
		c.String(http.StatusBadRequest, "must specify either city or zipcode")
		return nil
	}

	wr, err := w.GetWeather(zipcode, city, unit)
	if err != nil {
		if err.Error() == "invalid unit" {
			c.String(http.StatusBadRequest,
			"Please enter a valid unit (celsius, fahrenheit, imperial, metric, kelvin)")
			return nil
		}
		if err.Error() == "city/zip not found" {
			c.String(http.StatusNotFound, "City/zip not found")
			return nil
		}
		c.NoContent(http.StatusInternalServerError)
		return err
	}
	c.JSON(http.StatusOK, wr)
	return nil
}

func (w WeatherApp) GetWeather(zipcode, city, unit string) (WeatherResponse, error) {
	if unit != "" {
		unit = convertUnit(unit, w.Conversions)
	}
	if unit != "" && !isValidUnit(unit, w.Units) {
		return WeatherResponse{}, errors.New("invalid unit")
	}
	var url strings.Builder
	fmt.Fprintf(&url, "https://%s", w.BaseURL)
	fmt.Fprintf(&url, "?APPID=%s", w.Apikey)
	if unit != "" {
		fmt.Fprintf(&url, "&units=%s", unit)
	}
	if zipcode == "" {
		fmt.Fprintf(&url, "&q=%s", city)
	}
	if city == "" {
		fmt.Fprintf(&url, "&zip=%s", zipcode)
	}
	fmt.Printf("url: %s\n", url.String())
	resp, err := http.Get(url.String())
	if err != nil {
		return WeatherResponse{}, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return WeatherResponse{}, errors.New("city/zip not found")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return WeatherResponse{}, err
	}
	
	var ret WeatherResponse
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return WeatherResponse{}, err
	}
	if unit == "" {
		ret.Temp.Unit = "kelvin"
	} else {
		ret.Temp.Unit = unit
	}
	return ret, nil
}

func CreateWeatherApp(configFile string) WeatherApp {
	dat, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("failed to read file %s: %s", configFile, err)
		os.Exit(-1)
	}
	var config WeatherApp
	err = json.Unmarshal(dat, &config)
	if err != nil {
		fmt.Printf("failed to unmarshal config json %s", err)
		os.Exit(-1)
	}
	return config
}
func isValidUnit(unit string, validUnits []string) bool {
	for _, validUnit := range validUnits {
		if unit == validUnit {
			return true
		}
	}
	return false
}

func convertUnit(unit string, conversions [][]string) string {
	for _, pair := range conversions {
		if len(pair) == 2 {
			if unit == pair[0] {
				return pair[1]
			}
		}
	}
	return unit
}
