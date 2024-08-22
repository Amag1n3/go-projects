package main

import(
  "fmt"
)

func main(){
  res, err := https.Get("https://api.weatherbit.io/v2.0/current?access_key=%s&query=%s")
  if err!=nil{
    panic(err)
  }
  defer res.Body.close()

  if 
}


