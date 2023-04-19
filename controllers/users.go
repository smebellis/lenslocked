package controllers

import (
	"fmt"
	"net/http"

	"github.com/smebellis/lenslocked/models"
)

type Users struct {
	Templates struct {
		New Template
	}

	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// TODO: Render the signup page
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "something went wrong.", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User created: %+v", user)
}
