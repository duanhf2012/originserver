package simple_gogorpc

import (
	"github.com/duanhf2012/origin/log"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
	"github.com/duanhf2012/origin/util/timer"
	"github.com/duanhf2012/origin/util/uuid"
	"originserver/common/gogoproto/gogorpc"
	"time"
)

func init(){
	node.Setup(&TestService11{})
}

type TestService11 struct {
	service.Service
}

func (slf *TestService11) OnInit() error {
	//开始定时器
	slf.AfterFunc(10 * time.Second, slf.AsyncCallServer12TestOne)
	slf.AfterFunc(5 * time.Second, slf.CallServer12TestOne)
	return nil
}

func (slf *TestService11) AsyncCallServer12TestOne(t *timer.Timer) {
	for i := 0; i < 10; i++ {
		go func() {
			arg := gogorpc.TestOne{Msg: uuid.Rand().HexEx()}
			errCall := slf.AsyncCall("TestService12.RPC_Service12TestOne",
				&arg, func(ret *gogorpc.TestOneRet, err error) {
					if err != nil || ret.Msg != arg.Msg {
						log.Error("TestService11 AsyncCallServer12TestOne err[%+v], arg[%+v], ret[%+v]", err, arg, ret)
					}
					log.Release("Async call RPC_Service12TestOne receive[%+v]", ret)
				})
			if errCall != nil {
				log.Error("TestService11 AsyncCallServer12TestOne err[%+v]", errCall)
			}
		}()
	}
	slf.AfterFunc(10 * time.Second, slf.AsyncCallServer12TestOne)
}

func (slf *TestService11) CallServer12TestOne(t *timer.Timer) {
	for i := 0; i < 10; i++ {
		go func() {
			arg := gogorpc.TestOne{Msg: uuid.Rand().HexEx()}
			ret := gogorpc.TestOneRet{}
			errCall := slf.Call("TestService12.RPC_Service12TestOne", &arg, &ret)
			if errCall != nil || arg.Msg != ret.Msg {
				log.Error("TestService11 CallServer12TestOne err[%+v], arg[%+v], ret[%+v]", errCall, &arg, &ret)
			}
			log.Release("call RPC_Service12TestOne receive[%+v]", ret)
		}()
	}
	slf.AfterFunc(5 * time.Second, slf.CallServer12TestOne)
}