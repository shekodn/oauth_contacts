package main

import (
  "fmt"
  "net/http"

  "github.com/julienschmidt/httprouter"
  "github.com/shekodn/oauth_contacts/version"

)

// home path
func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  log.Info("Request received.")
  fmt.Fprintf(w, "Processing URL %s...\n", r.URL.Path)
}

func getVersion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  log.Info("Request received.")
  fmt.Fprintf(w, "Repo: %s, Commit: %s, Version: %s", version.REPO, version.COMMIT, version.RELEASE)
}
