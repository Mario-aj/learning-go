package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func testeHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/teste", testeHandler)

	http.Handle("/", router)

	fmt.Println("Everything is going to be alright")
}
