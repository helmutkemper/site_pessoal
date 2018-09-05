package main

import (
  "log"
  "net/http"
)

func main() {
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

  log.Println("Listening... :3000")
  http.ListenAndServe(":3000", nil)
}
