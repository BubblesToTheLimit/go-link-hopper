package validator

import (
    "browser"
    "time"
)

type Result struct {
    Id int
    Url string
    Target string
    Trace []string `xorm:`
    CreatedAt time.Time `xorm:"created"`
    Error error
}

type Validate struct {
    Id int
    Url string
    Country string
    OsVersion string
    Timeout int
}

func Validator(dto Validate) Result {
    var agent = generateUserAgent(dto.OsVersion)
    var client = browser.NewBrowser(agent)

    client.Timeout = time.Duration(dto.Timeout) * time.Second

    var navRes = client.Navigate(dto.Url)

    var res = Result{
        Id: dto.Id,
        Url: dto.Url,
        Target: navRes.Target,
        Trace: navRes.Trace,
        CreatedAt: time.Now(),
        Error: navRes.Error,
    }

    client.ShutDown()

    return res
}

func generateUserAgent(osVersion string) string {
    return "Mozilla/5.0 (Linux; Android " + osVersion + "; SM-G920T Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.78 Mobile Safari/537.36";
}