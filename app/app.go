package app

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"html/template"
)

type Conf struct {
	Appname string
	Company string
	Port    int
}

type App struct {
	Router   *mux.Router
	Database *sqlx.DB
	Tmpl     *template.Template
	Config   Conf
}
