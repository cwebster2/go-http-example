package main

import (
  "fmt"
  "./galactica"
  "net/http"
)

func main() {
  fmt.Println("Test App Started")

  var opts galactica.GalacticaOptions
  opts.Port = 8080
  // create galactica app
  app := galactica.Galactica(opts)
  fmt.Printf("%T",app)

  // add a middleawre
  app.Use(func(w http.ResponseWriter, req *http.Request){
    fmt.Println("Recieved Request")
  })
  // add another middleware
  // add a route handler

  fmt.Println(app)
  // run app
  app.Run()
}
