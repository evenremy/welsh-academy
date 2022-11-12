package handler

import (
	"api/model"
	"context"
	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"github.com/maxatome/go-testdeep/td"
	"net/http"
	"testing"
)

type recipeReq struct {
	Title          string                       `json:"title"`
	Description    string                       `json:"description"`
	IngredientList []linkIngredientWithQuantity `json:"ingredientList"`
	StageList      []stage                      `json:"stageList"`
}

type linkIngredientWithQuantity struct {
	IngredientId int64   `json:"ingredientId"`
	Quantity     float32 `json:"quantity"`
	Unit         string  `json:"unit"`
}

type stage struct {
	Order       int32  `json:"order"`
	Description string `json:"description"`
}

func TestAddRecipeHandlerOk(t *testing.T) {
	DeleteIngredients()
	var recipeId int64
	ingredientsId := AddSomeIngredients()
	recipeBodyQuery := createRecipeAuto(ingredientsId[0:3], 6)

	testApi := tdhttp.NewTestAPI(t, AddRecipeHandler(testCtx))
	testApi.AutoDumpResponse().
		Name("AddRecipe : should succeed").
		PostJSON("/recipe", recipeBodyQuery).
		CmpStatus(http.StatusOK).
		CmpJSONBody(td.JSON(`
{
	"title": "A Welsh recipe Auto",
	"id":$id
}
`, td.Tag("id", td.Catch(&recipeId, td.Gt(0)))))

	// Check the number of created quantities and stages in the db
	checkQuantitiesAndStages(t, recipeId, 3, 6)
}

func TestAddRecipeHandlerWithNoIngredientsAndStages(t *testing.T) {
	recipeBodyQuery := createRecipeAuto([]int64{}, 0)
	var recipeId int64

	testApi := tdhttp.NewTestAPI(t, AddRecipeHandler(testCtx))
	testApi.
		AutoDumpResponse().
		Name("AddRecipe : With No Ingredients And Stages").
		PostJSON("/recipe", recipeBodyQuery).
		CmpStatus(http.StatusOK).
		CmpJSONBody(td.JSON(`
{
	"title": "A Welsh recipe Auto",
	"id":$id
}
`, td.Tag("id", td.Catch(&recipeId, td.Gt(0)))))

	checkQuantitiesAndStages(t, recipeId, 0, 0)
}

// checkQuantitiesAndStages check the number of quantities and stages in the db for given recipeId
func checkQuantitiesAndStages(t *testing.T, recipeId int64, expectedQuantities int, expectedStages int) {
	// Check that there is no quantity and stage row in the db
	quantities, err := testCtx.QuantityModel.FindByRecipe(context.Background(), recipeId)
	if (expectedQuantities == 0 && err != model.ErrNotFound && len(quantities) != 0) || len(quantities) != expectedQuantities {
		t.Error("wrong number of quantities for the recipe", recipeId, "expected", expectedQuantities, "got", len(quantities), "error", err)
	}

	stages, err := testCtx.StageModel.FindByRecipe(context.Background(), recipeId)
	if (expectedStages == 0 && err != model.ErrNotFound && len(stages) != 0) || len(stages) != expectedStages {
		t.Error("wrong number of stages for the recipe", recipeId, "expected", expectedStages, "got", len(stages), "error", err)
	}
}

func createRecipeAuto(listIngredientId []int64, stageSize int) *recipeReq {
	recipe := recipeReq{
		Title:       "A Welsh recipe Auto",
		Description: "Blalkslkh, BlalkslkhBlalkslkhBlalkslkh BlalkslkhBlalkslkhBlalkslkhBlalkslkh Blalkslkh",
	}
	recipe.IngredientList = newIngredientList(listIngredientId)
	recipe.StageList = newStageList(stageSize)
	return &recipe
}

func newIngredientList(listId []int64) []linkIngredientWithQuantity {
	liwq := make([]linkIngredientWithQuantity, len(listId))
	for i, id := range listId {
		liwq[i].IngredientId = id
		liwq[i].Quantity = float32(i * 3)
		liwq[i].Unit = "ml"
	}
	return liwq
}

func newStageList(size int) []stage {
	stages := make([]stage, size)
	for i := range stages {
		stages[i].Order = int32(i)
		stages[i].Description = "Perferendis voluptatem sit aut accusantium consequatur. Consequatur voluptatem accusantium aut perferendis sit. Consequatur perferendis aut voluptatem sit accusantium. Consequatur sit voluptatem aut accusantium perferendis. Sit voluptatem aut perferendis consequatur accusantium. Aut voluptatem accusantium consequatur sit perferendis. Perferendis aut sit voluptatem accusantium consequatur. Voluptatem aut sit perferendis consequatur accusantium."
	}
	return stages
}
