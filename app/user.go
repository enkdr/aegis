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
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Department string    `json:"department"`
	Created    time.Time `json:"created"`
}

func (app *App) getUser(w http.ResponseWriter, r *http.Request, api bool) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Fatal("No id requested")
	}

	user := &User{}

	err := app.Database.QueryRow("SELECT id, created, username, department FROM  userinfo WHERE id = $1", id).
		Scan(&user.ID, &user.Created, &user.Username, &user.Department)
	if err != nil {
		log.Fatal("Database SELECT failed")
	}

	if api == true {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			panic(err)
		}
		return
	}
	app.Tmpl.ExecuteTemplate(w, "user.tmpl", user)
}

func (app *App) newUser(w http.ResponseWriter, r *http.Request) {

	user := "Dev"
	dep := "Sysadmin"
	_, err := app.Database.Exec("INSERT INTO userinfo (username,department) VALUES ($1,$2)", user, dep)

	if err != nil {
		fmt.Println(err)
		log.Fatal("Database insert failed")
	}

	w.WriteHeader(http.StatusOK)
}
