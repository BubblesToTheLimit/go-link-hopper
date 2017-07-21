package browser

import (
    "context"
    "log"
    "time"

    cdp "github.com/knq/chromedp"
)

type Client struct {
    Context context.Context
    Client *cdp.CDP
    Cancel context.CancelFunc
    Timeout time.Duration
}

func NewBrowser(timeout int) *Client {
    var err error

    // Create Context
    client := new(Client)
    ctxt, cancel := context.WithCancel(context.Background())

    // Create chrome instance
    chrome, err := cdp.New(ctxt, cdp.WithLog(log.Printf))  // For headless use cdp.WithRunnerOptions(runner.Flag("headless", true) as third parameter
    if err != nil {
        log.Fatal(err)
    }

    client.Context = ctxt
    client.Client = chrome
    client.Cancel = cancel
    client.Timeout = time.Duration(timeout) * time.Second

    return client
}

func (client *Client) Navigate(site string) bool {
    var err = client.Client.Run(client.Context, cdp.Navigate(site))

    if err != nil {
        log.Fatal(err)
        return false
    }

    err = client.Client.Run(client.Context, cdp.Sleep(client.Timeout))
    if err != nil {
        log.Fatal(err)
        return false
    }

    return true
}

func (client *Client) GetLocation() string {
    var site string

    var err = client.Client.Run(client.Context, cdp.Location(&site))
    if err != nil {
        log.Fatal(err)
        return ""
    }

    return site
}

func (client *Client) ShutDown() bool {
    var err error

    err = client.Client.Shutdown(client.Context)
    if err != nil {
        log.Fatal(err)
        return false
    }

    // wait for chrome to finish
    err = client.Client.Wait()
    if err != nil {
        log.Fatal(err)
        return false
    }

    client.Cancel()

    return true
}