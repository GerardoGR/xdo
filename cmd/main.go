package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		log.Print("Can not connect to postgres:", err)
		return
	}

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	if err != nil {
		log.Print("Can not ping postgres:", err)
		return
	}

	// TODO: Validate listenAddress: $host:$port
	listenAddress := flag.String("listen-address", "localhost:8000", "Domain name where shorty is being served")
	flag.Parse()

	http.HandleFunc("/", homePage)

	fmt.Printf("Starting server: http://%s\n", *listenAddress)
	err = http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		todosIndex(w, r)
	} else if r.Method == http.MethodPost {
		createTodo(w, r)
	} else {
		http.Error(w, "Only GET or POST are supported", http.StatusMethodNotAllowed)
		return
	}
}

func todosIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Print("Error parsing template:", err)
		http.Error(w, "Unexpected server error", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print("Error parsing form:", err)
		http.Error(w, "Bad request check the form data", http.StatusBadRequest)
		return
	}

	// TODO: Get postgresql configuration
	todoName := r.FormValue("todo")

	insertDynStmt := `insert into "todo" ("name") values($1)`
	_, err := db.Exec(insertDynStmt, todoName)
	if err != nil {
		log.Print("Can not insert into todo table:", err)
		http.Error(w, "Unexpected server error", http.StatusInternalServerError)
		return
	}

	log.Printf("New ToDo: %s", todoName)
	fmt.Fprintf(w, "New ToDo: %s", todoName)
}

const (
	host     = "localhost"
	port     = 5432
	user     = "xdo"
	password = "my-secret-password"
	dbname   = "xdo"
)
