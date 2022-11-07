package logic

import (
	quantity2 "api/model/quantity"
	"api/model/recipe"
	stage2 "api/model/stage"
	"context"
	"database/sql"

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

	recipeId, err := l.svcCtx.RecipeModel.InsertReturningId(l.ctx, &newRecipe)
	if err != nil {
		l.Logger.Errorf("Error during insertion for Recipe : id = %d", recipeId)
		return nil, err
	}

	// Stages
	err = l.addStages(req.StageList, recipeId)
	if err != nil {
		return nil, err
	}

	// Ingredients
	err = l.addIngredients(req.IngredientList, recipeId)
	if err != nil {
		return nil, err
	}

	reply := types.AddRecipeReply{
		Id:    recipeId,
		Title: req.Title,
	}

	return &reply, nil
}

func (l *AddRecipeLogic) addIngredients(ingredients []types.LinkIngredientWithQuantity, recipeId int64) error {
	for i, ingredient := range ingredients {
		newQuantity := quantity2.Quantity{
			Recipe:     recipeId,
			Ingredient: ingredient.IngredientId,
			Unit:       ingredient.Unit,
			Quantity:   sql.NullFloat64{},
		}
		err := newQuantity.Quantity.Scan(ingredient.Quantity)
		if err != nil {
			return err
		}
		_, err = l.svcCtx.QuantityModel.Insert(l.ctx, &newQuantity)
		if err != nil {
			l.Logger.Errorf("Error during insertion in Quantity at %d/%d ingredients, for: %+v", i, len(ingredients), newQuantity)
			return err
		}
	}
	return nil
}

func (l *AddRecipeLogic) addStages(stages []types.Stage, recipeId int64) error {
	for i, stage := range stages {
		newStage := stage2.Stages{
			Recipe:      recipeId,
			Order:       int64(stage.Order),
			Description: sql.NullString{},
		}
		err := newStage.Description.Scan(stage.Description)
		if err != nil {
			return err
		}

		_, err = l.svcCtx.StageModel.InsertStageFixed(l.ctx, &newStage)
		if err != nil {
			l.Logger.Errorf("Error during insertion in Stages at %d/%d stages, for: %+v", i, len(stages), newStage)
			return err
		}
	}
	return nil
}
