package browser

import (
    "context"
    "log"
    "time"

    "github.com/knq/chromedp/runner"
    cdp "github.com/knq/chromedp"
)

type Result struct {
    Target string
    Trace []string
    Error error
}

type Client struct {
    Context context.Context
    Client *cdp.CDP
    Cancel context.CancelFunc
    Timeout time.Duration
    Url string
    Res Result
}

func NewBrowser(agent string) *Client {
    var err error

    // Create Context
    client := new(Client)
    ctxt, cancel := context.WithCancel(context.Background())

    // Create chrome instance
    chrome, err := cdp.New(ctxt, cdp.WithLog(log.Printf), cdp.WithRunnerOptions(runner.UserAgent(agent)))  // For headless use cdp.WithRunnerOptions(runner.Flag("headless", true) as third parameter
    if err != nil {
        log.Fatal(err)
    }

    client.Context = ctxt
    client.Client = chrome
    client.Cancel = cancel

    return client
}

func (client *Client) Navigate(site string) Result {
    var err = client.Client.Run(client.Context, cdp.Navigate(site))

    if err != nil {
        client.Res.Error = err
        return client.Res
    }

    err = client.Client.Run(client.Context, cdp.Sleep(client.Timeout))
    if err != nil {
        client.Res.Error = err
        return client.Res
    }

    // Get location
    err = client.Client.Run(client.Context, cdp.Location(&client.Res.Target))
    if err != nil {
        client.Res.Error = err
        return client.Res
    }

    return client.Res
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