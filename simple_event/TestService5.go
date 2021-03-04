package simple_event

import (
	"fmt"
	"github.com/duanhf2012/origin/event"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
)

func init(){
	node.Setup(&TestService5{})
}

type TestService5 struct {
	service.Service
}

type TestModule struct {
	service.Module
}

func (slf *TestModule) OnInit() error{
	//在当前node中查找TestService4
	pService := node.GetService("TestService4")

	//在TestModule中，往TestService4中注册EVENT1类型事件监听
	pService.(*TestService4).GetEventProcessor().RegEventReciverFunc(EVENT1,slf.GetEventHandler(),slf.OnModuleEvent)
	return nil
}

func (slf *TestModule) OnModuleEvent(ev event.IEvent){
	event := ev.(*event.Event)
	fmt.Printf("OnModuleEvent type :%d data:%+v\n",event.GetEventType(),event.Data)
}


//服务初始化函数，在安装服务时，服务将自动调用OnInit函数
func (slf *TestService5) OnInit() error {
	//通过服务名获取服务对象
	pService := node.GetService("TestService4")

	////在TestModule中，往TestService4中注册EVENT1类型事件监听
	pService.(*TestService4).GetEventProcessor().RegEventReciverFunc(EVENT1,slf.GetEventHandler(),slf.OnServiceEvent)
	slf.AddModule(&TestModule{})
	return nil
}

func (slf *TestService5) OnServiceEvent(ev event.IEvent){
	event := ev.(*event.Event)
	fmt.Printf("OnServiceEvent type :%d data:%+v\n",event.Type,event.Data)
}

