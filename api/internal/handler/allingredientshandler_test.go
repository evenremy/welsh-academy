package handler

import (
	"api/internal/types"
	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"github.com/maxatome/go-testdeep/td"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/maxatome/go-testdeep/helpers/tdhttp"
	_ "github.com/maxatome/go-testdeep/td"
)

func TestAllIngredientsHandler(t *testing.T) {
	DeleteAllRecipes()
	DeleteIngredients()

	var id int64
	expectedIngredientNumber := 3
	addNIngredients(expectedIngredientNumber)

	testAPI := tdhttp.NewTestAPI(t, AllIngredientsHandler(testCtx))
	testAPI.Get("/ingredients").
		Name("Get the n ingredients previously created").
		CmpStatus(http.StatusOK).
		CmpJSONBody(td.JSON(`
{
	"ingredientList": [{
			"id":$id,
			"name":$name,
		},{
			"id":$id,
			"name":$name,
		},{
			"id":$id,
			"name":$name,
		}]
}`,
			td.Tag("id", td.Catch(&id, td.Gt(0))),
			td.Tag("name", td.NotEmpty())))
}

func addNIngredients(n int) {
	for i := 0; i < n; i++ {
		dataReader := prepareFakeJson(&types.AddIngredientReq{})

		req := httptest.NewRequest(TEST_METHOD, TEST_TARGET, dataReader)
		req.Header.Add("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		httpHandler := AddIngredientHandler(testCtx)
		httpHandler(rr, req)
	}
}

func TestNoIngredient(t *testing.T) {
	DeleteAllRecipes()
	DeleteIngredients()
	testAPI := tdhttp.NewTestAPI(t, AllIngredientsHandler(testCtx))
	testAPI.Get("/ingredients").
		Name("Get the 0 ingredient result").
		CmpStatus(http.StatusOK).
		CmpJSONBody(td.JSON(`
{
	"ingredientList": []
}`))
}
