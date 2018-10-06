package routes

import (
	"github.com/bradford-hamilton/explore-the-chi/api/handlers"
	"github.com/bradford-hamilton/explore-the-chi/config"
	"github.com/go-chi/chi"
)

// Routes sets up the transaction routes, methods, and handlers
func Routes(dbConn *config.DBConfig) *chi.Mux {
	router := chi.NewRouter()

	// tx endpoints
	router.Get("/{txID}", handlers.GetTransaction(dbConn))
	router.Delete("/{txID}", handlers.DeleteTransaction(dbConn))
	router.Post("/", handlers.CreateTransaction(dbConn))
	router.Get("/", handlers.GetAllTransactions(dbConn))

	return router
}
