package routes

import (
	"bookstore/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/maxId/", controllers.GetMaxID).Methods("GET")
	router.HandleFunc("/orders/", controllers.GetOrders).Methods("GET")
	router.HandleFunc("/orders/", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/bookOrders/{bookId}", controllers.GetBookOrders).Methods("GET")

}
