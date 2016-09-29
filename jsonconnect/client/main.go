package main

import (
  "fmt"
  "github.com/godsboss/gotraining/jsonconnect/config"
  "net"
)

func main() {
  _, err := net.Dial("tcp", "127.0.0.1:" + config.PORT)
  if err != nil {
    fmt.Println("Could not connect: " + err.Error())
    return
  }
}
