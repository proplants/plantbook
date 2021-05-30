package controllers

import (
	"fmt"
	"net/http"
	"github.com/kaatinga/plantbook/cmd/plantbook-front/views"
)


func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("index", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}


func (u *Users) GetUserByName(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}


type UserLoginPassword struct {
	Login string `json:"login"`
	Password string `json:"password"`
}


func (u *Users) CreateUser(w http.ResponseWriter, r *http.Request) {
	var form UserLoginPassword
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
