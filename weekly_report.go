package main

import (
  "fmt"
  "net/http"
)

func main() {
  resp, err := http.Get("http://hazlo.herokuapp.com/emails/weekly")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(resp.Status)
}
