package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFavoriteRecipeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFavoriteRecipeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFavoriteRecipeLogic {
	return &DeleteFavoriteRecipeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFavoriteRecipeLogic) DeleteFavoriteRecipe(req *types.FavReq) error {
	err := l.svcCtx.FavoriteModel.DeleteFavoriteByRecipeAndUserId(l.ctx, req.RecipeId, req.UserId)
	if err != nil {
		return err
	}

	return nil
}
