package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
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
	var u User
	err := app.Database.QueryRowx("SELECT firstname, created FROM account WHERE id = $1", id).
		StructScan(&u)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		// log.Fatal("Database SELECT failed")
	}

	if api == true {
		js, err := json.Marshal(u)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
		return
	}

	// not api? return in a template
	app.Tmpl.ExecuteTemplate(w, "user.tmpl", u)
}

func (app *App) newUser(w http.ResponseWriter, r *http.Request) {

	firstname := "Raphael"
	lastname := "Enkoder"
	email := "raphael@enkoder.com.au"
	created := time.Now()

	_, err := app.Database.Exec("INSERT INTO account (firstname,lastname,email, created) VALUES ($1,$2, $3, $4)", firstname, lastname, email, created)

	if err != nil {
		fmt.Printf("error: %v\n", err)
		// log.Fatal("Database insert failed")
	}

	w.WriteHeader(http.StatusOK)
}
