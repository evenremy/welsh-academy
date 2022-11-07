package quantity

import (
	"api/model"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ QuantityModel = (*customQuantityModel)(nil)

type (
	// QuantityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQuantityModel.
	QuantityModel interface {
		quantityModel
		FindByRecipe(ctx context.Context, recipeId int64) ([]Quantity, error)
	}

	customQuantityModel struct {
		*defaultQuantityModel
	}
)

func (c customQuantityModel) FindByRecipe(ctx context.Context, recipeId int64) ([]Quantity, error) {
	var resp []Quantity
	query := fmt.Sprintf("select %s from %s where recipe = $1", quantityRows, c.table)
	err := c.conn.QueryRowsCtx(ctx, &resp, query, recipeId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

// NewQuantityModel returns a model for the database table.
func NewQuantityModel(conn sqlx.SqlConn) QuantityModel {
	return &customQuantityModel{
		defaultQuantityModel: newQuantityModel(conn),
	}
}
