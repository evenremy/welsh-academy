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
	favRecipeList, err := l.svcCtx.FavoriteModel.FindFavRecipeByUser(l.ctx, req.UserId)

	reply := types.RecipesReply{}
	reply.RecipeList = make([]types.LiteRecipe, len(favRecipeList))
	for i, r := range favRecipeList {
		reply.RecipeList[i] = types.LiteRecipe(struct {
			Id    int64
			Title string
		}{Id: r.Id, Title: r.Title})
	}
	return &reply, nil
}
