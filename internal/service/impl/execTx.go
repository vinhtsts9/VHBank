package impl

import (
	"Golang-Masterclass/simplebank/internal/database"
	"context"
	"database/sql"
	"fmt"
)

func ExecTx(ctx context.Context, db *sql.DB, fn func(*database.Queries) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := database.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
