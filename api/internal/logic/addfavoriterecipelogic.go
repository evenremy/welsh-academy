package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFavoriteRecipeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFavoriteRecipeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFavoriteRecipeLogic {
	return &AddFavoriteRecipeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFavoriteRecipeLogic) AddFavoriteRecipe(req *types.FavReq) error {
	// todo: add your logic here and delete this line

	return nil
}
