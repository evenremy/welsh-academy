package quantity

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ QuantityModel = (*customQuantityModel)(nil)

type (
	// QuantityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQuantityModel.
	QuantityModel interface {
		quantityModel
	}

	customQuantityModel struct {
		*defaultQuantityModel
	}
)

// NewQuantityModel returns a model for the database table.
func NewQuantityModel(conn sqlx.SqlConn) QuantityModel {
	return &customQuantityModel{
		defaultQuantityModel: newQuantityModel(conn),
	}
}
