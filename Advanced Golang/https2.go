
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
)

type TimeDateApi struct {
    Datetime  string `json:"datetime"`
    Timezone  string `json:"timezone"`
    Weeknumber int64  `json:"week_number"`
}

type GeoIP struct{
  Country string `json:"country_name"`
  State string `json:"state_prov"`
  City string `json:"city"`
  Pincode string `json:"zipcode"`
  Latitude string `json:"latitude"`
  Longitude string `json:"longitude"`
}

type CityWeather struct{
  Name string `json:"name"`
  Coord struct{
    Lat float64 `json:"Lat"`
    Lon float64 `json:"lon"`
  }
  Weather []struct{
    Description string `json:"description"`
  }
  Main struct{
    Temp float64 `json:"temp"`
    Min float64 `json:"temp_min"`
    Max float64 `json:"temp_max"`
    Humidity int64 `json:"humidity"`
  }

}

func getTimeData(timezone string) (*TimeDateApi, error) {
    url := fmt.Sprintf("http://worldtimeapi.org/api/timezone/%s", timezone)

    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        panic(fmt.Sprintf("unexpected status code: %v", resp.Status))
    }

    var data TimeDateApi
    err = json.NewDecoder(resp.Body).Decode(&data)
    if err != nil {
        return nil, fmt.Errorf("failed to decode JSON: %v", err)
    }
    return &data, nil
}

func getGeoData(ip string) (*GeoIP, error){
  url := fmt.Sprintf("https://api.ipgeolocation.io/ipgeo?apiKey=b46376b370b44554b5c34a59eb560cd1&ip=%s", ip)
  resp, err := http.Get(url)
  if err != nil{
    return nil, err
  }
  defer resp.Body.Close()
  if resp.StatusCode != http.StatusOK {
    return nil, fmt.Errorf("unexpected status code: %v", resp.Status)
  }

  var output GeoIP
  err = json.NewDecoder(resp.Body).Decode(&output)
  if err!=nil{
    fmt.Println("error somewhere")
  }
  return &output, nil

}

func getWeatherData(city string) (*CityWeather,error){
  const api = "d9651e9c9a71679c8a59f5fd2f1d1f14"
  url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s",city, api)
  //check for server side problems like server down or no internet connection
  resp, err := http.Get(url)
  if err!=nil{
    return nil, err
  }

  defer resp.Body.Close()

  if resp.StatusCode!=http.StatusOK{
     return nil, fmt.Errorf("unexpected status code: %v", resp.Status)
  }

  var wea_output CityWeather
  err = json.NewDecoder(resp.Body).Decode(&wea_output)
  if err!=nil{
    fmt.Println("error somewhere")
  }
  return &wea_output, nil
}



func main() {
  var timezone string
  var ip string
  var choice int64
  fmt.Println("Enter choice. 1: timezone details, 2: GeoLocation using IP, 3: Weather of a city")
  fmt.Scanln(&choice)


  if choice==1{
    fmt.Println("Enter timezone. (eg:Asia/Kolkata)")
  fmt.Scanln(&timezone)
  data, err := getTimeData(timezone)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }
  fmt.Printf("Current date and time in %s: %s\n", timezone, data.Datetime)
  fmt.Println("Current week is:", data.Weeknumber)
  }
  

  if choice==2{
    fmt.Println("Enter your IP (you can find your IP from https://www.myip.com/ or its likes)")
    fmt.Scanln(&ip)
    data, err := getGeoData(ip)
    if err != nil{
      fmt.Println("yet another error")
      return
    }
  lat, err := strconv.ParseFloat(data.Latitude, 64)
  if err != nil {
    fmt.Println("Error parsing latitude:", err)
    return
  }
  lon, err := strconv.ParseFloat(data.Longitude, 64)
  if err != nil {
    fmt.Println("Error parsing longitude:", err)
    return
    }
    fmt.Printf("Country, State and City are %s, %s and %s respectively", data.Country, data.State, data.City)
    fmt.Println()
    fmt.Printf("Your latitude and longitude are: %f and %f respectively", lat, lon)
    fmt.Println()
    fmt.Println("Lastly, your Pincode is: ", data.Pincode)

  }

  if choice == 3{
    var city string
    fmt.Println("Enter name of city. (eg: London)")
    fmt.Scanln(&city)
    data, err := getWeatherData(city)
    if err!=nil{
      fmt.Println("You fucked up somewhere")
    }

    fmt.Printf("Country: %s", data.Name)
    fmt.Println()
    fmt.Printf("Latitude: %f", data.Coord.Lat)
    fmt.Println()
    fmt.Printf("Longitude: %f", data.Coord.Lon)
    fmt.Println()
    fmt.Printf("Weather descritpion: %s", data.Weather[0].Description)
    fmt.Println()
    fmt.Printf("Temperature: %f", data.Main.Temp)
    fmt.Println()
    fmt.Printf("Min. temp: %f", data.Main.Min)
    fmt.Println()
    fmt.Printf("Max. temp: %f", data.Main.Max)
    fmt.Println()
    fmt.Println("Humidity:", data.Main.Humidity,"%")
    

  }
}




