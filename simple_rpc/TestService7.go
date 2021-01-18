package simple_rpc

import (
	"fmt"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/rpc"
	"github.com/duanhf2012/origin/service"
	"github.com/duanhf2012/origin/util/timer"
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
	slf.AfterFunc(time.Second*2,slf.RawTest)
	slf.AfterFunc(time.Second*2,slf.SyncTest)
	return nil
}

func (slf *TestService7) CallTest(t *timer.Timer){
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


func (slf *TestService7) AsyncCallTest(t *timer.Timer){
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

func (slf *TestService7) GoTest(t *timer.Timer){
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

type RawInputArgs struct {
	rawData       []byte
	additionParam []byte
}

func (args RawInputArgs) DoFree() {
}

func (args RawInputArgs) DoEscape() {

}

func (args RawInputArgs) GetRawData() []byte {
	return args.rawData
}

func (slf *TestService7) RawTest(t *timer.Timer){
	var inputArgs RawInputArgs
	inputArgs.rawData = []byte("hello world!")

	slf.RawGoNode(rpc.RpcProcessorGoGoPB, 1, 1, "TestService6", &inputArgs)
}

func (slf *TestService7) SyncTest(t *timer.Timer) {
	var input int = 3333
	slf.AsyncCall("TestService6.RPC_SyncTest",&input,func(output *int,err error){
		if err != nil {
			fmt.Printf("AsyncCall error :%+v\n",err)
		}else{
			fmt.Printf("AsyncCall output %d\n",*output)
		}
	})

	var output int = 444
	err := slf.Call("TestService6.RPC_SyncTest",&input,&output)
	fmt.Println(err,output)

}