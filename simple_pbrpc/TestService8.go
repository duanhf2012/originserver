package simple_pbrpc

import (
	"github.com/duanhf2012/origin/log"
	"github.com/duanhf2012/origin/node"
	rpcHandle "github.com/duanhf2012/origin/rpc"
	"github.com/duanhf2012/origin/service"
	"github.com/duanhf2012/origin/util/timer"
	"github.com/duanhf2012/origin/util/uuid"
	"github.com/gogo/protobuf/proto"
	"math/rand"
	"originserver/common/proto/rpc"
	"time"
)

func init(){
	node.Setup(&TestService8{})
}

type RawInputArgs struct {
	rawData       []byte
}

func (args RawInputArgs) GetRawData() []byte {
	return args.rawData
}

func (args RawInputArgs) DoFree() {
}

func (args RawInputArgs) DoEscape() {
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

	//slf.AfterFunc(5 * time.Second, slf.TestGoParameter)
	//slf.AfterFunc(5 * time.Second, slf.TestCallError)
	slf.AfterFunc(3 * time.Second, slf.TestRpcResponder)
	slf.AfterFunc(5 * time.Second, slf.TestRpcRegister)
	//slf.AfterFunc(5 * time.Second, slf.TestCallPanic)
	slf.AfterFunc(5 * time.Second, slf.TestCallList)
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

func (slf *TestService8) TestCallList(t *timer.Timer) {
	arg := rpc.TestThree{
		UList:                make([]uint64, 0, 10),
	}
	arg.UList = append(arg.UList, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	slf.Go("TestService9.RPC_Service9TestSix", &arg)
}

func (slf *TestService8) TestRpcRegister(t *timer.Timer) {
	arg := rpc.TestOne{Msg: "test Rpc Register"}
	sendByte, _ := proto.Marshal(&arg)

	var inputArgs RawInputArgs
	inputArgs.rawData = sendByte
	slf.RawGoNode(rpcHandle.RpcProcessorGoGoPB, 3, 1, "TestService10", &inputArgs)

	slf.AfterFunc(5 * time.Second, slf.TestRpcRegister)
}

func (slf *TestService8) TestRpcResponder(t *timer.Timer) {
	argCall := rpc.TestTwo{
		Data:                 100,
		Msg:                  "test responder",
	}
	retCall := rpc.TestTwoRet{}
	errCall := slf.Call("TestService10.RPC_TestResponder", &argCall, &retCall)
	if errCall != nil {
		log.Error("%+v", errCall)
	}
	log.Release("call receive data[%+v]", &retCall)

	argAsyncCall := rpc.TestTwo{
		Data:                 200,
		Msg:                  "test responder AsyncCall",
	}
	errAsyncCall := slf.AsyncCall("TestService10.RPC_TestResponder", &argAsyncCall, func(ret *rpc.TestTwoRet, err error) {
		if err != nil {
			log.Error("%+v", err)
		}

		log.Release("asyncCall receive data[%+v]", ret)
	})
	if errAsyncCall != nil {
		log.Error("%+v", errAsyncCall)
	}

	slf.AfterFunc(3 * time.Second, slf.TestRpcResponder)
}

func (slf *TestService8) TestGoParameter(t *timer.Timer) {
	argOne := rpc.TestOne{Msg: "Test111111111111111111111"}
	errGo := slf.Go("TestService9.RPC_Service9TestThree", &argOne)
	if errGo != nil {
		log.Error("TestService8 RPC_Service9TestThree err[%+v], arg[%+v]", errGo, &argOne)
	}

	argTwo := rpc.TestOne{Msg: "Test22222222222"}
	errGo = slf.Go("TestService9.RPC_Service9TestThree", &argTwo)
	if errGo != nil {
		log.Error("TestService8 RPC_Service9TestThree err[%+v], arg[%+v]", errGo, &argTwo)
	}

	slf.AfterFunc(5 * time.Second, slf.TestGoParameter)
}

func (slf *TestService8) TestCallPanic() {
	argOne := rpc.TestOne{Msg: "Test111111111111111111111"}
	retOne := rpc.TestOneRet{}
	err := slf.Call("TestService9.RPC_Service9TestFive", &argOne, &retOne)
	if err != nil {
		log.Error("TestService8 RPC_Service9TestThree err[%+v], arg[%+v]", err, &argOne)
	}

	argTwo := rpc.TestOne{Msg: "Test22222222222"}
	err = slf.AsyncCall("TestService9.RPC_Service9TestFive", &argTwo, func(ret *rpc.TestOneRet, err error){
		if err != nil {
			log.Error("TestService8 RPC_Service9TestFive err[%+v], arg[%+v]", err, &argOne)
		}
	})
	if err != nil {
		log.Error("TestService8 RPC_Service9TestFive err[%+v], arg[%+v]", err, &argTwo)
	}
}

func (slf *TestService8) TestCallError(t *timer.Timer) {
	argOne := rpc.TestOne{Msg: "Test111111111111111111111"}
	retOne := rpc.TestOneRet{}
	err := slf.Call("TestService9.RPC_Service9TestFour", &argOne, &retOne)
	if err != nil {
		log.Error("TestService8 RPC_Service9TestThree err[%+v], arg[%+v]", err, &argOne)
	}

	argTwo := rpc.TestOne{Msg: "Test22222222222"}
	err1 := slf.AsyncCall("TestService9.RPC_Service9TestFour", &argTwo, func(ret *rpc.TestOneRet, err error){
		if err != nil {
			log.Error("TestService8 RPC_Service9TestFour err[%+v], arg[%+v]", err, &argTwo)
		}
	})
	if err1 != nil {
		log.Error("TestService8 RPC_Service9TestFour err[%+v], arg[%+v]", err1, &argTwo)
	}

	slf.AfterFunc(5 * time.Second, slf.TestCallError)
}

func (slf *TestService8) PrintMsg(t *timer.Timer) {
	slf.AfterFunc(5 * time.Second, slf.PrintMsg)
}

func (slf *TestService8) AsyncCallServer9TestOne(t *timer.Timer) {
	for i := 0; i < 10; i++ {
		go func() {
			arg := rpc.TestOne{Msg: uuid.Rand().HexEx()}
			errCall := slf.AsyncCall("TestService9.RPC_Service9TestOne",
				&arg, func(ret *rpc.TestOneRet, err error) {
				if err != nil || ret.Msg != arg.Msg {
					log.Error("TestService8 AsyncCallServer9TestOne err[%+v], arg[%+v], ret[%+v]", err, arg, ret)
					return
				}
				log.Release("Async call RPC_Service9TestOne receive[%+v]", ret)
			})
			if errCall != nil {
				log.Error("TestService8 AsyncCallServer9TestOne err[%+v]", errCall)
			}
		}()
	}
	slf.AfterFunc(10 * time.Second, slf.AsyncCallServer9TestOne)
}

func (slf *TestService8) AsyncCallServer9TestTwo(t *timer.Timer) {
	for i := 0; i < 10; i++ {
		go func() {
			arg := rpc.TestTwo{Msg: uuid.Rand().HexEx(), Data: int32(rand.Int())}
			errCall := slf.AsyncCall("TestService9.RPC_Service9TestTwo", &arg, func(ret *rpc.TestTwoRet, err error) {
				if err != nil || ret.Msg != arg.Msg || ret.Data != arg.Data {
					log.Error("TestService8 AsyncCallServer9TestTwo err[%+v], arg[%+v], ret[%+v]", err, arg, ret)
					return
				}
				log.Release("Async call RPC_Service9TestTwo receive[%+v]", ret)
			})
			if errCall != nil {
				log.Error("TestService8 AsyncCallServer9TestTwo err[%+v]", errCall)
			}
		}()
	}
	slf.AfterFunc(10 * time.Second, slf.AsyncCallServer9TestTwo)
}

func (slf *TestService8) CallServer9TestOne(t *timer.Timer) {
	for i := 0; i < 10; i++ {
		go func() {
			arg := rpc.TestOne{Msg: uuid.Rand().HexEx()}
			ret := rpc.TestOneRet{}
			errCall := slf.Call("TestService9.RPC_Service9TestOne", &arg, &ret)
			if errCall != nil || arg.Msg != ret.Msg {
				log.Error("TestService8 CallServer9TestOne err[%+v], arg[%+v], ret[%+v]", errCall, &arg, &ret)
				return
			}
			log.Release("call RPC_Service9TestOne receive[%+v]", ret)
		}()
	}
	slf.AfterFunc(5 * time.Second, slf.CallServer9TestOne)
}

func (slf *TestService8) CallServer9TestTwo(t *timer.Timer) {
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
	slf.AfterFunc(5 * time.Second, slf.CallServer9TestTwo)
}
