package main

import (
    "browser"
    "fmt"
)

func main() {
    fmt.Print("Start")

    var client = browser.NewBrowser(3)

    client.Navigate("http://google.com")

    var site = client.GetLocation()

    client.ShutDown()

    fmt.Print("Reached site " + site)
}
