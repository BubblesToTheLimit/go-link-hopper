package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "validator"
    "io/ioutil"
    "strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // trying to read parameter
    r.ParseForm()

    var url = r.Form.Get("url")
    if url == "" {
       http.Error(w, "invalid url provided: " + url, http.StatusBadRequest)
    }

    var os = r.Form.Get("os")
    if os == "" {
        os = "2.3"
    }

    var country = r.Form.Get("country")

    var id = -1
    var sId = r.Form.Get("id")
    if sId != "" {
        i, err := strconv.ParseInt(sId, 10, 32)
        if err != nil {
            http.Error(w, "non numerical id provided: " + sId, http.StatusBadRequest)
            return
        }

        id = int(i)
    }

    var timeout = 8
    var sTimeout = r.Form.Get("timeout")
    if sTimeout != "" {
        i, err := strconv.ParseInt(sTimeout, 10, 32)
        if err != nil {
            http.Error(w, "non numerical timeout provided: " + sTimeout, http.StatusBadRequest)
            return
        }

        timeout = int(i)
    }

    var dto = validator.Validate{
        Id: id,
        Url: url,
        OsVersion: os,
        Timeout: timeout,
        Country: country,
    }

    var result = validator.Validator(dto)
    if result.Error != nil {
        http.Error(w, "Validator error: " + result.Error.Error(), http.StatusInternalServerError)
        return
    }

    jData, err := json.Marshal(result)
    if err != nil {
       http.Error(w, "Serialization error: " + err.Error(), http.StatusInternalServerError)
       return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jData)

    logResult(string(jData))
}

func logResult(text string) {
    dat, _ := ioutil.ReadFile("log.txt")

    stuff := []byte(string(dat) + text + "\n" )

    ioutil.WriteFile("log.txt", stuff, 0644)

}

func main() {
    fmt.Printf("Hello, friend")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
