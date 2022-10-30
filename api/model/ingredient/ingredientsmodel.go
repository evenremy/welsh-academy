package ingredient

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ IngredientsModel = (*customIngredientsModel)(nil)

type (
	// IngredientsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIngredientsModel.
	IngredientsModel interface {
		ingredientsModel
		InsertReturningId(ctx context.Context, data *Ingredients) (int64, error)
	}

	customIngredientsModel struct {
		*defaultIngredientsModel
	}
)

type ReturningId struct {
	Id int64 `db:"id"`
}

func (c customIngredientsModel) InsertReturningId(ctx context.Context, data *Ingredients) (int64, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1) returning id", c.table, ingredientsRowsExpectAutoSet)
	resp := ReturningId{}
	err := c.conn.QueryRowCtx(ctx, &resp, query, data.Name)
	return resp.Id, err
}

// NewIngredientsModel returns a model for the database table.
func NewIngredientsModel(conn sqlx.SqlConn) IngredientsModel {
	return &customIngredientsModel{
		defaultIngredientsModel: newIngredientsModel(conn),
	}
}
