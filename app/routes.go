package app

import (
	"fmt"
	"net/http"
)

func (app *App) InitializeRoutes() {

	fs := http.FileServer(http.Dir("./static/"))
	app.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// set true for api
	app.Router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		app.getUser(w, r, true)
	}).Methods("GET")

	app.Router.HandleFunc("/newuser", app.newUser).Methods("POST")

	fmt.Printf("app: %s\n", app.Config.Appname)

}
