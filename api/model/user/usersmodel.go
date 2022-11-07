package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		InsertReturningId(ctx context.Context, username string, isAdmin bool) (int64, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

func (c customUsersModel) InsertReturningId(ctx context.Context, username string, isAdmin bool) (id int64, err error) {
	query := "insert into users (username, is_admin) VALUES ($1, $2) returning id"
	err = c.conn.QueryRowCtx(ctx, &id, query, username, isAdmin)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn),
	}
}
