package errorx

import "github.com/lib/pq"

const (
	PgErrorCodeForeignKeyViolation = pq.ErrorCode("23503") // FOREIGN KEY VIOLATION
)
