package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/smebellis/lenslocked/controllers"
	"github.com/smebellis/lenslocked/views"
)

func main() {

	r := chi.NewRouter()

	tpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting server on port 3000...")
	http.ListenAndServe(":3000", r)

}
