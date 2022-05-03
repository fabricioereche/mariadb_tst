package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/fabricioereche/mariadb_tst/cmd/web/handlers"
	"github.com/fabricioereche/mariadb_tst/core/beer"
	"github.com/fabricioereche/mariadb_tst/infraestructure/persistence"
	"github.com/fabricioereche/mariadb_tst/infraestructure/persistence/repository"
	"github.com/gorilla/mux"
)

func main() {
	db, err := persistence.GetDrive()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repository := repository.NewRepository(db)
	service := beer.NewService(repository)
	r := mux.NewRouter()
	n := negroni.New(negroni.NewLogger())

	handlers.MakeBeerHandlers(r, n, service)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// used to health check, will return 200
	})

	http.Handle("/", r)

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":" + os.Getenv("HTTP_PORT"),
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
