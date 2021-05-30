package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/kaatinga/plantbook/cmd/plantbook-front/controllers"
	"github.com/kaatinga/plantbook/cmd/plantbook-front/views"
)

var (
	homeView     *views.View
	notfoundView *views.View
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	must(notfoundView.Render(w, nil))
}

func main() {
	homeView = views.NewView("index", "views/home.gohtml")
	notfoundView = views.NewView("index", "views/404.gohtml")
	usersC := controllers.NewUsers()
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home).Methods("GET")

	r.HandleFunc("/api/v1/user/login", usersC.GetUserByName).Methods("GET")
	r.HandleFunc("/api/v1/user/login", usersC.CreateUser).Methods("POST")
	staticfile := http.FileServer(http.Dir("assets"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", staticfile))
	fmt.Println("Frontend compiled successfully")
	http.ListenAndServe(":8080", r)
}
