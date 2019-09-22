package main

import(
  "net/http"
  "os"

  "github.com/sirupsen/logrus"
  "github.com/shekodn/oauth_contacts/handlers"
  "github.com/shekodn/oauth_contacts/version"
)

var (
  log = logrus.New()
)

// Prepare vendoring config: glide init; glide install

// Run server: go build -o app && ./app
func main() {
  log.Printf("Starting the service...\n")
  r := handlers.Router(version.BuildTime, version.Commit, version.Release)
  log.Print("The service is ready to listen and serve.")
  log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
