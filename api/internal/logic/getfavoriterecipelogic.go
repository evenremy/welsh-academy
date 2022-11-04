package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoriteRecipeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFavoriteRecipeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteRecipeLogic {
	return &GetFavoriteRecipeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFavoriteRecipeLogic) GetFavoriteRecipe(req *types.AuthReq) error {
	// todo: add your logic here and delete this line

	return nil
}
