package main

import(
  "net/http"

  "github.com/julienschmidt/httprouter"
  "github.com/sirupsen/logrus"
)

var log = logrus.New()

// Prepare vendoring config: glide init; glide install

// Run server: go build -o app && ./app
// Try requests: curl http://127.0.0.1:8000/any
func main() {
  log.Info("Initialize service...")

  router := httprouter.New()
	router.GET("/", home)
  router.GET("/version", getVersion)

  log.Info("Service is ready to listen and serve.")
  http.ListenAndServe(":8000", router)
}
