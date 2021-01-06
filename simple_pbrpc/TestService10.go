package simple_pbrpc

import (
	"github.com/duanhf2012/origin/log"
	"github.com/duanhf2012/origin/node"
	rpcHandle "github.com/duanhf2012/origin/rpc"
	"github.com/duanhf2012/origin/service"
	"github.com/golang/protobuf/proto"
	"originserver/common/proto/rpc"
)

func init(){
	node.Setup(&TestService10{})
}

type TestRequest struct {
	request 	*rpc.TestTwo
	responder 	rpcHandle.Responder
}

type TestService10 struct {
	service.Service
	channelOptData chan TestRequest
}

func (slf *TestService10) OnInit() error {
	slf.channelOptData = make(chan TestRequest, 50)
	go slf.ExecuteOptData(slf.channelOptData)

	slf.RegRawRpc(1, slf.TestRpcRegister)
	return nil
}

func (slf *TestService10) ExecuteOptData(channelOptData chan TestRequest) {
	for {
		select {
		case optData := <-channelOptData:
			slf.DoDealData(optData)
		}
	}
}

func (slf *TestService10) DoDealData(dataReq TestRequest) {
	retInfo := rpc.TestTwoRet{
		Data:                 dataReq.request.Data,
		Msg:                  dataReq.request.Msg,
	}
	dataReq.responder(&retInfo, rpcHandle.RpcError("test err"))
}

func (slf *TestService10) RPC_TestResponder(responder rpcHandle.Responder, request *rpc.TestTwo) error{
	var tRequest TestRequest
	tRequest.request = request
	tRequest.responder = responder

	slf.channelOptData <- tRequest
	return nil
}

func (slf *TestService10) TestRpcRegister(byteBuffer []byte) {
	argInfo := rpc.TestOne{}
	proto.UnmarshalMerge(byteBuffer, &argInfo)
	log.Release("TestRpcRegister receive[%+v]", &argInfo)

	return
}
