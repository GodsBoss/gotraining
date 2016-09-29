package main

import (
  "encoding/json"
  "fmt"
  "github.com/godsboss/gotraining/jsonconnect/config"
  "github.com/godsboss/gotraining/jsonconnect/data"
  "io"
  "net"
)

func main() {
  listener, err := net.Listen("tcp", "127.0.0.1:" + config.PORT)
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
  reader, writer := io.Pipe()
  go func() {
    encodeRandomDataInto(writer)
  }()
  go func() {
    defer connection.Close()
    io.Copy(connection, reader)
  }()
}

func encodeRandomDataInto(writer *io.PipeWriter) {
  defer writer.Close()
  encoder := json.NewEncoder(writer)
  for i := 0; i<5; i++ {
    encoder.Encode(data.GetRandom())
  }
}
