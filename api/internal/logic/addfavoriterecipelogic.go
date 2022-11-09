package logic

import (
	"api/errorx"
	"api/model"
	"api/model/favorite"
	"context"
	"net/http"

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
	// Check user presence
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	switch err {
	case nil:
	case model.ErrNotFound:
		return errorx.NewCodeError(22, "unable to add favorite, user not found", http.StatusForbidden)
	}

	_, err = l.svcCtx.FavoriteModel.InsertReturningId(l.ctx, &favorite.Favorites{
		User:   req.UserId,
		Recipe: req.RecipeId,
	})
	if err != nil {
		l.Error(err)
		return errorx.NewCodeError(23, "Saving favorite failed", http.StatusNotModified)
	}
	return nil
}
