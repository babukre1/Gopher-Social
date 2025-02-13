package main

import (
	"encoding/json"
	"net/http"

)



// "ecommerce-api/cmd/api"
// "log"

// func main() {
// 	server := api.NewApiServer(":8081", nil)

// 	if err := server.Run(); err != nil {
// 		log.Fatal(err)
// 	}
// }

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	w.Write([]byte("Hello From the Server"))
	// switch r.Method {
	// case http.MethodGet:
	// 	switch r.URL.Path {
	// 	case "/":
	// 		w.Write([]byte("Index Page"))
	// 		return
	// 	case r.URL.Path:
	// 		w.Write([]byte(r.URL.Path))
	// 		return
	// 	}
	// default:
	// 	w.Write([]byte("404 not found"))
	// 	return
	// }
}

func (s *api) CreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Created Succesfully"))
}

func (s *api) DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Deleted Succesfully"))
}

type User struct {
	Name   string `json"name"`
	Age    int `json"age"`
	Gender string `json"gender"`
}

var users = []User{}

func (s *api) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)

}
func (s *api) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	u := User{
		Name:   payload.Name,
		Age:    payload.Age,
		Gender: payload.Gender,
	}
	users = append(users, u)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User Created Succesfully"))

}

// func insertUser(u User) error {
// 	if u.Name == "" {
// 		return errors.New("Name is Required")
// 	}
// 	if u.Age <= 0 {
// 		return errors.New("Not a valid Age. ")
// 	}

// 	users = append(users, u)
// 	return nil
// }

func main() {
	api := &api{addr: ":8081"}
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("POST /createUser", api.CreateUsersHandler)
	mux.HandleFunc("POST /deleteUser", api.DeleteUsersHandler)
	mux.HandleFunc("GET /getUsers", api.GetUserHandler)
	mux.HandleFunc("POST /createUser", api.CreateUserHandler)
	srv.ListenAndServe()
}
