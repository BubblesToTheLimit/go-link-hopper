package main

import (
    "fmt"
    "os"
    "storage"
    "server"
    "validator"
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
    case "validate":
        if len(os.Args) < 3 {
            fmt.Print("No url given")
            return
        }

        var dto = validator.Validate{
            ExternalId: 1,
            Url: os.Args[2],
            OsVersion: "2.3",
            Timeout: 2,
        }

        var res = validator.Validator(dto)
        fmt.Print("Target: " + res.Target)
    default:
        fmt.Print("Invalid command given\n")
    }
}
