package server

import (
    "net/http"
    "encoding/json"
    "validator"
    "io/ioutil"
    "strconv"
    "fmt"
    "storage"
    "log"
)

func handleValidation(w http.ResponseWriter, r *http.Request) {
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
        ExternalId: id,
        Url:        url,
        OsVersion:  os,
        Timeout:    timeout,
        Country:    country,
    }

    var result = validator.Validator(dto)
    if result.Error != nil {
        http.Error(w, "Validator error: " + result.Error.Error(), http.StatusInternalServerError)
        return
    }

    // Persist
    storage.Save(result)

    // Write result to html result as json
    jsonOutput(result, w)

    // Write result to log file
    logResult(result, w)

    fmt.Print("Validation Handler terminated\n")
}

func jsonOutput (result interface{}, w http.ResponseWriter) {
    jData, err := json.Marshal(result)
    if err != nil {
    http.Error(w, "Serialization error: " + err.Error(), http.StatusInternalServerError)
    return
    }

    // Write result to return html page
    w.Header().Set("Content-Type", "application/json")
    w.Write(jData)
}

func logResult(result interface{}, w http.ResponseWriter) {
    jData, err := json.Marshal(result)
    if err != nil {
        http.Error(w, "Serialization error: " + err.Error(), http.StatusInternalServerError)
        return
    }
    text := string(jData)

    dat, _ := ioutil.ReadFile("log.txt")

    stuff := []byte(string(dat) + text + "\n" )

    ioutil.WriteFile("log.txt", stuff, 0644)
}

func handleStatistics(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Server", "Go Link Hopper")
    w.Header().Set("Route", "statistics")
    w.WriteHeader(200)


    var results = storage.Read()
    if results == nil {
        fmt.Print("results is nil")
        return
    }
    jsonOutput(results, w)
}

func handleProxies(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Server", "Go Link Hopper")
    w.Header().Set("Route", "proxies")
    w.WriteHeader(200)
}

func Init() {
    fmt.Print("Init server\n")

    // Add route handlers
    http.HandleFunc("/validate", handleValidation)
    http.HandleFunc("/statistics", handleStatistics)
    http.HandleFunc("/proxies", handleProxies)

    // Start web server
    var err = http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
