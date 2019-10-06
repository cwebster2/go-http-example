package galactica

import (
  "log"
  "bytes"
  "fmt"
  "net/http"
)

func init() {
  log.Printf("Initializing Galactica\n")
}

type GalacticaOptions struct {
  Port int
}

type RequestHandlerFunc func(w http.ResponseWriter, req *http.Request)

type GalacticaApp struct{
  port string
  middleware map[string][]RequestHandlerFunc
}

func Galactica(options GalacticaOptions) *GalacticaApp{
  var app GalacticaApp
  //app.port, _ = fmt.Scanf(":%d", options.port)
  app.port = ":8080"
  app.middleware = make(map[string][]RequestHandlerFunc)
  return &app
}

func (g *GalacticaApp) Run() {
  log.Printf("Galactica App Listening on port %s\n", g.port)
  log.Fatal(http.ListenAndServe(g.port, nil))
}

func (g *GalacticaApp) Use(m RequestHandlerFunc) {
  g.middleware["*"] = append(g.middleware["*"], m)
}

func (g *GalacticaApp) String() string {
  var b bytes.Buffer
  b.WriteString("{\n")
  fmt.Fprintf(&b, "%s\n", g.port)
  for route, funcs := range g.middleware {
    fmt.Fprintf(&b, "%s: %v", route, funcs)
  }
  b.WriteString("}\n")
  return b.String()
}
