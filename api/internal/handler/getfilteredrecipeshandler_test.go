package handler

import (
	"api/internal/types"
	"api/model/quantity"
	"api/model/recipe"
	"context"
	"database/sql"
	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"github.com/maxatome/go-testdeep/td"
	"net/http"
	"testing"
)

func TestGetFilteredRecipesHandler(t *testing.T) {
	DeleteAllRecipes()
	listId := AddSomeIngredients()
	// Create 2 recipes with 1 common ingredients -> listId[3]
	recipeId1 := addRecipe(listId[0:4])
	recipeId2 := addRecipe(listId[3:6])
	// Create 1 recipe without any common ingredients
	//recipeId3 := addRecipe(listId[6:10])

	testApi := tdhttp.NewTestAPI(t, GetFilteredRecipesHandler(testCtx)).
		AutoDumpResponse()
	// Match with recipeId1 & recipeId2
	filterReq1 := types.IngredientConstraintsReq{
		WithIngredientIdList:    []int64{listId[3]},
		WithoutIngredientIdList: []int64{},
	}
	// Match with recipeId2
	/*filterReq2 := types.IngredientConstraintsReq{
		WithIngredientIdList:    []int64{listId[3]},
		WithoutIngredientIdList: []int64{listId[4]},
	}*/

	testApi.PostJSON("/filtered_recipes", filterReq1).
		CmpStatus(http.StatusOK).
		CmpJSONBody(td.JSON(`{
"recipeList": [
	{
	"id": $id1,
	"title": $title
	},
	{	
	"id": $id2,
	"title": $title}
	]
}
`,
			td.Tag("id1", recipeId1),
			td.Tag("id2", recipeId2),
			td.Tag("title", td.NotEmpty())))

}

func addRecipe(ingredientsId []int64) int64 {
	recipeId, _ := testCtx.RecipeModel.InsertReturningId(context.Background(), &recipe.Recipes{
		Title:       "Test recipe",
		Description: sql.NullString{},
		Owner:       sql.NullInt64{},
	})

	for _, id := range ingredientsId {
		_, _ = testCtx.QuantityModel.Insert(context.Background(), &quantity.Quantity{
			Recipe:     recipeId,
			Ingredient: id,
			Unit:       "cup",
			Quantity:   sql.NullFloat64{},
		})
	}

	return recipeId
}
