package logic

import (
	"api/model/quantity"
	"api/model/stage"
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecipeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRecipeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecipeLogic {
	return &GetRecipeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRecipeLogic) GetRecipe(req *types.RecipeReq) (resp *types.FullRecipeReply, err error) {
	recipe, err := l.svcCtx.RecipeModel.FindOne(l.ctx, req.RecipeId)
	if err != nil {
		return nil, err
	}
	reply := types.FullRecipeReply{
		Id:    recipe.Id,
		Title: recipe.Title,
	}
	if recipe.Description.Valid {
		reply.Description = recipe.Description.String
	}

	quantities, err := l.svcCtx.QuantityModel.FindByRecipe(l.ctx, req.RecipeId)
	if err != nil {
		return nil, err
	}
	reply.IngredientList = getIngredientListFromModel(quantities)

	stages, err := l.svcCtx.StageModel.FindByRecipe(l.ctx, req.RecipeId)
	if err != nil {
		return nil, err
	}
	reply.StageList = getStageListFromModel(stages)

	return &reply, nil
}

func getStageListFromModel(stages []stage.Stages) []types.Stage {
	replyStages := make([]types.Stage, len(stages))
	for i, s := range stages {
		replyStages[i].Order = int32(s.Order)
		if s.Description.Valid {
			replyStages[i].Description = s.Description.String
		}
	}

	return replyStages
}

func getIngredientListFromModel(quantities []quantity.QuantityWithName) []types.IngredientNameWithQuantity {
	replyIngredients := make([]types.IngredientNameWithQuantity, len(quantities))
	for i, q := range quantities {
		replyIngredients[i].Unit = q.Unit
		replyIngredients[i].Name = q.IngredientName
		if q.Quantity.Valid {
			replyIngredients[i].Quantity = float32(q.Quantity.Float64)
		}
	}
	return replyIngredients
}
