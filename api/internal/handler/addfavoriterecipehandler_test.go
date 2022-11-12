package handler

import (
	"api/internal/types"
	"api/model"
	"api/model/recipe"
	"context"
	"database/sql"
	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"github.com/maxatome/go-testdeep/td"
	"net/http"
	"testing"
)

func TestAddFavoriteRecipeHandlerErrors(t *testing.T) {
	testApi := tdhttp.NewTestAPI(t, AddFavoriteRecipeHandler(testCtx))
	testApi.Name("Add favorite : Not existing user should be rejected")
	body := &types.FavReq{
		UserId:   999999,
		RecipeId: 99999,
	}
	testApi.PostJSON("/favorite_recipe", body).
		CmpStatus(http.StatusForbidden).
		CmpJSONBody(td.JSON(`
{
	"code": 22,
	"msg": NotEmpty()
}
`))
}

func TestAddFavoriteRecipeHandlerNotExistingUser(t *testing.T) {
	userId := createOrGetAbcUserId(t)

	testApi := tdhttp.NewTestAPI(t, AddFavoriteRecipeHandler(testCtx))
	testApi.Name("Add favorite : Not existing recipe should be rejected")
	body := &types.FavReq{
		UserId:   userId,
		RecipeId: 999999,
	}

	testApi.PostJSON("/favorite_recipe", body).
		CmpStatus(http.StatusNotModified).
		CmpJSONBody(td.JSON(`
{
	"code": 23,
	"msg": NotEmpty()
}
`))
}

func TestAddFavoriteRecipeHandlerWorking(t *testing.T) {
	userId := createOrGetAbcUserId(t)

	// create recipe
	recipeData := &recipe.Recipes{
		Title:       "Fav recipe",
		Description: sql.NullString{},
		Owner:       sql.NullInt64{},
	}
	recipeId, _ := testCtx.RecipeModel.InsertReturningId(context.Background(), recipeData)

	body := &types.FavReq{
		UserId:   userId,
		RecipeId: recipeId,
	}

	testApi := tdhttp.NewTestAPI(t, AddFavoriteRecipeHandler(testCtx))
	testApi.Name("Add favorite : With existing user and recipe, should succeed")
	testApi.PostJSON("/favorite_recipe", body).
		CmpStatus(http.StatusOK).
		NoBody()
}

func createOrGetAbcUserId(t *testing.T) int64 {
	var userId int64
	user, err := testCtx.UserModel.FindOneByUsername(context.Background(), "abc")
	switch err {
	case nil:
		userId = user.Id
	case model.ErrNotFound:
		userId, _ = testCtx.UserModel.InsertReturningId(context.Background(), "abc", false)
	default:
		t.Error(err)
	}

	return userId
}
