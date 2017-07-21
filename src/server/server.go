package main

import (
    "fmt"
    "net/http"
    "html/template"
    //"net/url"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // trying to read parameter
    r.ParseForm()
    url := r.Form.Get("url")

    if url == "" {
       //url := "no url found"
       fmt.Printf("url not found")
    }
    
    t, err := template.ParseFiles("handler.html")
    if err != nil{
      panic(err)
    }
    t.Execute(w, "static")
}

func main() {
    fmt.Printf("Hello, friend")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
