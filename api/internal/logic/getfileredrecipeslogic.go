package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileredRecipesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileredRecipesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileredRecipesLogic {
	return &GetFileredRecipesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileredRecipesLogic) GetFileredRecipes(req *types.IngredientConstraintsReq) (resp *types.RecipesReply, err error) {
	// todo: add your logic here and delete this line

	return
}
