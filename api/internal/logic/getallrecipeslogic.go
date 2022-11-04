package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllRecipesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllRecipesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllRecipesLogic {
	return &GetAllRecipesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllRecipesLogic) GetAllRecipes() (resp *types.RecipesReply, err error) {
	// todo: add your logic here and delete this line

	return
}
