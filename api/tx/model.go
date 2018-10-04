package tx

// Transaction is the data model for a transaction
type Transaction struct {
	ID     string `dynamo:"Id"`
	Input  string `dynamo:"input"`
	Output string `dynamo:"output"`
}
