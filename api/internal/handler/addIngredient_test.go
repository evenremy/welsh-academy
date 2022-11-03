package handler

import (
	"api/internal/config"
	"api/internal/svc"
	"api/internal/types"
	"bytes"
	"encoding/json"
	"github.com/go-faker/faker/v4"
	_ "github.com/go-faker/faker/v4"
	"github.com/zeromicro/go-zero/core/conf"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var configFilePath = "../../etc/welsh-academy-api.yaml"
var ctx *svc.ServiceContext

func init() {
	c := config.Config{}
	conf.MustLoad(configFilePath, &c)
	ctx = svc.NewServiceContext(c)
}

// should work
func TestAddIngredient(t *testing.T) {
	dataReader := prepareFakeJson(&types.AddIngredientReq{})

	req := httptest.NewRequest("POST", "/ingredient", dataReader)
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	httpHandler := AddIngredientsHandler(ctx)
	httpHandler(rr, req)
	if rr.Code != http.StatusOK {
		t.Error(rr.Body)
	}
}

// TODO update test
// Should fail
func TestAddRejectSameNameIngredient(t *testing.T) {
	dataReader := prepareFakeJson(&types.AddIngredientReq{})

	req := httptest.NewRequest("POST", "/ingredient", dataReader)
	req.Header.Add("Content-Type", "application/json")
	rr1 := httptest.NewRecorder()
	httpHandler := AddIngredientsHandler(ctx)

	httpHandler(rr1, req) // first try should work
	if rr1.Code != http.StatusOK {
		t.Error(rr1.Body)
	}

	_, err := dataReader.Seek(0, io.SeekStart)
	if err != nil {
		t.Error(err)
	}
	rr2 := httptest.NewRecorder()
	httpHandler(rr2, req) // second try should be rejected and explicitly
	if rr2.Code == http.StatusOK {
		t.Error(rr2.Body)
	}

}

func prepareFakeJson(emptyStruct *types.AddIngredientReq) *bytes.Reader {
	err := faker.FakeData(emptyStruct)
	if err != nil {
		panic(err.Error())
	}

	jsonData, err := json.Marshal(emptyStruct)
	if err != nil {
		panic(err.Error())
	}
	return bytes.NewReader(jsonData)
}

// TODO test missing/empty values
// TODO values should be secured
// TODO values should be striped
