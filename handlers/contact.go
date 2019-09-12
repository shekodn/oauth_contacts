package handlers

import(
"html/template"
"net/http"

"github.com/julienschmidt/httprouter"

)

type Contact struct {
    Name string
}

type ContactData struct {
    PageTitle string
    Contacts []Contact
}

var (
  data = ContactData{
    PageTitle: "Contacts",
    Contacts: []Contact{
        {Name: "Aníbal Troilo"},
        {Name: "Juan d'Arienzo"},
        {Name: "Julio Sosa"},
    },
  }
)

func getContacts(w http.ResponseWriter, r *http.Request,  _ httprouter.Params) {
  log.Info("Processing URL ", r.URL.Path)

  tmpl := template.Must(template.ParseFiles("contact/contact.html"))
  tmpl.Execute(w, data)
}
