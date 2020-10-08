package gameservice

import (
	"fmt"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/rpc"
	"github.com/duanhf2012/origin/service"
	"github.com/duanhf2012/origin/sysservice/tcpgateway"
	"github.com/golang/protobuf/proto"
	"net"
)

type GameService struct {
	service.Service
	tcpConn         *net.TCPConn
	gateProxyModule *tcpgateway.GateProxyModule
}

func init() {
	node.Setup(&GameService{})
}

func (slf *GameService) OnInit() error {
	slf.gateProxyModule = tcpgateway.NewGateProxyModule()
	slf.AddModule(slf.gateProxyModule)
	return nil
}

func GetClientId(addition rpc.IRawAdditionParam) uint64 {
	return addition.GetParamValue().(uint64)
}

func (slf *GameService) RPC_Login(addition rpc.IRawAdditionParam, input *rpc.PBRpcRequestData) error {
	clientId := addition.GetParamValue().(uint64)

	input.Seq = proto.Uint64(311111)
	slf.gateProxyModule.Send(clientId, 1000, input)

	return nil
}

func (slf *GameService) RPC_OnConnect(addition rpc.IRawAdditionParam, input *tcpgateway.PlaceHolders) error {
	clientId := addition.GetParamValue().(uint64)
	fmt.Print("client ", clientId, " onconnect...\n")
	return nil
}

func (slf *GameService) RPC_OnDisConnect(addition rpc.IRawAdditionParam, input *tcpgateway.PlaceHolders) error {
	clientId := addition.GetParamValue().(uint64)
	fmt.Print("client ", clientId, " ondisconnect...\n")

	return nil
}
