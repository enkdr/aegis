package app

import (
	"database/sql"
	"github.com/gorilla/mux"
	"html/template"
)

type Conf struct {
	Appname string
}

type App struct {
	Router   *mux.Router
	Database *sql.DB
	Tmpl     *template.Template
	Config   Conf
}
