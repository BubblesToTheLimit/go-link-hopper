package validator

import (
    "browser"
    "fmt"
)

func Validator(url string) string {
    fmt.Print("Start")

    var client = browser.NewBrowser(3)

    //client.Navigate("http://google.com")
    client.Navigate(url)

    var site = client.GetLocation()

    client.ShutDown()

    fmt.Print("Reached site " + site)
    return site
}
