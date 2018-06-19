package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/roiperelman/client-site-server/models"
	"github.com/roiperelman/client-site-server/handlers"
	"flag"
	"fmt"
	"os"
)


func main() {

	var port, staticLocation string
	flag.StringVar(&port ,"port", "1111", "The Port the App Will Listen To")
	flag.StringVar(&staticLocation,"static", "./build", "The Location the App statically serves from")
	flag.Parse()

	fmt.Println("the staticLocation is " + staticLocation)
	fmt.Println("The port is " + getEnv("PORT", port))

	models.InitDB()
	r := mux.NewRouter()

	r.HandleFunc("/api/user/authorize", handlers.AuthorizeUser).Methods("GET")
	r.HandleFunc("/api/user/signup", handlers.SignupUser).Methods("POST")
	r.HandleFunc("/api/user/login", handlers.LoginUser).Methods("POST")
	r.PathPrefix("/static").Handler(http.FileServer(http.Dir(staticLocation)))
	r.PathPrefix("/").HandlerFunc(staticFileHandler(staticLocation + "/index.html"))

	http.ListenAndServe(":" + getEnv("PORT", port), r)
}

func staticFileHandler(filePath string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	}
	return http.HandlerFunc(fn)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}