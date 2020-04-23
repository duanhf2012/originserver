package simple_rpc

import (
	"fmt"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
	"time"
)

func init(){
	node.Setup(&TestService7{})
}

type TestService7 struct {
	service.Service
}

func (slf *TestService7) OnInit() error {
	slf.AfterFunc(time.Second*2,slf.CallTest)
	slf.AfterFunc(time.Second*2,slf.AsyncCallTest)
	slf.AfterFunc(time.Second*2,slf.GoTest)
	return nil
}

func (slf *TestService7) CallTest(){
	var input InputData
	input.A = 300
	input.B = 600
	var output int

	//同步调用其他服务的rpc,input为传入的rpc,output为输出参数
	err := slf.Call("TestService6.RPC_Sum",&input,&output)
	if err != nil {
		fmt.Printf("Call error :%+v\n",err)
	}else{
		fmt.Printf("Call output %d\n",output)
	}
}


func (slf *TestService7) AsyncCallTest(){
	var input InputData
	input.A = 300
	input.B = 600
	/*slf.AsyncCallNode(1,"TestService6.RPC_Sum",&input,func(output *int,err error){
	})*/
	//异步调用，在数据返回时，会回调传入函数
	//注意函数的第一个参数一定是RPC_Sum函数的第二个参数，err error为RPC_Sum返回值
	slf.AsyncCall("TestService6.RPC_Sum",&input,func(output *int,err error){
		if err != nil {
			fmt.Printf("AsyncCall error :%+v\n",err)
		}else{
			fmt.Printf("AsyncCall output %d\n",*output)
		}
	})
}

func (slf *TestService7) GoTest(){
	var input InputData
	input.A = 300
	input.B = 600

	//在某些应用场景下不需要数据返回可以使用Go，它是不阻塞的,只需要填入输入参数
	err := slf.Go("TestService6.RPC_Sum",&input)
	if err != nil {
		fmt.Printf("Go error :%+v\n",err)
	}

	//以下是广播方式，如果在同一个子网中有多个同名的服务名，CastGo将会广播给所有的node
	//slf.CastGo("TestService6.RPC_Sum",&input)
}
