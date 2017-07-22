package browser

import (
    "context"
    "log"
    "time"
    "fmt"

    "github.com/knq/chromedp/runner"
    cdp "github.com/knq/chromedp"
    "github.com/knq/chromedp/cdp/network"
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

func NewBrowser(agent string, country string) *Client {
    var err error
    var proxy Proxy
    var chrome *cdp.CDP

    if country != "" {
        proxy = ByCountry(country)
    }

    // Create Context
    client := new(Client)
    ctxt, cancel := context.WithCancel(context.Background())

    // Create chrome instance
    var proxyOption = runner.Proxy(fmt.Sprintf("%s://%s:%s@%s:%s", proxy.Protocol, proxy.Credentials.User, proxy.Credentials.Password, proxy.Host, proxy.Port))
    var agentOption = runner.UserAgent(agent)

    if country == "" {
        chrome, err = cdp.New(ctxt, cdp.WithLog(log.Printf), cdp.WithRunnerOptions(agentOption))  // For headless use cdp.WithRunnerOptions(runner.Flag("headless", true) as third parameter
    } else {
        chrome, err = cdp.New(ctxt, cdp.WithRunnerOptions(agentOption, proxyOption))  // For headless use cdp.WithRunnerOptions(runner.Flag("headless", true) as third parameter
    }

    if err != nil {
        log.Fatal(err)
    }

    network.Enable()
    network.SetRequestInterceptionEnabled(true)

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