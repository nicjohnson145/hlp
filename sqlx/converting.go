package sqlx

import (
	"database/sql"
)

// PointerToSqlNull converts a pointer to a type to a sql.Null for that type
func PointerToSqlNull[T any](x *T) sql.Null[T] {
	if x != nil {
		return sql.Null[T]{
			V:     *x,
			Valid: true,
		}
	}
	return sql.Null[T]{
		Valid: false,
	}
}

// SqlNullToPointer is the inverse of PointerToSqlNull, converting a sql.Null to a pointer to a type
func SqlNullToPointer[T any](x sql.Null[T]) *T {
	if !x.Valid {
		return nil
	}
	return &x.V
}
