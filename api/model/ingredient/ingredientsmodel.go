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
		DeleteAllIngredients(ctx context.Context) (int64, error)
		FindAll(ctx context.Context) ([]Ingredients, error)
	}

	customIngredientsModel struct {
		*defaultIngredientsModel
	}
)

func (c customIngredientsModel) FindAll(ctx context.Context) ([]Ingredients, error) {
	var ingredients []Ingredients
	query := fmt.Sprintf("select * from %s", c.table)
	err := c.conn.QueryRowsCtx(ctx, &ingredients, query)
	if err != nil {
		return nil, err
	}
	return ingredients, nil
}

type ReturningId struct {
	Id int64 `db:"id"`
}

//goland:noinspection SqlWithoutWhere
func (c customIngredientsModel) DeleteAllIngredients(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("Delete from %s", c.table)
	result, err := c.conn.ExecCtx(ctx, query)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (c customIngredientsModel) InsertReturningId(ctx context.Context, data *Ingredients) (int64, error) {
	var query = fmt.Sprintf("insert into %s (%s) values ($1) returning id", c.table, ingredientsRowsExpectAutoSet)
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
