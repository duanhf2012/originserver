package simple_pbrpc

import (
	"github.com/duanhf2012/origin/log"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
	"github.com/duanhf2012/origin/util/uuid"
	"math/rand"
	"originserver/common/proto/rpc"
	"time"
)

func init(){
	node.Setup(&TestService8{})
}

type TestService8 struct {
	service.Service
}

func (slf *TestService8) OnInit() error {
	//开始定时器
	slf.AfterFunc(10 * time.Second, slf.AsyncCallServer9TestOne)
	slf.AfterFunc(10 * time.Second, slf.AsyncCallServer9TestTwo)
	slf.AfterFunc(5 * time.Second, slf.CallServer9TestOne)
	slf.AfterFunc(5 * time.Second, slf.CallServer9TestTwo)
	slf.AfterFunc(5 * time.Second, slf.PrintMsg)
	return nil
}

func (slf *TestService8) RPC_Service8TestOne(arg *rpc.TestOne, ret *rpc.TestOneRet) error {
	log.Release("RPC_Service8TestOne[%+v]", arg)
	ret.Msg = arg.Msg
	return nil
}

func (slf *TestService8) RPC_Service8TestTwo(arg *rpc.TestTwo, ret *rpc.TestTwoRet) error {
	log.Release("RPC_Service8TestTwo[%+v]", arg)
	ret.Msg = arg.Msg
	ret.Data = arg.Data
	return nil
}

func (slf *TestService8) PrintMsg() {
	slf.AfterFunc(5 * time.Second, slf.PrintMsg)
}

func (slf *TestService8) AsyncCallServer9TestOne() {
	for i := 0; i < 10; i++ {
		go func() {
			arg := rpc.TestOne{Msg: uuid.Rand().HexEx()}
			errCall := slf.AsyncCall("TestService9.RPC_Service9TestOne",
				&arg, func(ret *rpc.TestOneRet, err error) {
				if err != nil || ret.Msg != arg.Msg {
					log.Error("TestService8 AsyncCallServer9TestOne err[%+v], arg[%+v], ret[%+v]", err, arg, ret)
				}
				//log.Release("Async call RPC_Service9TestOne receive[%+v]", ret)
			})
			if errCall != nil {
				log.Error("TestService8 AsyncCallServer9TestOne err[%+v]", errCall)
			}
		}()
	}
	slf.AfterFunc(10 * time.Second, slf.AsyncCallServer9TestOne)
}

func (slf *TestService8) AsyncCallServer9TestTwo() {
	for i := 0; i < 10; i++ {
		go func() {
			arg := rpc.TestTwo{Msg: uuid.Rand().HexEx(), Data: int32(rand.Int())}
			errCall := slf.AsyncCall("TestService9.RPC_Service9TestTwo", &arg, func(ret *rpc.TestTwoRet, err error) {
				if err != nil || ret.Msg != arg.Msg || ret.Data != arg.Data {
					log.Error("TestService8 AsyncCallServer9TestTwo err[%+v], arg[%+v], ret[%+v]", err, arg, ret)
				}
				//log.Release("Async call RPC_Service9TestTwo receive[%+v]", ret)
			})
			if errCall != nil {
				log.Error("TestService8 AsyncCallServer9TestTwo err[%+v]", errCall)
			}
		}()
	}
	slf.AfterFunc(10 * time.Second, slf.AsyncCallServer9TestTwo)
}

func (slf *TestService8) CallServer9TestOne() {
	for i := 0; i < 10; i++ {
		go func() {
			arg := rpc.TestOne{Msg: uuid.Rand().HexEx()}
			ret := rpc.TestOneRet{}
			errCall := slf.Call("TestService9.RPC_Service9TestOne", &arg, &ret)
			if errCall != nil || arg.Msg != ret.Msg {
				log.Error("TestService8 CallServer9TestOne err[%+v], arg[%+v], ret[%+v]", errCall, &arg, &ret)
			}
			//log.Release("call RPC_Service9TestOne receive[%+v]", ret)
		}()
	}
}

func (slf *TestService8) CallServer9TestTwo() {
	for i := 0; i < 10; i++ {
		go func() {
			arg := rpc.TestTwo{Msg: uuid.Rand().HexEx(), Data: int32(rand.Int())}
			ret := rpc.TestTwoRet{}
			errCall := slf.Call("TestService9.RPC_Service9TestTwo", &arg, &ret)
			if errCall != nil || ret.Msg != arg.Msg || ret.Data != arg.Data {
				log.Error("TestService8 CallServer9TestTwo err[%+v], arg[%+v], ret[%+v]", errCall, &arg, &ret)
			}
			//log.Release("call RPC_Service9TestTwo receive[%+v]", ret)
		}()
	}
}
