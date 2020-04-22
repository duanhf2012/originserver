package simple_module

import (
	"fmt"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
)

func init(){
	node.Setup(&TestService3{})
}

type TestService3 struct {
	service.Service
}

type Module1 struct {
	service.Module
}

type Module2 struct {
	service.Module
}

func (slf *Module1) OnInit()error{
	fmt.Printf("Module1 OnInit.\n")
	return nil
}

func (slf *Module1) OnRelease(){
	fmt.Printf("Module1 Release.\n")
}

func (slf *Module2) OnInit()error{
	fmt.Printf("Module2 OnInit.\n")
	return nil
}

func (slf *Module2) OnRelease(){
	fmt.Printf("Module2 Release.\n")
}


func (slf *TestService3) OnInit() error {
	//新建两个Module对象
	module1 := &Module1{}
	module2 := &Module2{}
	//将module1添加到服务中
	module1Id,_ := slf.AddModule(module1)
	//在module1中添加module2模块
	module1.AddModule(module2)
	fmt.Printf("module1 id is %d, module2 id is %d\n",module1Id,module2.GetModuleId())

	//释放模块module1
	slf.ReleaseModule(module1Id)
	fmt.Printf("xxxxxxxxxxx")
	return nil
}


