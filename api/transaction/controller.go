package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bradford-hamilton/explore-the-chi/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// GetTransaction returns a single Transaction struct as JSON
func GetTransaction(dbConn *config.DBConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transactionID := chi.URLParam(r, "transactionID")
		table := dbConn.DB.Table("btc-transaction")

		var tx Transaction
		err := table.Get("Id", transactionID).One(&tx)
		if err != nil {
			fmt.Println(err)
		}

		render.JSON(w, r, tx) // a chi router helper for serializing and returning json
	}
}

// DeleteTransaction deletes a transaction and returns a success JSON message
func DeleteTransaction(dbConn *config.DBConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]string)
		response["message"] = "Deleted transaction successfully"

		render.JSON(w, r, response)
	}
}

// CreateTransaction creates a transaction and returns a JSON success message.
// Submit format below:
// {
//		"ID": "0xSuperC00lId123",
//		"Input":  "0xNeatInput",
//		"Output": "0xNeatOutput"
// }
// Ensure content type is application/json
func CreateTransaction(dbConn *config.DBConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var data Transaction
		err := decoder.Decode(&data)
		if err != nil {
			fmt.Println(err)
		}

		table := dbConn.DB.Table("btc-transaction")
		transaction := Transaction{
			ID:     data.ID,
			Input:  data.Input,
			Output: data.Output,
		}

		err = table.Put(transaction).Run()
		if err != nil {
			fmt.Println(err)
		}

		response := make(map[string]string)
		response["message"] = "Transaction with ID: " + data.ID + " created successfully"

		render.JSON(w, r, response) // a chi router helper for serializing and returning json
	}
}

// GetAllTransactions returns a slice of Transaction structs containing
// transactions. Response is JSON.
func GetAllTransactions(dbConn *config.DBConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		table := dbConn.DB.Table("btc-transaction")

		var txs []Transaction
		err := table.Scan().All(&txs)
		if err != nil {
			fmt.Println(err)
		}

		render.JSON(w, r, txs) // a chi router helper for serializing and returning json
	}
}
