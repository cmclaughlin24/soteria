package postgres

import "database/sql"

func newNullInt64(val int) sql.NullInt64 {
	if val == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{Int64: int64(val), Valid: true}
}
