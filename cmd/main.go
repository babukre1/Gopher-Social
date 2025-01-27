package main

import (
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
	// w.Write([]byte("Hello From the Server"))
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("Index Page"))
			return
		case r.URL.Path:
			w.Write([]byte(r.URL.Path))
			return
		}
	default:
		w.Write([]byte("404 not found"))
		return
	}
}

func (s *api) CreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Created Succesfully"))
}

func (s *api) DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Deleted Succesfully"))
}

func main() {
	api := &api{addr: ":8081"}
	mux:= http.NewServeMux()
	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("POST /createUser", api.CreateUsersHandler)
	mux.HandleFunc("POST /deleteUser", api.DeleteUsersHandler)
	
	srv.ListenAndServe()
}
