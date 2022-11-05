package stage

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ StagesModel = (*customStagesModel)(nil)

type (
	// StagesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStagesModel.
	StagesModel interface {
		stagesModel
	}

	customStagesModel struct {
		*defaultStagesModel
	}
)

// NewStagesModel returns a model for the database table.
func NewStagesModel(conn sqlx.SqlConn) StagesModel {
	return &customStagesModel{
		defaultStagesModel: newStagesModel(conn),
	}
}
