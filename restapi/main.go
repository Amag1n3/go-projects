package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type friend struct{
  Name string `json:"name"`
  Batch int `json:"batch"`
}

var friends = []friend{
  {Name: "Amogh", Batch: 2023},
  {Name: "Pathak", Batch: 2023},
}


func main(){
  var choice int
  fmt.Println("Enter choice. 1: Add, 2: Search")
  fmt.Scanln(&choice)
  router := gin.Default()
  router.GET("/friends", func(c *gin.Context){
    c.IndentedJSON(http.StatusOK, friends)
  })

  switch choice{
  case 1:
    addFriend()
  case 2:
    queryFriend()
  }

  router.Run(":8080")
}

func queryFriend(){
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Enter Friend's Name: ")
  Qname,_ := reader.ReadString('\n')
  Qname = strings.TrimSpace(Qname)


  fmt.Println("Enter Friend's Batch: ")
  QbatchStr,_ := reader.ReadString('\n')
  QbatchStr = strings.TrimSpace(QbatchStr)


  Qbatch, err := strconv.Atoi(QbatchStr)
  if err != nil{
    fmt.Printf("Error changing batch to int: %v\n", err)
  }

  for _, friend := range friends{
    if Qname == friend.Name && Qbatch == friend.Batch{
      fmt.Println(friend)
    }
  }

}


func addFriend(){
  reader := bufio.NewReader(os.Stdin)

  fmt.Println("Enter New Friend's Name: ")
  name,_ := reader.ReadString('\n')
  name = strings.TrimSpace(name)


  fmt.Println("Enter New Friend's Batch: ")
  batchStr,_ := reader.ReadString('\n')
  batchStr = strings.TrimSpace(batchStr)


  batch, err := strconv.Atoi(batchStr)
  if err != nil{
    fmt.Printf("Error changing batch to string: %v\n", err)
  }
  friends = append(friends, friend{Name: name, Batch: batch})
  fmt.Println("Friend added successfully!")
}
