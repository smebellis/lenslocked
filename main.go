package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/smebellis/lenslocked/controllers"
	"github.com/smebellis/lenslocked/templates"
	"github.com/smebellis/lenslocked/views"
)

func main() {

	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))

	r.Get("/faq", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting server on port 3000...")
	http.ListenAndServe(":3000", r)

}
