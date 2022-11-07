package recipe

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ RecipesModel = (*customRecipesModel)(nil)

type (
	// RecipesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecipesModel.
	RecipesModel interface {
		recipesModel
		InsertReturningId(ctx context.Context, data *Recipes) (int64, error)
		FindAll(ctx context.Context) ([]LiteRecipe, error)
		FindFiltered(ctx context.Context, withIngredients []int64, withoutIngredients []int64) ([]LiteRecipe, error)
		DeleteAllRecipes(ctx context.Context) (*sql.Result, error)
	}

	customRecipesModel struct {
		*defaultRecipesModel
	}

	returningId struct {
		Id int64 `db:"id"`
	}

	LiteRecipe struct {
		Id    int64  `db:"id"`
		Title string `db:"title"`
	}
)

func (c customRecipesModel) DeleteAllRecipes(ctx context.Context) (*sql.Result, error) {
	//goland:noinspection SqlWithoutWhere
	result, err := c.conn.ExecCtx(ctx, "delete from recipes")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// arrayToSqlString returns a string representation of the array parameter in the format "(1,2,3,4)"
func arrayToSqlString(array []int64) string {
	str := fmt.Sprintf("(%v)", array)
	str = strings.ReplaceAll(str, " ", ",")
	str = strings.Replace(str, "[", "", 1)
	str = strings.Replace(str, "]", "", 1)
	return str
}

func (c customRecipesModel) FindFiltered(ctx context.Context, withIngredients []int64, withoutIngredients []int64) ([]LiteRecipe, error) {
	queryBoth := `select r.id, r.title 
				from recipes r
				join quantity q on r.id = q.recipe
				where q.ingredient not in %s and q.ingredient in %s`
	queryWithout := `select r.id, r.title 
				from recipes r
				join quantity q on r.id = q.recipe
				where q.ingredient not in $1`
	queryWith := `select r.id, r.title 
				from recipes r
				join quantity q on r.id = q.recipe
				where q.ingredient in %s`

	var recipeList []LiteRecipe
	with := arrayToSqlString(withIngredients)
	without := arrayToSqlString(withoutIngredients)

	var err error
	if len(withIngredients) > 0 && len(withoutIngredients) > 0 {
		err = c.conn.QueryRowsCtx(ctx, &recipeList, fmt.Sprintf(queryBoth, without, with))
	} else if len(withIngredients) > 0 {
		err = c.conn.QueryRowsCtx(ctx, &recipeList, fmt.Sprintf(queryWith, with))
	} else {
		err = c.conn.QueryRowsCtx(ctx, &recipeList, fmt.Sprintf(queryWithout, without))
	}
	if err != nil {
		return nil, err
	}
	return recipeList, nil
}

func (c customRecipesModel) FindAll(ctx context.Context) ([]LiteRecipe, error) {
	query := fmt.Sprintf("select id, title from recipes")
	var recipeList []LiteRecipe
	err := c.conn.QueryRowsCtx(ctx, &recipeList, query)
	if err != nil {
		return nil, err
	}
	return recipeList, nil
}

func (c customRecipesModel) InsertReturningId(ctx context.Context, data *Recipes) (int64, error) {
	// TODO should add transaction management (and regroup quantity & stage insertion in one method)
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
