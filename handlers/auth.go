package handlers

import(
  "fmt"
  "net/http"
  "os"

  "github.com/joho/godotenv"
  "github.com/julienschmidt/httprouter"

  "golang.org/x/net/context"
  "golang.org/x/oauth2"
  "google.golang.org/api/people/v1"
  "golang.org/x/oauth2/google"
)

var (
  googleOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

func init() {

  err := godotenv.Load()

  if err != nil {
    log.Fatal("Error loading .env file")
  }

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:" + os.Getenv("PORT") + "/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
    Scopes:       []string{"https://www.googleapis.com/auth/contacts.readonly"},
		Endpoint:     google.Endpoint,
	}
}

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
