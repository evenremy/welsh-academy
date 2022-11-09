package handler

import (
	"api/internal/types"
	"context"
	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"github.com/maxatome/go-testdeep/td"
	"net/http"
	"testing"
)

func TestAddFavoriteRecipeHandlerErrors(t *testing.T) {
	testApi := tdhttp.NewTestAPI(t, AddFavoriteRecipeHandler(testCtx))
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

func TestAddFavoriteRecipeHandlerExistingUser(t *testing.T) {
	userId, err := testCtx.UserModel.InsertReturningId(context.Background(), "abc", false)
	if err != nil {
		return
	}
	testApi := tdhttp.NewTestAPI(t, AddFavoriteRecipeHandler(testCtx))
	body := &types.FavReq{
		UserId:   userId,
		RecipeId: 999999,
	}
	testApi.PostJSON("/favorite_recipe", body).
		CmpStatus(http.StatusForbidden).
		CmpJSONBody(td.JSON(`
{
	"code": 23,
	"msg": NotEmpty()
}
`))
}
