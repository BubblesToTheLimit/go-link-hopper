package main

import (
    "fmt"
    "os"
    "storage"
    "server"
)

func main() {
    fmt.Printf("Let's GO!")

    switch command := os.Args[1]; command {
    case "init":
        storage.Init()
    case "run":
        server.Init()
    default:
        fmt.Print("Invalid command given")
    }
}
