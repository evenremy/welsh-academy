package logic

import (
	"api/errorx"
	"api/model/recipe"
	"api/model/recipe/quantity"
	"api/model/recipe/stage"
	"context"
	"database/sql"
	"net/http"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRecipeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRecipeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRecipeLogic {
	return &AddRecipeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddRecipe adds the recipe from request parameter AddRecipeReq and returns a AddRecipeReply (id + title)
// IngredientList should be consistent with existing ingredients
func (l *AddRecipeLogic) AddRecipe(req *types.AddRecipeReq) (resp *types.AddRecipeReply, err error) {
	// Recipe
	newRecipe := recipe.Recipes{
		Title: req.Title,
		Owner: sql.NullInt64{},
	}
	err = newRecipe.Description.Scan(req.Description)
	if err != nil {
		return nil, err
	}

	// Ingredients
	newQuantities, err := l.quantitiesFromIngredients(req.IngredientList)
	if err != nil {
		return nil, err
	}

	// Stages
	newStages, err := l.stagesFromStagesRequest(req.StageList)
	if err != nil {
		return nil, err
	}

	recipeId, err := l.svcCtx.RecipeModel.InsertReturningId(l.ctx, &newRecipe, newQuantities, newStages)
	switch err.(type) {
	case nil:
	case errorx.ApiError:
		return nil, err
	default:
		l.Logger.Errorf("Unknown Error during insertion for Recipe", newRecipe)
		return nil, errorx.NewCodeError(45, "unable to add recipe, unknown error", http.StatusBadRequest)
	}

	reply := types.AddRecipeReply{
		Id:    recipeId,
		Title: req.Title,
	}

	return &reply, nil
}

func (l *AddRecipeLogic) quantitiesFromIngredients(ingredients []types.LinkIngredientWithQuantity) ([]quantity.Quantity, error) {
	quantities := make([]quantity.Quantity, len(ingredients))
	for i, ingredient := range ingredients {
		quantities[i] = quantity.Quantity{
			Ingredient: ingredient.IngredientId,
			Unit:       ingredient.Unit,
			Quantity:   sql.NullFloat64{},
		}
		err := quantities[i].Quantity.Scan(ingredient.Quantity)
		if err != nil {
			return nil, err
		}
	}

	return quantities, nil
}

func (l *AddRecipeLogic) stagesFromStagesRequest(stagesReq []types.Stage) ([]stage.Stages, error) {
	stageList := make([]stage.Stages, len(stagesReq))
	for i, s := range stagesReq {
		stageList[i] = stage.Stages{
			Order:       int64(s.Order),
			Description: sql.NullString{},
		}
		err := stageList[i].Description.Scan(s.Description)
		if err != nil {
			return nil, err
		}

	}
	return stageList, nil
}
