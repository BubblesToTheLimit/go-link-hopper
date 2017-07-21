package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type Result struct {
    Url string
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
       message = "url received"
    }

    myresult := Result{
                 Url: url,
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
    //fmt.Fprintf(w, string(jData))
}

func main() {
    fmt.Printf("Hello, friend")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
