package stage

import (
	"api/model"
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StagesModel = (*customStagesModel)(nil)

type (
	// StagesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStagesModel.
	StagesModel interface {
		stagesModel
		InsertStageFixed(ctx context.Context, data *Stages) (sql.Result, error)
		FindByRecipe(ctx context.Context, id int64) ([]Stages, error)
	}

	customStagesModel struct {
		*defaultStagesModel
	}
)

func (c customStagesModel) FindByRecipe(ctx context.Context, id int64) ([]Stages, error) {
	var resp []Stages
	query := fmt.Sprintf("select %s from %s where recipe = $1", stagesRows, c.table)
	err := c.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

func (c customStagesModel) InsertStageFixed(ctx context.Context, data *Stages) (sql.Result, error) {
	query := "insert into stages (recipe, \"order\", description) values ($1, $2, $3)"
	ret, err := c.conn.ExecCtx(ctx, query, data.Recipe, data.Order, data.Description)
	return ret, err
}

// NewStagesModel returns a model for the database table.
func NewStagesModel(conn sqlx.SqlConn) StagesModel {
	return &customStagesModel{
		defaultStagesModel: newStagesModel(conn),
	}
}
