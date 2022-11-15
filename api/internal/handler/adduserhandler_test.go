package handler

import (
	"api/internal/types"
	"context"
	"github.com/go-faker/faker/v4"
	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"github.com/maxatome/go-testdeep/td"
	"net/http"
	"testing"
)

func TestAddUserHandlerWorking(t *testing.T) {
	newUser := types.AddUserReq{}
	newUser.Username = faker.Username()

	testApi := tdhttp.NewTestAPI(t, AddUserHandler(testCtx))
	var userId int64
	testApi.
		Name("Add user").
		PostJSON("/user", &newUser).
		CmpStatus(http.StatusOK).
		CmpJSONBody(td.JSON(`
{
	"userId": $1
}
`, td.Catch(&userId, td.Gt(0))))
	defer func() {
		_ = testCtx.UserModel.Delete(context.Background(), userId)
	}()

	testUser, err := testCtx.UserModel.FindOne(context.Background(), userId)
	td.CmpNoError(t, err)
	td.EqDeeply(testUser.Id, userId)
}

func TestAddUserHandlerEmpty(t *testing.T) {
	newUser := types.AddUserReq{}
	newUser.Username = ""

	testApi := tdhttp.NewTestAPI(t, AddUserHandler(testCtx))
	testApi.
		Name("Add user").
		PostJSON("/user", &newUser).
		CmpStatus(http.StatusBadRequest).
		CmpJSONBody(td.JSON(`
{
	"code" : 10,
	"msg" : "Bad username"
}
`))
	_, err := testCtx.UserModel.FindOneByUsername(context.Background(), "")
	td.CmpError(t, err)
}
