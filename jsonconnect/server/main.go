package main

import (
  "fmt"
  "github.com/godsboss/gotraining/jsonconnect/config"
  "net"
)

func main() {
  listener, err := net.Listen("tcp", "127.0.0.1:" + string(config.PORT))
  if err != nil {
    fmt.Println("Could not open socket connection: " + err.Error())
    return
  }
  for {
    connection, err := listener.Accept()
    if err != nil {
      fmt.Println("Error while accepting connections: " + err.Error())
    } else {
      fmt.Println("Accepted connection from: " + connection.RemoteAddr().String())
      go handleConnection(connection)
    }
  }
}

func handleConnection(connection net.Conn) {
  connection.Close()
}
