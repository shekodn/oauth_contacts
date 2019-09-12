package handlers

import (
  "github.com/julienschmidt/httprouter"
  "github.com/sirupsen/logrus"
)

var (
  log = logrus.New()
)

// Router register necessary routes and returns an instance of a router.
func Router(buildTime, commit, release string) *httprouter.Router {

	r := httprouter.New()

  // home
  r.GET("/", home)

  // contacts
  r.GET("/contacts", getContacts)

  // auth
  r.GET("/importContacts", importContacts)
  r.GET("/callback", callback)

  return r
}
