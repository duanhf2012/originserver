package simple_gogorpc

import (
	"github.com/duanhf2012/origin/log"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
	"originserver/common/gogoproto/gogorpc"
)

func init(){
	node.Setup(&TestService12{})
}

type TestService12 struct {
	service.Service
}

func (slf *TestService12) OnInit() error {
	return nil
}

func (slf *TestService12) PrintMsg() {
}

func (slf *TestService12) RPC_Service12TestOne(arg *gogorpc.TestOne, ret *gogorpc.TestOneRet) error {
	log.Release("RPC_Service12TestOne[%+v]", arg)
	ret.Msg = arg.Msg
	return nil
}
