package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/roiperelman/client-site-server/models"
	"github.com/roiperelman/client-site-server/handlers"
	"github.com/roiperelman/client-site-server/utils"
	"github.com/joho/godotenv"
	"flag"
	"log"
	"github.com/roiperelman/client-site-server/middlewares"
)

func main() {

	var port, staticLocation string
	flag.StringVar(&port ,"port", "1111", "The Port the App Will Listen To")
	flag.StringVar(&staticLocation,"static", "./client/build", "The Location the App statically serves from")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}

	models.InitDB()
	r := mux.NewRouter()

	log.Println("The port is " + utils.GetEnv("PORT", port))
	log.Println("the staticLocation is " + staticLocation)

	//r.HandleFunc("/api/user/authorize", handlers.AuthorizeUser).Methods("GET")
	r.HandleFunc("/api/user/signup", handlers.SignupUser).Methods("POST")
	r.HandleFunc("/api/user/login", handlers.LoginUser).Methods("POST")
	r.Handle("/api/user/authorize", middlewares.Authenticate(http.HandlerFunc(handlers.AuthorizeUser))).Methods("GET")
	r.Handle("/api/user/multipleSections", middlewares.Authenticate(http.HandlerFunc(handlers.MultipleSectionsUser))).Methods("POST")
	r.Handle("/api/user/addSection", middlewares.Authenticate(http.HandlerFunc(handlers.AddSectionUser))).Methods("POST")
	r.Handle("/api/user/delSection", middlewares.Authenticate(http.HandlerFunc(handlers.DelSectionUser))).Methods("POST")
	r.Handle("/api/user/addContextItem", middlewares.Authenticate(http.HandlerFunc(handlers.AddContextItem))).Methods("POST")
	r.Handle("/api/user/delContextItem", middlewares.Authenticate(http.HandlerFunc(handlers.DelContextItem))).Methods("POST")
	r.PathPrefix("/static").Handler(http.FileServer(http.Dir(staticLocation)))
	r.PathPrefix("/").HandlerFunc(staticFileHandler(staticLocation + "/index.html"))

	http.ListenAndServe(":" + utils.GetEnv("PORT", port), r)
}

func staticFileHandler(filePath string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	}
	return http.HandlerFunc(fn)
}

