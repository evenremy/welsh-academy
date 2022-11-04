package handler

import (
	"api/internal/config"
	"api/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
)

var testCtx *svc.ServiceContext

const TestConfigFilePath = "../../etc/welsh-academy-api.yaml"

func InitTestCtx() {
	if testCtx != nil {
		return
	}
	c := config.Config{}
	conf.MustLoad(TestConfigFilePath, &c)
	testCtx = svc.NewServiceContext(c)
}
