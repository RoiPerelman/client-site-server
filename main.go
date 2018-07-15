package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/roiperelman/client-site-server/handlers"
	"github.com/roiperelman/client-site-server/middlewares"
	"github.com/roiperelman/client-site-server/models"
	"github.com/roiperelman/client-site-server/utils"
)

func main() {
	var port, staticLocation string
	flag.StringVar(&port, "port", "1111", "The Port the App Will Listen To")
	flag.StringVar(&staticLocation, "static", "./client/build", "The Location the App statically serves from")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}

	log.Println("The port is " + utils.GetEnv("PORT", port))
	log.Println("the staticLocation is " + staticLocation)

	db, err := models.InitDB()
	if err != nil {
		log.Panicf("Error Connecting to DB: %v", err)
	}

	dbStore := &models.DBStore{db}

	r := mux.NewRouter()

	r.Use(dbStore.DBStoreMiddleware)

	//r.HandleFunc("/api/user/authorize", handlers.AuthorizeUser).Methods("GET")
	r.HandleFunc("/api/user/signup", handlers.SignupUser).Methods("POST")
	r.HandleFunc("/api/user/login", handlers.LoginUser).Methods("POST")
	r.Handle("/api/user/authorize", middlewares.AuthenticateDBUserMiddleware(http.HandlerFunc(handlers.AuthorizeUser))).Methods("GET")
	r.Handle("/api/user/multipleSections", middlewares.AuthenticateClaimsMiddleware(http.HandlerFunc(handlers.MultipleSectionsUser))).Methods("POST")
	r.Handle("/api/user/addSection", middlewares.AuthenticateClaimsMiddleware(http.HandlerFunc(handlers.AddSectionUser))).Methods("POST")
	r.Handle("/api/user/delSection", middlewares.AuthenticateDBUserMiddleware(http.HandlerFunc(handlers.DelSectionUser))).Methods("POST")
	r.Handle("/api/user/addContextItem", middlewares.AuthenticateClaimsMiddleware(http.HandlerFunc(handlers.AddContextItem))).Methods("POST")
	r.Handle("/api/user/delContextItem", middlewares.AuthenticateClaimsMiddleware(http.HandlerFunc(handlers.DelContextItem))).Methods("POST")
	r.Handle("/api/user/updateJSCode", middlewares.AuthenticateClaimsMiddleware(http.HandlerFunc(handlers.UpdateJSCode))).Methods("POST")
	r.PathPrefix("/static").Handler(http.FileServer(http.Dir(staticLocation)))
	r.PathPrefix("/").HandlerFunc(staticFileHandler(staticLocation + "/index.html"))

	http.ListenAndServe(":"+utils.GetEnv("PORT", port), r)
}

func staticFileHandler(filePath string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	}
	return http.HandlerFunc(fn)
}
