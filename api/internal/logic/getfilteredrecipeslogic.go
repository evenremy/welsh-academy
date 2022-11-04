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
	// todo: add your logic here and delete this line

	return
}
