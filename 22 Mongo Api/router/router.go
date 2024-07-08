package route

import (
	"github.com/gorilla/mux"

	"myMongoApi/controller"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controller.CreateOneMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	r.HandleFunc("/api/deleteall", controller.DeleteAllMovies).Methods("DELETE")

	return r
}
