package simple_service

import (
	"fmt"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
	"github.com/duanhf2012/origin/util/timer"
	"time"
)

//模块加载时自动安装TestService1服务
func init(){
	node.Setup(&TestService1{})
}

//新建自定义服务TestService1
type TestService1 struct {

	//所有的自定义服务必需加入service.Service基服务
	//那么该自定义服务将有各种功能特性
	//例如: Rpc,事件驱动,定时器等
	service.Service

	crontabModuleId int64
}

type CrontabModule struct {
	service.Module
}


func (slf *CrontabModule) OnInit()error {
  //cron定时器使用
  pCron,err := timer.NewCronExpr("* * * * * *")
  if err != nil {
  	return err
  }

  //开始定时器
  slf.CronFunc(pCron,slf.OnRun)
  return nil
}

func (slf *CrontabModule) OnRun(){
	fmt.Printf("CrontabModule OnRun.\n")
}

//服务初始化函数，在安装服务时，服务将自动调用OnInit函数
func (slf *TestService1) OnInit() error {
	fmt.Printf("TestService1 OnInit.\n")
	//打开性能分析工具
	slf.OpenProfiler()
	//监控超过1秒的慢处理
	slf.GetProfiler().SetOverTime(time.Second*1)
	//监控超过10秒的超慢处理，您可以用它来定位是否存在死循环
	//比如以下设置10秒，我的应用中是不会发生超过10秒的一次函数调用
	//所以设置为10秒。
	slf.GetProfiler().SetMaxOverTime(time.Second*10)

	slf.AfterFunc(time.Second*2,slf.Loop)
	//打开多线程处理模式，10个协程并发处理
	//slf.SetGoRouterNum(10)

	//增加module，在module中演示定时器
	var err error
	slf.crontabModuleId,err = slf.AddModule(&CrontabModule{})
	if err!= nil {
		return err
	}

	//10秒后删除module
	slf.AfterFunc(time.Second*10,slf.ReleaseCrontabModule)
	return nil
}

func (slf *TestService1) ReleaseCrontabModule(){
	//释放module后，定时器也会一起释放
	slf.ReleaseModule(slf.crontabModuleId)
}

func (slf *TestService1) Loop(){
	//for {
		time.Sleep(time.Second*1)
	//}
}

