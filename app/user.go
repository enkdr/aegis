package app

import (
	"time"

	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	ID        int       `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Created   time.Time `json:"created"`
}

func (app *App) getUser(w http.ResponseWriter, r *http.Request, api bool) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Fatal("No id requested")
	}

	// user := &User{}

	var u User
	err := app.Database.QueryRowx("SELECT firstname FROM account WHERE id = $1", id).
		StructScan(&u)
	if err != nil {
		log.Fatal("Database SELECT failed")
	}

	if api == true {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(u); err != nil {
			panic(err)
		}
		return
	}
	app.Tmpl.ExecuteTemplate(w, "user.tmpl", u)
}

func (app *App) newUser(w http.ResponseWriter, r *http.Request) {

	firstname := "Dev"
	lastname := "Enkoder"
	email := "dev@enkoder.com.au"
	_, err := app.Database.Exec("INSERT INTO account (firstname,lastname,email) VALUES ($1,$2, $3)", firstname, lastname, email)

	if err != nil {
		fmt.Println(err)
		log.Fatal("Database insert failed")
	}

	w.WriteHeader(http.StatusOK)
}
