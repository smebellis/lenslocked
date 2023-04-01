package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Helper function to handle repeated code in handlers
func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//Parse the template
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("processing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
	}
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
		<li>
			<b>Is there a free version?</b>
			Yes! We offer a free trial for 30 days on any paid plans.
		</li>
		<li>
			<b>What are your support hours?</b>
			We have support staff answering emails 24/7, though response
			times may be a bit slower on weekends.
		</li>
		<li>
			<b>How do I contact support?</b>
			Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
		</li>
	</ul>
	`)
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
