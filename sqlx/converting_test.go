package sqlx

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPointerToSqlNull(t *testing.T) {
	t.Run("not null", func(t *testing.T) {
		var x int = 7
		require.Equal(t, sql.Null[int]{V: 7, Valid: true}, PointerToSqlNull(&x))
	})

	t.Run("null", func(t *testing.T) {
		require.Equal(t, sql.Null[int]{V: 0, Valid: false}, PointerToSqlNull[int](nil))
	})
}

func TestSqlNullToPointer(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		require.Nil(t, SqlNullToPointer(sql.Null[int]{Valid: false}))
	})

	t.Run("not null", func(t *testing.T) {
		x := 7
		require.Equal(t, &x, SqlNullToPointer(sql.Null[int]{V: 7, Valid: true}))
	})
}
