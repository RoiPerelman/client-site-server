package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/roiperelman/client-site-server/models"
	"github.com/roiperelman/client-site-server/handlers"
)


func main() {
	models.InitDB()
	r := mux.NewRouter()

	r.HandleFunc("/api/user/authorize", handlers.AuthorizeUser).Methods("GET")
	r.HandleFunc("/api/user/signup", handlers.SignupUser).Methods("POST")
	r.HandleFunc("/api/user/login", handlers.LoginUser).Methods("POST")
	r.PathPrefix("/static").Handler(http.FileServer(http.Dir("./build")))
	r.PathPrefix("/").HandlerFunc(staticFileHandler("./build/index.html"))

	http.ListenAndServe(":1111", r)
}

func staticFileHandler(filePath string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	}

	return http.HandlerFunc(fn)
}