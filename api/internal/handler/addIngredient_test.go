package handler

import (
	"api/internal/config"
	"api/internal/svc"
	"api/internal/types"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/conf"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var configFile = "../../etc/welsh-academy-api.yaml"

func TestAddIngredient(t *testing.T) {
	var c config.Config
	conf.MustLoad(configFile, &c)

	jsonData, _ := json.Marshal(types.AddIngredientReq{Name: "hello world"})

	req := httptest.NewRequest("POST", "/ingredient", strings.NewReader(string(jsonData)))
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	ctx := svc.NewServiceContext(c)
	httpHandler := addIngredientsHandler(ctx)
	httpHandler(rr, req)
	if rr.Code != http.StatusOK {
		t.Error(rr.Body)
	}
}
