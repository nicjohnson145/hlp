package sqlx

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var (
	ErrNotFoundError           = errors.New("not found")
	ErrUnexpectedRowCountError = errors.New("unexpected row count")
)

// SelectNamedCtx executes a query with named paramters, and scans the results into the supplied struct
func SelectNamedCtx[T any](ctx context.Context, db sqlx.ExtContext, query string, args any) ([]T, error) {
	rows, err := sqlx.NamedQueryContext(ctx, db, query, args)
	if err != nil {
		return nil, fmt.Errorf("error querying: %w", err)
	}
	defer rows.Close()

	return ScanRows[T](rows)
}

// RequireExactSelectNamedCtx is like SelectNamedCtx, except it enforces that the number of rows returned matches an
// expected value
func RequireExactSelectNamedCtx[T any](ctx context.Context, expected int, db sqlx.ExtContext, query string, args any) ([]T, error) {
	rows, err := SelectNamedCtx[T](ctx, db, query, args)
	if err != nil {
		return nil, err
	}

	if count := len(rows); count != expected {
		if expected > 0 && count == 0 {
			return nil, ErrNotFoundError
		}
		return nil, fmt.Errorf("%w: expected %v, got %v", ErrUnexpectedRowCountError, expected, count)
	}

	return rows, nil
}
