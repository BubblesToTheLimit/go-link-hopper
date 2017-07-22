package main

import (
    "fmt"
    "os"
    "storage"
    "server"
)

func main() {
    fmt.Printf("Let's GO!\n")

    if len(os.Args) < 2 {
        fmt.Print("No command given!\n")
        return
    }

    switch command := os.Args[1]; command {
    case "init":
        storage.Init()
    case "run":
        server.Init()
    default:
        fmt.Print("Invalid command given\n")
    }
}
