package transaction

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Transaction struct {
	Slug   string `json:"slug"`
	Input  string `json:"input"`
	Output string `json:"output"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{transactionID}", GetTransaction)
	router.Delete("/{transactionID}", DeleteTransaction)
	router.Post("/", CreateTransaction)
	router.Get("/", GetAllTransactions)

	return router
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	transactionID := chi.URLParam(r, "transactionID")

	transactions := Transaction{
		Slug:   transactionID,
		Input:  "0xNeatInput",
		Output: "0xNeatOutput",
	}

	render.JSON(w, r, transactions) // a chi router helper for serializing and returning json
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted transaction successfully"

	render.JSON(w, r, response)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created transaction successfully"

	render.JSON(w, r, response)
}

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	transactions := []Transaction{
		{
			Slug:   "ImmaSlugOne",
			Input:  "0xInputOne",
			Output: "0xOutputOne",
		},
		{
			Slug:   "ImmaSlugTwo",
			Input:  "0xInputTwo",
			Output: "0xOutputTwo",
		},
		{
			Slug:   "ImmaSlugThree",
			Input:  "0xInputThree",
			Output: "0xOutputThree",
		},
	}

	render.JSON(w, r, transactions)
}
