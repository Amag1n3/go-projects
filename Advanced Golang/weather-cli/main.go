package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const apiKey = "b1471fc2c76d4085bcc316e81023aa78" // Replace with your Weatherbit API key

type WeatherbitData struct {
	Data []struct {
		CityName    string  `json:"city_name"`
		Temp        float64 `json:"temp"`
		Weather     struct {
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"data"`
}

func getWeather(city string) (*WeatherbitData, error) {
	url := fmt.Sprintf("https://api.weatherbit.io/v2.0/current?city=%s&key=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get weather data: %s", resp.Status)
	}

	var data WeatherbitData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("JSON decode error: %v", err)
	}
	return &data, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a city name")
		os.Exit(1)
	}

	city := os.Args[2]
	weather, err := getWeather(city)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if len(weather.Data) == 0 {
		fmt.Println("No weather data found for the provided city")
		os.Exit(1)
	}

	fmt.Printf("Weather in %s: %s, %.2f°C\n", weather.Data[0].CityName, weather.Data[0].Weather.Description, weather.Data[0].Temp)
}
