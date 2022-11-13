package recipe

import (
	"api/errorx"
	"api/model"
	"api/model/recipe/quantity"
	"api/model/recipe/stage"
	"context"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"net/http"
	"strings"
)

var _ RecipesModel = (*customRecipesModel)(nil)

type (
	// RecipesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecipesModel.
	RecipesModel interface {
		recipesModel
		StageModel() stage.StagesModel
		QuantityModel() quantity.QuantityModel
		InsertReturningId(ctx context.Context, recipeData *Recipes, quantitiesData []quantity.Quantity, stagesDate []stage.Stages) (int64, error)
		FindAll(ctx context.Context) ([]LiteRecipe, error)
		FindFiltered(ctx context.Context, withIngredients []int64, withoutIngredients []int64) ([]LiteRecipe, error)
		DeleteAllRecipes(ctx context.Context) (*sql.Result, error)
		FindByTitle(ctx context.Context, title string) ([]LiteRecipe, error)
		FindFullRecipeById(ctx context.Context, recipeId int64) (*Recipes, []quantity.QuantityWithName, []stage.Stages, error)
	}

	customRecipesModel struct {
		*defaultRecipesModel
		quantityModel quantity.QuantityModel
		stagesModel   stage.StagesModel
	}

	returningId struct {
		Id int64 `db:"id"`
	}

	LiteRecipe struct {
		Id    int64  `db:"id"`
		Title string `db:"title"`
	}
)

func (c customRecipesModel) StageModel() stage.StagesModel {
	return c.stagesModel
}

func (c customRecipesModel) QuantityModel() quantity.QuantityModel {
	return c.quantityModel
}

func (c customRecipesModel) FindFullRecipeById(ctx context.Context, recipeId int64) (*Recipes, []quantity.QuantityWithName, []stage.Stages, error) {
	recipe, err := c.FindOne(ctx, recipeId)
	if err != nil {
		return nil, nil, nil, err
	}

	quantities, err := c.quantityModel.FindByRecipe(ctx, recipeId)
	if err != nil {
		return nil, nil, nil, err
	}

	stages, err := c.stagesModel.FindByRecipe(ctx, recipeId)
	if err != nil {
		return nil, nil, nil, err
	}

	return recipe, quantities, stages, nil
}

func (c customRecipesModel) FindByTitle(ctx context.Context, title string) ([]LiteRecipe, error) {
	query := "select * from recipes where title = $1"
	var recipes []LiteRecipe
	err := c.conn.QueryRowsCtx(ctx, &recipes, query, title)
	switch err {
	case nil:
		if recipes == nil {
			return nil, model.ErrNotFound
		}
	case sqlx.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}

	return recipes, nil
}

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

// InsertReturningId do a transaction insertion of the recipe and the linked quantities and stages
func (c customRecipesModel) InsertReturningId(ctx context.Context, recipeData *Recipes, quantitiesData []quantity.Quantity, stagesDate []stage.Stages) (int64, error) {
	resp := returningId{}

	// Start transaction
	err := c.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		var query = fmt.Sprintf("insert into recipes (title, description, owner) values ($1, $2, $3) returning id")
		err := session.QueryRowCtx(ctx, &resp, query, recipeData.Title, recipeData.Description, recipeData.Owner)
		if err != nil {
			return err
		}

		// Inserting quantities
		for _, q := range quantitiesData {
			q.Recipe = resp.Id
			_, err := c.quantityModel.TransactInsert(ctx, session, &q)
			switch e := err.(type) {
			case *pq.Error:
				switch e.Code {
				case errorx.PgErrorCodeForeignKeyViolation:
					return errorx.NewCodeError(55, "Could not add this recipe, wrong ingredient", http.StatusNotAcceptable)
				default:
					return e
				}
			default:
				if e != nil {
					return e
				}
			}
		}

		// Inserting stages
		for _, s := range stagesDate {
			s.Recipe = resp.Id
			_, err := c.stagesModel.TransactInsert(ctx, session, &s)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return resp.Id, nil
}

// NewRecipesModel returns a model for the database table.
func NewRecipesModel(conn sqlx.SqlConn) RecipesModel {
	return &customRecipesModel{
		defaultRecipesModel: newRecipesModel(conn),
		quantityModel:       quantity.NewQuantityModel(conn),
		stagesModel:         stage.NewStagesModel(conn),
	}
}
