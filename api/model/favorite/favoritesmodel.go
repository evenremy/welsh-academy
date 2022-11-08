package favorite

import (
	"api/model"
	"api/model/recipe"
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FavoritesModel = (*customFavoritesModel)(nil)

type (
	// FavoritesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFavoritesModel.
	FavoritesModel interface {
		favoritesModel
		InsertReturningId(ctx context.Context, f *Favorites) (*model.ReturningId, error)
		DeleteFavoriteByRecipeAndUserId(ctx context.Context, recipeId int64, userId int64) error
		FindFavRecipeByUser(ctx context.Context, userId int64) ([]recipe.Recipes, error)
	}

	customFavoritesModel struct {
		*defaultFavoritesModel
	}
)

func (c customFavoritesModel) FindFavRecipeByUser(ctx context.Context, userId int64) ([]recipe.Recipes, error) {
	query := `select r.id, r.title
from favorites f
join recipes r on f.recipe = r.id
where \"user\" = $1`

	var favs []recipe.Recipes
	err := c.conn.QueryRowsCtx(ctx, &favs, query, userId)
	if err != nil {
		return nil, err
	}
	return favs, nil
}

func (c customFavoritesModel) DeleteFavoriteByRecipeAndUserId(ctx context.Context, recipeId int64, userId int64) error {
	query := "delete from favorites where recipe = $1 and \"user\" = $2"
	result, err := c.conn.ExecCtx(ctx, query, recipeId, userId)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// should be 0 or 1 (unique constraint)
	if rowsAffected != 1 {
		return model.ErrNotFound
	}
	return nil
}

func (c customFavoritesModel) InsertReturningId(ctx context.Context, f *Favorites) (*model.ReturningId, error) {
	query := "insert into favorites (\"user\", recipe) values ($1, $2) returning id"
	id := model.ReturningId{}
	err := c.conn.QueryRowCtx(ctx, &id, query, f.User, f.Recipe)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

// NewFavoritesModel returns a model for the database table.
func NewFavoritesModel(conn sqlx.SqlConn) FavoritesModel {
	return &customFavoritesModel{
		defaultFavoritesModel: newFavoritesModel(conn),
	}
}
