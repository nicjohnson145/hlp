package sqlx

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// ScanRows scans each row object into the indicated struct, returning a list of all objects. It delegates the closing
// of the rows object to the caller
func ScanRows[T any](rows *sqlx.Rows) ([]T, error) {
	outRows := make([]T, 0)

	err := IScanRows(rows, func(out T) error {
		outRows = append(outRows, out)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error iterating: %w", err)
	}

	return outRows, nil
}

// IScanRows scans each row into the indicated struct, and calls the process function on it, stopping on the first
// error. It delegates the closing of the rows object to the caller
func IScanRows[T any](rows *sqlx.Rows, processFunc func(T) error) error {
	for rows.Next() {
		var out T
		if err := rows.StructScan(&out); err != nil {
			return fmt.Errorf("error scanning: %w", err)
		}

		if err := processFunc(out); err != nil {
			return fmt.Errorf("error processing: %w", err)
		}
	}
	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}
