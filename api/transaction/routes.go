package transaction

import (
	"github.com/bradford-hamilton/explore-the-chi/config"
	"github.com/go-chi/chi"
)

// Routes sets up the transaction routes, methods, and handlers
func Routes(dbConn *config.DBConfig) *chi.Mux {
	router := chi.NewRouter()

	// transaction endpoints
	router.Get("/{transactionID}", GetTransaction(dbConn))
	router.Delete("/{transactionID}", DeleteTransaction(dbConn))
	router.Post("/", CreateTransaction(dbConn))
	router.Get("/", GetAllTransactions(dbConn))

	return router
}
