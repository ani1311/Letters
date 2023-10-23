package main

import (
	"html/template"
	"log"
	"net/http"
)

type Contact struct {
	Name  string
	Email string
}

type Letter struct {
	Title     string
	ContactId int
	Content   string
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")
	email := r.FormValue("Email")

	tmpl := template.Must(template.ParseFiles("templates/send.gohtml"))

	err := tmpl.Execute(w, Contact{
		Name:  name,
		Email: email,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func contactsHandler(w http.ResponseWriter, r *http.Request) {
	contacts := []Contact{
		{Name: "John Doe", Email: "john.doe@example.com"},
		{Name: "Jane Smith", Email: "jane.smith@example.com"},
		{Name: "Bob Johnson", Email: "bob.johnson@example.com"},
	}

	tmpl := template.Must(template.ParseFiles("templates/contacts.gohtml"))

	err := tmpl.Execute(w, contacts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func lettersHandler(w http.ResponseWriter, r *http.Request) {
	letters := []Letter{
		{
			Title:     "Letter 1",
			ContactId: 1,
			Content:   "Hello, World!",
		},
	}
	tmpl := template.Must(template.ParseFiles("templates/letters.gohtml"))

	err := tmpl.Execute(w, struct {
		Name    string
		Letters []Letter
	}{
		Name:    "John Doe",
		Letters: letters,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/send", sendHandler)
	r.HandleFunc("/contacts", contactsHandler)
	r.HandleFunc("/letters", lettersHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
