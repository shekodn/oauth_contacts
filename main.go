package main

import(
  "net/http"
  "os"

  "github.com/julienschmidt/httprouter"
  "github.com/sirupsen/logrus"
  "github.com/joho/godotenv"

  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
)

var (
  googleOauthConfig *oauth2.Config
  log = logrus.New()
)

// Prepare vendoring config: glide init; glide install

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

// Run server: go build -o app && ./app
// Try requests: curl http://127.0.0.1:8000/any
func main() {
  log.Info("Initialize service...")

  router := httprouter.New()
	router.GET("/", home)
  router.GET("/version", getVersion)

  router.GET("/contacts", getContacts)
  router.GET("/importContacts", importContacts)
  router.GET("/callback", callback)

  log.Info("Service is ready to listen and serve.")
  http.ListenAndServe(":" + os.Getenv("PORT"), router)
}
