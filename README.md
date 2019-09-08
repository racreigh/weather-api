# weather-api

API for getting the weather from https://openweathermap.org/current

Supports getting current weather data by zipcode or city name.

Temperature unit can be specified, defaults to Kelvin.

Using go modules is recommended.

## Configuration

`apikey` is required and not provided.

`conversions` allows you to convert units to other units under the covers as the openweather API does no support names like "celsius"

`port` allows you to configure the port the server listens on

`units` defines the allowed units (at this time the openweather API only supports "imperial", and "metric")

`url` defines the endpoint this server retrieves its data from



## Example Requests:

GET <hostname>/weather?city=Seattle&unit=imperial
Response:
```
{
  "coord": {
    "lat": 47.6,
    "lon": -122.33
  },
  "weather": [
    {
      "main": "Clouds",
      "description": "overcast clouds"
    }
  ],
  "main": {
    "Unit": "imperial",
    "temp": 66,
    "temp_max": 68,
    "temp_min": 64
  },
  "name": "Seattle"
}
```

GET <hostname>/weather?zipcode=98058&unit=fahrenheit (fahrenheit gets converted to "imperial" if `conversions` is defined as in the sample config)
Response:
```
{
  "coord": {
    "lat": 47.45,
    "lon": -122.12
  },
  "weather": [
    {
      "main": "Clouds",
      "description": "overcast clouds"
    }
  ],
  "main": {
    "Unit": "imperial",
    "temp": 66.04,
    "temp_max": 68,
    "temp_min": 64
  },
  "name": "Renton"
}
```
