package alarm

import (
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/smartystreets/goconvey/convey"

	"go-common/app/service/main/resource/conf"
	"go-common/app/service/main/resource/model"

	"gopkg.in/h2non/gock.v1"
)

var (
	d *Dao
)

func TestMain(m *testing.M) {
	if os.Getenv("DEPLOY_ENV") != "" {
		flag.Set("app_id", "main.app-svr.resource-service")
		flag.Set("conf_token", "y79sErNhxggjvULS0O8Czas9PaxHBF5o")
		flag.Set("tree_id", "3232")
		flag.Set("conf_version", "docker-1")
		flag.Set("deploy_env", "uat")
		flag.Set("conf_host", "config.bilibili.co")
		flag.Set("conf_path", "/tmp")
		flag.Set("region", "sh")
		flag.Set("zone", "sh001")
	} else {
		flag.Set("conf", "../../cmd/resource-service-test.toml")
	}
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	d = New(conf.Conf)
	d.httpClient.SetTransport(gock.DefaultTransport)
	m.Run()
	os.Exit(0)
}

func httpMock(method, url string) *gock.Request {
	r := gock.New(url)
	r.Method = strings.ToUpper(method)
	return r
}

func TestDaoCheckURL(t *testing.T) {
	var (
		originURL = "https://www.bilibili.com"
		wis       = []*model.ResWarnInfo{}
	)
	convey.Convey("CheckURL", t, func(ctx convey.C) {
		httpMock("GET", "http://www.bilibili.com").Reply(200)
		d.CheckURL(originURL, wis)
		ctx.Convey("no return values", func() {

		})
	})
}
