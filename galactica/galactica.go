package galactica

import (
  "log"
  "net/http"
)

func init() {
  log.Printf("Initializing Galactica\n")
}

type RequestHandlerFunc func(w http.ResponseWriter, req *http.Request)
type GalacticaApp struct{
  port int
  middleware []RequestHandlerFunc
  endpoint string
}

func Galactica() *GalacticaApp{
  var app GalacticaApp
  app.port = 8080
  return &app
}

func (g *GalacticaApp) Run() {
  g.endpoint = ":8080"
  log.Printf("Galactica App Listening on port %s\n", g.endpoint)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func (g *GalacticaApp) Use(m RequestHandlerFunc) {
  g.middleware = append(g.middleware, m)
}
