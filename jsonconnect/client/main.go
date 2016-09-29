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
  connection, err := net.Dial("tcp", "127.0.0.1:" + config.PORT)
  if err != nil {
    fmt.Println("Could not connect: " + err.Error())
    return
  }
  reader, writer := io.Pipe()
  go func() {
    defer connection.Close()
    defer writer.Close()
    io.Copy(writer, connection)
  }()
  decoder := json.NewDecoder(reader)
  var value data.Data
  for decoder.More() {
    err := decoder.Decode(&value)
    if err != nil {
      fmt.Println("Could not decode value: " + err.Error())
    } else {
      fmt.Printf("Received value %d.\n", value.Number)
    }
  }
  reader.Close()
}
