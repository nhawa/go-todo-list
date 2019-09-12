package application

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
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

	// Healtcheck
	healthCheck := func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()

		err := application.ApplicationDB.Ping()

		if err != nil {
			logrus.Error("Health check: Can't connect to application Database")
			return
		}

		err = application.ApplicationDB.Ping()
		if err != nil {
			logrus.Error("Health check: Can't connect to Otten Database")
			return
		}

		fmt.Printf("Ok|%s", hostname)
	}

	r.HandleFunc("/", healthCheck).Methods("GET", "HEAD")
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
