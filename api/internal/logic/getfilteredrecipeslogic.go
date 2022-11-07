package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFilteredRecipesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFilteredRecipesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFilteredRecipesLogic {
	return &GetFilteredRecipesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFilteredRecipesLogic) GetFilteredRecipes(req *types.IngredientConstraintsReq) (resp *types.RecipesReply, err error) {
	recipes, err := l.svcCtx.RecipeModel.FindFiltered(l.ctx, req.WithIngredientIdList, req.WithoutIngredientIdList)
	if err != nil {
		return nil, err
	}

	destRecipeList := make([]types.LiteRecipe, len(recipes))
	for i, r := range recipes {
		destRecipeList[i] = types.LiteRecipe(r)
	}
	return &types.RecipesReply{RecipeList: destRecipeList}, nil
}
