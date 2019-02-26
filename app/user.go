package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	// "log"
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
		fmt.Print("error: no id in request")
		return
	}
	var u User
	err := app.Database.QueryRowx("SELECT firstname, email, created FROM users WHERE id = $1", id).
		StructScan(&u)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
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

	// vars := mux.Vars(r)

	// firstname := vars["firstname"]
	firstname := "Dev"
	lastname := "Enkoder"
	email := "raphael@enkoder.com.au"
	created := time.Now()

	_, err := app.Database.Exec("INSERT INTO users (firstname,lastname,email, created) VALUES ($1,$2, $3, $4)", firstname, lastname, email, created)

	if err != nil {
		fmt.Printf("error: %v\n", err)
		// log.Fatal("Database insert failed")
	}

	w.WriteHeader(http.StatusOK)
}
