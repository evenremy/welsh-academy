package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoriteRecipesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFavoriteRecipesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteRecipesLogic {
	return &GetFavoriteRecipesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFavoriteRecipesLogic) GetFavoriteRecipes(req *types.FavReq) (resp *types.RecipesReply, err error) {
	// todo: add your logic here and delete this line

	return
}
