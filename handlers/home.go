package handlers

import(
"net/http"

"github.com/julienschmidt/httprouter"

)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  log.Info("Processing URL ", r.URL.Path)
  http.Redirect(w, r, "/contacts", http.StatusMovedPermanently)
}
