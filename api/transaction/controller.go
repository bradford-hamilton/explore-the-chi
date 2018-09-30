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

		transaction := Transaction{
			ID:     transactionID,
			Input:  "0xNeatInput",
			Output: "0xNeatOutput",
		}

		render.JSON(w, r, transaction) // a chi router helper for serializing and returning json
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
		// TODO: check if we want table reference above this return to run once on program start.
		// We will have to see if that works - and if it's better performaance or if we need to define
		// the table like below each time the api is called
		table := dbConn.DB.Table("btc-transaction")
		transaction := Transaction{
			ID:     "neat-id",
			Input:  "0xInputThree",
			Output: "0xOutputThree",
		}

		err := table.Put(transaction).Run()

		if err != nil {
			fmt.Println(err)
		}

		var result Transaction
		err = table.Get("Id", transaction.ID).One(&result)

		render.JSON(w, r, result) // a chi router helper for serializing and returning json
	}
}
