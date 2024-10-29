package sqlx

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// WithTransaction is like WithTransactionReturning except does not return a value
func WithTransaction(db *sqlx.DB, workFunc func(*sqlx.Tx) error) error {
	_, err := WithTransactionReturning(db, func(tx *sqlx.Tx) (bool, error) {
		return false, workFunc(tx)
	})
	return err
}

// WithTransactionReturning executes the provided function against the provided DB, rolling back if the given function
// returns error, or committing otherwise. The return value from the work function will be returned on success
func WithTransactionReturning[T any](db *sqlx.DB, workFunc func(tx *sqlx.Tx) (T, error)) (T, error) {
	var empty T

	txn, err := db.Beginx()
	if err != nil {
		return empty, fmt.Errorf("error opening transaction: %w", err)
	}

	out, workErr := workFunc(txn)
	if workErr != nil {
		if err := txn.Rollback(); err != nil {
			return empty, fmt.Errorf("error rolling back: %w, rollback caused by: %w", err, workErr)
		}
		return empty, workErr
	}

	if err := txn.Commit(); err != nil {
		return empty, fmt.Errorf("error committing: %w", err)
	}

	return out, nil
}
