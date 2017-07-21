package browser

type Credentials struct {
    User string
    Password string
}

type Proxies struct {
    Credentials Credentials
    Hosts struct {

    }
}

type Proxy struct {
    Protocol string
    Host string
    Port int
    Credentials Credentials
}

func NewProxy() *Proxy {
    var proxy = new(Proxy)

    proxy.Protocol = "socks5"
    proxy.Port = 65336

    return proxy
}

func (proxy *Proxy) loadByCountry(country string) {

}