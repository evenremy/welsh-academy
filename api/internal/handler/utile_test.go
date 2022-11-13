package handler

import (
	"api/errorx"
	"api/internal/config"
	"api/internal/svc"
	ingredient2 "api/model/ingredient"
	"context"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// Give access to models/db
var testCtx *svc.ServiceContext

const TestConfigFilePath = "../../etc/welsh-academy-api.yaml"

func init() {
	// testCtx initialisation
	if testCtx != nil {
		return
	}
	c := config.Config{}
	conf.MustLoad(TestConfigFilePath, &c)
	testCtx = svc.NewServiceContext(c)

	// set error handler to test json error formatting
	httpx.SetErrorHandler(errorx.ErrorHandler)
}

func DeleteIngredients() {
	_, _ = testCtx.IngredientModel.DeleteAllIngredients(context.Background())
}

func DeleteAllRecipes() {
	_, _ = testCtx.RecipeModel.DeleteAllRecipes(context.Background())
}

// AddSomeIngredients Insert at least 10 ingredients and returns all ingredients id from the db (could be more than 10)
func AddSomeIngredients() []int64 {
	var listid []int64
	someIngredients := []ingredient2.Ingredients{
		{Name: "Cheddar"},
		{Name: "Ham"},
		{Name: "Mustard"},
		{Name: "Bread"},
		{Name: "Black pepper"},
		{Name: "Bacon"},
		{Name: "Eggs"},
		{Name: "Beer"},
		{Name: "Comté"},
		{Name: "Mayonnaise"},
	}
	for _, si := range someIngredients {
		_, err := testCtx.IngredientModel.InsertReturningId(context.Background(), &si)
		if err != nil {
			break
		}
	}

	allIngredients, err := testCtx.IngredientModel.FindAll(context.Background())
	if err != nil {
		return nil
	}

	for _, ingr := range allIngredients {
		listid = append(listid, ingr.Id)
	}
	return listid
}
