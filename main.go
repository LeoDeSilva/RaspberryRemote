package main

import (
  "fmt"
  "net/http"
  "log"
  "github.com/pkg/browser"
)

const PORT = ":10000"

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Hello World")
}

func linkPage(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case "GET":
    fmt.Println("Expected POST Request at /link")
  case "POST":
    if err := r.ParseForm(); err != nil {
      fmt.Printf("ParseForm() err: %v", err);
      return
    }

    url := r.FormValue("url")
    fmt.Println("URL",url)
    browser.OpenURL(url)
  }
}   

func handleRequests() {
  http.HandleFunc("/", homePage) 
  http.HandleFunc("/link", linkPage) 
  log.Fatal(http.ListenAndServe(PORT, nil))
}

func main() {
  fmt.Printf("RemotePI - Server open on port: '%s'\n", PORT)
  handleRequests()
}
