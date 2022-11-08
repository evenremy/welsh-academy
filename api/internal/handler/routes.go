// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ingredients",
				Handler: AllIngredientsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/ingredient",
				Handler: AddIngredientsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/recipes",
				Handler: GetAllRecipesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/filtered_recipes",
				Handler: GetFilteredRecipesHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/recipe/:recipe_id",
				Handler: GetRecipeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/recipe",
				Handler: AddRecipeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/favorite_recipe/:recipe_id",
				Handler: AddFavoriteRecipeHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/favorite_recipe/:recipe_id",
				Handler: DeleteFavoriteRecipeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/favorite_recipes/:user_id",
				Handler: GetFavoriteRecipesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user",
				Handler: AddUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/example_recipe",
				Handler: GetFakeRecipeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/example_ingredient",
				Handler: GetFakeIngredientHandler(serverCtx),
			},
		},
	)
}
