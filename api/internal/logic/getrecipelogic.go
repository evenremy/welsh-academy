package logic

import (
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
	quantities, err := l.svcCtx.QuantityModel.FindByRecipe(l.ctx, req.RecipeId)
	if err != nil {
		return nil, err
	}
	stages, err := l.svcCtx.StageModel.FindByRecipe(l.ctx, req.RecipeId)
	if err != nil {
		return nil, err
	}
}
