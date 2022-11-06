package handler

import (
	"net/http"

	"api/internal/logic"
	"api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetFakeRecipeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetFakeRecipeLogic(r.Context(), svcCtx)
		resp, err := l.GetFakeRecipe()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
