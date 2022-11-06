package handler

import (
	"net/http"

	"api/internal/logic"
	"api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetFakeIngredientHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetFakeIngredientLogic(r.Context(), svcCtx)
		resp, err := l.GetFakeIngredient()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
