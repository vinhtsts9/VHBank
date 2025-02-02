package impl

import (
	"Golang-Masterclass/simplebank/internal/database"
	"Golang-Masterclass/simplebank/internal/models"
	"context"
	"database/sql"
)

type sTransfer struct {
	r  *database.Queries
	db *sql.DB
}

func NewTransferImpl(r *database.Queries, db *sql.DB) *sTransfer {
	return &sTransfer{
		r:  r,
		db: db,
	}
}
func (s *sTransfer) TransferTX(ctx context.Context, arg models.TransferTxParams) (models.TransferTxResult, error) {
	var result models.TransferTxResult
	err := ExecTx(ctx, s.db, func(q *database.Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, database.CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, database.CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}
		return err
	})
	return result, err
}
