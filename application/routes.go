package application
import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/diouri844/order_api/handler"
)
// define the function setup of app router :
func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	// add the middelware logger to the current router :
	router.Use(middleware.Logger)
	// define endpoint:
	router.Get("/Ping", func( w http.ResponseWriter, r *http.Request){
			w.WriteHeader(http.StatusOK)
	})
	// define endpoint with orders prefixer :
	router.Route("/Orders", loadOrderRoutes)
	return	router
}

func loadOrderRoutes(router chi.Router){
	// create an instance of order handler :
	orderHandler := &handler.Order{}
	// list all the orders :
	router.Get("/", orderHandler.List)
	// create new order :
	router.Post("/", orderHandler.Create)
	// get order by id :
	router.Get("/{id}", orderHandler.GetById)
	// Update order by id :
	router.Put("/{id}", orderHandler.UpdateById)
	// delete order by id :
	router.Delete("/{id}", orderHandler.DeleteById)
}