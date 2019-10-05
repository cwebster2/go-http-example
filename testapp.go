package main

import (
  "fmt"
  "./galactica"
)

func main() {
  fmt.Println("Test App Started")

  // create galactica app
  app := galactica.Galactica()
  fmt.Printf("%T",app)

  // add a middleawre
  // add another middleware
  // add a route handler

  // run app
  app.Run()
}
