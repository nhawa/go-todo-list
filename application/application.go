package application

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type (
	ApplicationInterface interface {
		ListenAndServe()
	}
	Application struct {
		mux          *mux.Router
		ApplicationDB *sql.DB
		Version       string
	}
)

func New() ApplicationInterface {
	return &Application{
		mux: mux.NewRouter(),
	}
}

func (application *Application) loadApplicationDatabase() {

	connStr := "user=otten host=localhost password=password dbname=todo_lists"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	log.Info("Connected to Database")
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Minute * 10)
	application.ApplicationDB = db

}

// ListenAndServe app
func (application *Application) ListenAndServe() {
	r := application.mux


	application.loadApplicationDatabase()

	defer application.ApplicationDB.Close()

	application.register()

	http.Handle("/", r)

	var address = "localhost:8080"
	fmt.Printf("server started at %s\n", address)

	server := new(http.Server)
	server.Addr = address
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
