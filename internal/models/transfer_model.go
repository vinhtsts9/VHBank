package models

import (
	"Golang-Masterclass/simplebank/internal/database"
	"time"
)

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    database.Transfer `json:"transfer"`
	FromAccount database.Account  `json:"from_account"`
	ToAccount   database.Account  `json:"to_account"`
	FromEntry   database.Entry    `json:"from_entry"`
	ToEntry     database.Entry    `json:"to_entry"`
}
