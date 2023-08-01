package simple_rpc

import (
	"encoding/json"
	"fmt"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/rpc"
	"github.com/duanhf2012/origin/service"
	"time"
)

func init() {
	node.Setup(&TestService6{})
}

type TestService6 struct {
	service.Service
}

func (slf *TestService6) OnInit() error {
	slf.RegRawRpc(1, slf.RawRpcCallBack)

	//监听其他Node结点连接和断开事件
	slf.RegRpcListener(slf)
	return nil
}

type InputData struct {
	A int
	B int
}

func (slf *TestService6) OnNodeConnected(nodeId int) {
	fmt.Printf("node id %d is conntected.\n", nodeId)
}

func (slf *TestService6) OnNodeDisconnect(nodeId int) {
	fmt.Printf("node id %d is disconntected.\n", nodeId)
}

func (slf *TestService6) RPC_Sum(input *InputData, output *int) error {
	*output = input.A + input.B
	//等待1.5s
	time.Sleep(1500 * time.Millisecond)
	return nil
}

func (slf *TestService6) RPC_SyncTest(resp rpc.Responder, input *int, out *int) error {
	go func() {
		time.Sleep(3 * time.Second)
		var output int = *input
		resp(&output, rpc.NilError)
	}()

	return nil
}

func (slf *TestService6) RawRpcCallBack(data []byte) {
	retData := InputData{}
	err := json.Unmarshal(data, &retData)
	fmt.Println(err, retData)
}
