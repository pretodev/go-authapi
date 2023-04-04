package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"main/src/controllers"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	router := mux.NewRouter()
	router.HandleFunc("/users/infos", controllers.GetInfos).Methods("GET")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "9099"
	}
	fmt.Printf("Servidor rodando na porta %s...\n", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
