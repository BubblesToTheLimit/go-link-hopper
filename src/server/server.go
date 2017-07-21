package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "validator"
//    "os"
    "io/ioutil"
//    "write"
)

type Result struct {
    Id int
    Url string
    Target string
//    Trace []string
//    CreatedAt time.Time
    Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
    // trying to read parameter
    r.ParseForm()
    url := r.Form.Get("url")
    message := ""

    if url == "" {
       fmt.Printf("url not found")
       message = "please enter a valid url"
    } else {
       message = "url ok"
    }

    var target = validator.Validator(url)
    
    myresult := Result{
                 Id: 1,
                 Url: url,
                 Target: target,
//                 Trace: trace,
//                 CreatedAt: time,
                 Message: message,
    }

    jData, err := json.Marshal(myresult)
    if err != nil {
       fmt.Println("json Marshal parsing error", err)
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jData)

    logresult(string(jData))
}

func logresult(text string) {
    dat, _ := ioutil.ReadFile("log.txt")

    stuff := []byte(string(dat) + "\n" + text )

    ioutil.WriteFile("log.txt", stuff, 0644)

}

func main() {
    fmt.Printf("Hello, friend")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
