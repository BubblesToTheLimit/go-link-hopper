package browser

import (
    "storage"
    "config"
)

type Proxy struct {
    Protocol string
    Host string
    Port int
    Credentials config.Credentials
}

func ByCountry(country string) Proxy {
    data, err := storage.GetProxyByCountry(country)
    if err != nil {
        panic(err)
    }

    if data == nil {
        panic("Proxy for country " + country + " not found")
    }

    var proxy = config.ForProxy()

    return Proxy{
        Protocol: "socks5",
        Host: data.IP,
        Port: 61336,
        Credentials: proxy.Credentials,
    }
}