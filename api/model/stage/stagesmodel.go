package stage

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StagesModel = (*customStagesModel)(nil)

type (
	// StagesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStagesModel.
	StagesModel interface {
		stagesModel
		InsertStageFixed(ctx context.Context, data *Stages) (sql.Result, error)
	}

	customStagesModel struct {
		*defaultStagesModel
	}
)

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
