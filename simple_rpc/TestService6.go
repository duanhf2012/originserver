package simple_rpc

import (
	"encoding/json"
	"fmt"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/rpc"
	"github.com/duanhf2012/origin/service"
	"time"
)

func init(){
	node.Setup(&TestService6{})
}

type TestService6 struct {
	service.Service
}

func (slf *TestService6) OnInit() error {
	slf.RegRawRpc(1, &RPCRawTestCallBack{})
	return nil
}

type InputData struct {
	A int
	B int
}

func (slf *TestService6) RPC_Sum(input *InputData,output *int) error{
	*output = input.A+input.B
	return nil
}

func (slf *TestService6) RPC_SyncTest(resp rpc.Responder,input *int,out *int) error{
	go func(){
		time.Sleep(3*time.Second)
		var output int = *input
		resp(&output,rpc.NilError)
	}()

	return nil
}

type RPCRawTestCallBack struct {
}

func (cb *RPCRawTestCallBack) Unmarshal(data []byte) (interface{},error){
	fmt.Println(string(data))

	retData := InputData{}
	err := json.Unmarshal(data, &retData)
	return retData,err
}

func (cb *RPCRawTestCallBack) CB(data interface{}){
	fmt.Println(data)
}

