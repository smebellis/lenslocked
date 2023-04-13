package controllers

import (
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// TODO: Render the signup page
	u.Templates.New.Execute(w, nil)
}
