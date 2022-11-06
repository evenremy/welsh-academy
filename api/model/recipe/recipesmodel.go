package recipe

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RecipesModel = (*customRecipesModel)(nil)

type (
	// RecipesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecipesModel.
	RecipesModel interface {
		recipesModel
		InsertReturningId(ctx context.Context, data *Recipes) (int64, error)
	}

	customRecipesModel struct {
		*defaultRecipesModel
	}

	returningId struct {
		Id int64 `db:"id"`
	}
)

func (c customRecipesModel) InsertReturningId(ctx context.Context, data *Recipes) (int64, error) {
	var query = fmt.Sprintf("insert into %s (%s) values ($1, $2, $3) returning id", c.table, recipesRowsExpectAutoSet)
	resp := returningId{}
	err := c.conn.QueryRowCtx(ctx, &resp, query, data.Title, data.Description, data.Owner)
	return resp.Id, err
}

// NewRecipesModel returns a model for the database table.
func NewRecipesModel(conn sqlx.SqlConn) RecipesModel {
	return &customRecipesModel{
		defaultRecipesModel: newRecipesModel(conn),
	}
}
