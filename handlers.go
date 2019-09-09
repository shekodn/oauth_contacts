package main

import (

  "fmt"
  "html/template"
  "net/http"

  "github.com/julienschmidt/httprouter"
  "github.com/shekodn/oauth_contacts/version"

  "golang.org/x/net/context"
  "golang.org/x/oauth2"
  "google.golang.org/api/people/v1"
)

type Contact struct {
    Name string
}

type ContactData struct {
    PageTitle string
    Contacts []Contact
}

var (
	// TODO: randomize it
	oauthStateString = "pseudo-random"

  data = ContactData{
    PageTitle: "Contacts",
    Contacts: []Contact{
        {Name: "Aníbal Troilo"},
        {Name: "Juan d'Arienzo"},
        {Name: "Julio Sosa"},
    },
  }
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  log.Info("Processing URL ", r.URL.Path)
  http.Redirect(w, r, "/contacts", http.StatusMovedPermanently)
}

func getContacts(w http.ResponseWriter, r *http.Request,  _ httprouter.Params) {
  log.Info("Processing URL ", r.URL.Path)

  tmpl := template.Must(template.ParseFiles("contact/contact.html"))
  tmpl.Execute(w, data)
}

// OAuth 
func importContacts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  log.Info("Processing URL ", r.URL.Path)

  url := googleOauthConfig.AuthCodeURL(oauthStateString)
  log.Info("Redirecting to URL ", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
  log.Info("Client is redirecting user agent to authorization endpoint.")
}

func callback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  log.Info("Processing URL ", r.URL.Path)

  err := getUserInfo(r.FormValue("state"), r.FormValue("code"))

	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

  http.Redirect(w, r, "/contacts", http.StatusMovedPermanently)
}

func getUserInfo(state string, code string) (error) {

  log.Info("Getting user info...")

  if state != oauthStateString {
		return fmt.Errorf("invalid oauth state")
	}

  log.Info("Client is sending authorization code and its own credentials to token endpoint")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)

  if err != nil {
    log.Fatalf("Error: %v", err)
    return fmt.Errorf("Error: %v", err)
  }

  client := googleOauthConfig.Client(context.Background(), token)

  srv, err := people.New(client)
  if err != nil {
      log.Fatalf("Unable to create people Client %v", err)
      return fmt.Errorf("Error: %v", err)
  }

  log.Info("Client is sending access token to protected resource")
  r, err := srv.People.Connections.List("people/me").PageSize(5).
        PersonFields("names,emailAddresses").Do()
  if err != nil {
      log.Fatalf("Unable to retrieve people. %v", err)
      return fmt.Errorf("Error: %v", err)
  }

  log.Info("Client is receiving the protected resource")
  if len(r.Connections) > 0 {
      fmt.Print("Listing connection names:\n", )
      for _, c := range r.Connections {
          names := c.Names
          if len(names) > 0 {
              name := names[0].DisplayName
              data.Contacts = append(data.Contacts, Contact{ Name: name })
          }
      }
  } else {
      fmt.Print("No connections found.")
  }

  return nil
}

func getVersion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  log.Info("Request received: getVersion")
  fmt.Fprintf(w, "Repo: %s, Commit: %s, Version: %s", version.REPO, version.COMMIT, version.RELEASE)
}
