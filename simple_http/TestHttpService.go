package simple_http

import (
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
	"github.com/duanhf2012/origin/sysservice"
)

func init(){
	node.Setup(&sysservice.HttpService{})
	node.Setup(&TestHttpService{})
}

//新建自定义服务TestService1
type TestHttpService struct {
	service.Service
}

func (slf *TestHttpService) OnInit() error {
	//
	httpervice := node.GetService("HttpService").(*sysservice.HttpService)
	httpRouter := sysservice.NewHttpHttpRouter()
	httpervice.SetHttpRouter(httpRouter,slf.GetEventHandler())

	httpRouter.GET("/get/query", slf.HttpGet)
	httpRouter.POST("/post/query", slf.HttpPost)
	httpRouter.SetServeFile(sysservice.METHOD_GET,"/img/head/","d:/img")
	return nil
}

func (slf *TestHttpService) HttpGet(session *sysservice.HttpSession){

}

func (slf *TestHttpService) HttpPost(session *sysservice.HttpSession){

}
