// Code generated by goctl. DO NOT EDIT!

package quantity

import (
	"api/model"
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	quantityFieldNames          = builder.RawFieldNames(&Quantity{}, true)
	quantityRows                = strings.Join(quantityFieldNames, ",")
	quantityRowsExpectAutoSet   = strings.Join(stringx.Remove(quantityFieldNames, "id", "create_time", "update_time", "create_at", "update_at"), ",")
	quantityRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(quantityFieldNames, "id", "create_time", "update_time", "create_at", "update_at"))
)

type (
	quantityModel interface {
		Insert(ctx context.Context, data *Quantity) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Quantity, error)
		FindOneByRecipeIngredient(ctx context.Context, recipe int64, ingredient int64) (*Quantity, error)
		Update(ctx context.Context, data *Quantity) error
		Delete(ctx context.Context, id int64) error
	}

	defaultQuantityModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Quantity struct {
		Recipe     int64           `db:"recipe"`
		Ingredient int64           `db:"ingredient"`
		Unit       string          `db:"unit"`
		Quantity   sql.NullFloat64 `db:"quantity"`
		Id         int64           `db:"id"`
	}
)

func newQuantityModel(conn sqlx.SqlConn) *defaultQuantityModel {
	return &defaultQuantityModel{
		conn:  conn,
		table: `"public"."quantity"`,
	}
}

func (m *defaultQuantityModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where id = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultQuantityModel) FindOne(ctx context.Context, id int64) (*Quantity, error) {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", quantityRows, m.table)
	var resp Quantity
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultQuantityModel) FindOneByRecipeIngredient(ctx context.Context, recipe int64, ingredient int64) (*Quantity, error) {
	var resp Quantity
	query := fmt.Sprintf("select %s from %s where recipe = $1 and ingredient = $2 limit 1", quantityRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, recipe, ingredient)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultQuantityModel) Insert(ctx context.Context, data *Quantity) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4)", m.table, quantityRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Recipe, data.Ingredient, data.Unit, data.Quantity)
	return ret, err
}

func (m *defaultQuantityModel) Update(ctx context.Context, newData *Quantity) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, quantityRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Id, newData.Recipe, newData.Ingredient, newData.Unit, newData.Quantity)
	return err
}

func (m *defaultQuantityModel) tableName() string {
	return m.table
}
