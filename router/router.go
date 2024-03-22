package router

import (
	"go-with-postgres/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock", middleware.GetAllStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/new-stock", middleware.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStocks).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/stock/delete/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")

	return router
}
