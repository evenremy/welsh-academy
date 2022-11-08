package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

type ReturningId struct {
	Id int64 `db:"id"`
}
