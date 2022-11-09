package handler

import (
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

func TestAddRecipeHandlerAutoJson(t *testing.T) {
	DeleteIngredients()
	ingredientsId := AddSomeIngredients()
	recipeBodyQuery := createRecipeAuto(ingredientsId[0:3], 6)

	testApi := tdhttp.NewTestAPI(t, AddRecipeHandler(testCtx))
	testApi.AutoDumpResponse()
	testApi.PostJSON("/recipe", recipeBodyQuery).
		Name("Add recipe").
		CmpStatus(http.StatusOK).
		CmpJSONBody(td.JSON(`
{
	"title": "A Welsh recipe Auto",
	"id":$id
}
`, td.Tag("id", td.Gt(0))))
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
