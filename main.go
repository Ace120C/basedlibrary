package main

import (
  "net/http"
  "log"
  "webdev/basedlibrary/API"
)

func main()  {
  fs := http.FileServer(http.Dir("./static/"))
  http.Handle("/", fs)

  api.APIhost()

  log.Println("Listening on :8080...")
	log.Println("& Serving database at http://localhost:8080/dbOUT.json")
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal(err)
  }
}
