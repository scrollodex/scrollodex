package main

import (
	"edit"
	"log"
	"logout"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"admin"
	"callback"
	"home"
	"login"
	"middlewares"
	"unauthorized"
)

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", home.HomeHandler)
	r.HandleFunc("/login", login.LoginHandler)
	r.HandleFunc("/logout", logout.LogoutHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.HandleFunc("/unauthorized", unauthorized.UnauthorizedHandler)
	//r.Handle("/user", negroni.New(
	//	negroni.HandlerFunc(middlewares.IsAuthenticated),
	//	negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	//))
	r.Handle("/admin", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.HandlerFunc(middlewares.IsRbacEditor),
		negroni.Wrap(http.HandlerFunc(admin.AdminHandler)),
	))
	r.Handle("/admin/edit/{site:[a-z]+}/{table:[a-z]+}", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.HandlerFunc(middlewares.IsRbacEditor),
		negroni.Wrap(http.HandlerFunc(edit.EditHandler)),
	)).Methods("GET", "PUT")
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/")))) // static files
	http.Handle("/", r)
	log.Print("Server listening on http://localhost:3000/")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}
