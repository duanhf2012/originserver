package simple_rpc

import (
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
)

func init(){
	node.Setup(&TestService6{})
}

type TestService6 struct {
	service.Service
}

func (slf *TestService6) OnInit() error {
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
