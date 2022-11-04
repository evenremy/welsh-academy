package handler

import (
	"api/internal/config"
	"api/internal/svc"
	"api/internal/types"
	"bytes"
	"encoding/json"
	"github.com/go-faker/faker/v4"
	_ "github.com/maxatome/go-testdeep/helpers/tdhttp"
	_ "github.com/maxatome/go-testdeep/td"
	"github.com/zeromicro/go-zero/core/conf"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

const configFilePath = "../../etc/welsh-academy-api.yaml"

var ctx *svc.ServiceContext

const TEST_METHOD string = "POST"
const TEST_TARGET string = "/ingredient"

func init() {
	c := config.Config{}
	conf.MustLoad(configFilePath, &c)
	ctx = svc.NewServiceContext(c)
}

// should work
func TestAddIngredient(t *testing.T) {
	dataReader := prepareFakeJson(&types.AddIngredientReq{})

	req := httptest.NewRequest(TEST_METHOD, TEST_TARGET, dataReader)
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	httpHandler := AddIngredientsHandler(ctx)
	httpHandler(rr, req)
	if rr.Code != http.StatusOK {
		t.Error(rr.Body)
	}
}

// Should fail
func TestAddRejectSameNameIngredient(t *testing.T) {
	dataReader := prepareFakeJson(&types.AddIngredientReq{})

	req1, rr1 := prepareNewRequestAndResponder(TEST_METHOD, TEST_TARGET, dataReader)
	testHandler := AddIngredientsHandler(ctx)
	testHandler(rr1, req1) // first try should work
	if rr1.Code != http.StatusOK {
		t.Error(rr1.Body)
	}

	req2, rr2 := prepareNewRequestAndResponder(TEST_METHOD, TEST_TARGET, dataReader)
	testHandler(rr2, req2) // second try should be rejected and explicitly
	if rr2.Code == http.StatusOK {
		t.Error(rr2.Body)
	}
	t.Log(rr2.Body)
	t.Log(rr2)
}

// Return a prepared Json request and response recorder ready to be passed to a http.HandlerFunc (testHandler)
func prepareNewRequestAndResponder(method string, target string, body *bytes.Reader) (*http.Request, *httptest.ResponseRecorder) {
	_, err := body.Seek(0, io.SeekStart)
	if err != nil {
		panic(err.Error())
	}

	req := httptest.NewRequest(method, target, body)
	req.Header.Add("Content-Type", "application/json")
	rr1 := httptest.NewRecorder()
	return req, rr1
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
