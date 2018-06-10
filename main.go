package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
)

type UserErrors struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Password    string `json:"password"`
}

type User struct {
	Email       string     `json:"email"`
	Username        string     `json:"username"`
	Password    string     `json:"password"`
	Jwt         string     `json:"jwt"`
	Errors      UserErrors `json:"errors"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/user/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(formatRequest(r))
		r.ParseForm()
		fmt.Println(r.PostForm.Get("email"))
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		} else {
			fmt.Fprintf(w, "email: %s, name: %s, PhoneNumber: %s, password: %s", user.Email, user.Username, user.Username, user.Password)
		}

	}).Methods("POST")
	http.ListenAndServe(":1111", r)
}

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
	r.ParseForm()
	request = append(request, "\n")
	request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}