package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	DBHost     = "127.0.0.1"
	DBPort     = "3306"
	DBUser     = "root"
	DBPassword = "rootpassword"
	DBName     = "go_blog_post_database"
	Port       = ":3000"
)

var database *sql.DB

type Page struct {
	Title   string
	Content string
	Date    string
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageId := vars["id"]
	page := Page{}

	fmt.Println(pageId)

	err := database.QueryRow("SELECT page_title,page_content,page_date FROM pages WHERE id=?", pageId).Scan(&page.Title, &page.Content, &page.Date)

	if err != nil {
		log.Println("Couldn't get page: ", pageId)
		log.Println(err.Error())
	}

	html := `<html><head><title>` + page.Title +
		`</title></head><body><h1>` + page.Title + `</h1><div>` +
		page.Content + `</div></body></html>`

	fmt.Fprintln(w, html)
}

func main() {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUser, DBPassword, DBHost, DBPort, DBName)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("Couldn't connect to" + DBName)
		log.Println(err.Error())
	}

	database = db

	routes := mux.NewRouter()
	routes.HandleFunc("/page/{id:[0-9]+}", ServePage)
	http.Handle("/", routes)

	log.Println("Server is running on port", Port)

	err = http.ListenAndServe(Port, nil)

	if err != nil {
		log.Panic("Couldn't start server", err.Error())
	}
}
