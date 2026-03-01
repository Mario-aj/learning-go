package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

const (
	Port = ":3000"
)

func serveDynamic(w http.ResponseWriter, _ *http.Request) {
	response := "The time now is " + time.Now().String()

	fmt.Fprintln(w, response)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pageId := vars["id"]
	filename := "files/" + pageId + ".html"

	_, err := os.Stat(filename)

	if err != nil {
		filename = "files/404.html"
	}

	http.ServeFile(w, r, filename)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", serveDynamic)
	router.HandleFunc("/static", serveStatic)
	router.HandleFunc("/pages/{id}", pageHandler)

	http.Handle("/", router)
	http.ListenAndServe(Port, nil)
}
