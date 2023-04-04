package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/smebellis/lenslocked/views"
)

// Helper function to handle repeated code in handlers
func executeTemplate(w http.ResponseWriter, filepath string) {

	//Parse the template
	tpl, err := views.Parse(filepath)
	if err != nil {
		log.Printf("processing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, nil)

}

// This is the function that is called anytime someone comes to the web server
func homeHandler(w http.ResponseWriter, r *http.Request) {

	// Creates a path that is OS agnostic
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)

}

// Creating a contact handler function to serve a contact page
func contactHandler(w http.ResponseWriter, r *http.Request) {

	// Creates a path that is OS agnostic
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

// Exercise to implement an FAQ page
func faqHandler(w http.ResponseWriter, r *http.Request) {
	// Creates a path that is OS agnostic
	tplPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tplPath)
}

func main() {

	r := chi.NewRouter()

	//Using Chi Middleware logger
	r.Use(middleware.Logger)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting server on port 3000...")
	http.ListenAndServe(":3000", r)

}
