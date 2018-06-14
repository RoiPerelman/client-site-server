package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/roiperelman/client-site-server/models"
	"github.com/roiperelman/client-site-server/handlers"
)


func main() {
	models.InitDB()
	r := mux.NewRouter()

	r.HandleFunc("/api/user/authorize", handlers.AuthorizeUser).Methods("GET")
	r.HandleFunc("/api/user/signup", handlers.SignupUser).Methods("POST")
	r.HandleFunc("/api/user/login", handlers.LoginUser).Methods("POST")
	http.ListenAndServe(":1111", r)
}

//func JsonResponse(response interface{}, w http.ResponseWriter) {
//
//	json, err := json.Marshal(response)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(json)
//}

//err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(user.Password))
//if err != nil {
//	http.Error(w, err.Error(), 500)
//}

//timestamp := claims["timestamp"]
//if t, ok := timestamp.(float64); ok {
//	duration := int64(time.Now().Sub(time.Unix(int64(t), 0)))
//	maximumDuration := int64(3600 * time.Second)
//	fmt.Println(maximumDuration - int64(duration))
//} else {
//	fmt.Println("err")
//	fmt.Println(time.Now().Unix())
//}