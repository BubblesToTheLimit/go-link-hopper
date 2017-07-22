package validator

import (
    "browser"
    "time"
    "storage"
)

type Validate struct {
    ExternalId int
    Url        string
    Country    string
    OsVersion  string
    Timeout    int
}

func Validator(dto Validate) storage.Result {
    var agent = generateUserAgent(dto.OsVersion)
    var client = browser.NewBrowser(agent, dto.Country)

    client.Timeout = time.Duration(dto.Timeout) * time.Second

    var navRes = client.Navigate(dto.Url)

    var res = storage.Result{
        ExternalId: dto.ExternalId,
        Url: dto.Url,
        Target: navRes.Target,
        Trace: navRes.Trace,
        Error: navRes.Error,
    }

    client.ShutDown()

    return res
}

func generateUserAgent(osVersion string) string {
    return "Mozilla/5.0 (Linux; Android " + osVersion + "; SM-G920T Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.78 Mobile Safari/537.36";
}