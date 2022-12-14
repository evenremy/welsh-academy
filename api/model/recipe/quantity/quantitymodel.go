package quantity

import (
	"api/model"
	"context"
	"database/sql"
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
		FindByRecipe(ctx context.Context, recipeId int64) ([]QuantityWithName, error)
		TransactInsert(ctx context.Context, session sqlx.Session, q *Quantity) (sql.Result, error)
	}

	customQuantityModel struct {
		*defaultQuantityModel
	}

	QuantityWithName struct {
		Recipe         int64           `db:"recipe"`
		IngredientId   int64           `db:"ingredient"`
		IngredientName string          `db:"name"`
		Unit           string          `db:"unit"`
		Quantity       sql.NullFloat64 `db:"quantity"`
		Id             int64           `db:"id"`
	}
)

func (c customQuantityModel) TransactInsert(ctx context.Context, session sqlx.Session, q *Quantity) (sql.Result, error) {
	query := fmt.Sprintf("insert into quantity (recipe, ingredient, unit, quantity) values ($1, $2, $3, $4)")
	ret, err := session.ExecCtx(ctx, query, q.Recipe, q.Ingredient, q.Unit, q.Quantity)
	return ret, err
}

func (c customQuantityModel) FindByRecipe(ctx context.Context, recipeId int64) ([]QuantityWithName, error) {
	var resp []QuantityWithName
	query := `select q.*, i.name
				from quantity q
			    join ingredients i on i.id = q.ingredient 
				where recipe = $1`
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
