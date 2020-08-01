package simple_service

import (
	"fmt"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/rpc"
	"github.com/duanhf2012/origin/service"
	"github.com/duanhf2012/origin/util/timer"
	"github.com/golang/protobuf/proto"
	"time"
)

func init(){
	node.Setup(&TestService2{})
}

type TestService2 struct {
	service.Service
}

func (slf *TestService2) OnInit() error {
	fmt.Printf("TestService2 OnInit.\n")

	//间隔时间定时器
	slf.AfterFunc(time.Second*5,slf.OnSecondTick)

	//crontab模式定时触发
	//NewCronExpr的参数分别代表:Seconds Minutes Hours DayOfMonth Month DayOfWeek
	//以下为每换分钟时触发
	cron,_:=timer.NewCronExpr("0 * * * * *")
	slf.CronFunc(cron,slf.OnCron)
	return nil
}

func (slf *TestService2) OnSecondTick(){
	fmt.Printf("tick.\n")
	//slf.AfterFunc(time.Second*10,slf.OnSecondTick)




	now:= time.Now()
	/*
	for i:=1;i<=1000000;i++{
		var a InputData
		var b OutputData
		a.A = 3
		a.B = 8
		a.C = i
		slf.AsyncCallNode(1,"TestService1.RPC_Sum",&a,func(output *OutputData,err error){
			if output.C ==1000000 {
				fmt.Printf("xxxxxxxxxx:%d\n\n",time.Now().Sub(now).Microseconds())
			}
		})
*/
	for i:=1;i<=1;i++{

		var input rpc.PBRpcRequestData
		//var output rpc.PBRpcResponseData
		input.Seq = proto.Uint64(uint64(i))
		input.ServiceMethod = proto.String("aaaaa.bbbb.ccc")
		input.NoReply = proto.Bool(true)
		slf.AsyncCallNode(1,"TestService1.RPC_Test",&input,func(output *rpc.PBRpcResponseData,err error){
			if output.GetSeq() ==1000000 {
				fmt.Printf("xxxxxxxxxx:%d\n\n",time.Now().Sub(now).Microseconds())
			}
		})
		/*
		continue
		err := slf.CallNode(1,"TestService1.RPC_Sum",&a,&b)
		fmt.Print("CallNode:",err,a,b,"\n")

		err = slf.Go("TestService1.RPC_Sum",&a)
		fmt.Print("GO:",err,"\n")

		a.B = 10
		slf.AsyncCallNode(1,"TestService1.RPC_Sum",&a,func(output *OutputData,err error){
			fmt.Print("AsyncCallNode:",err,*output,"\n")
		})

		err = slf.GoNode(1,"TestService1.RPC_Sum",&a)
		fmt.Print("GoNode:",err,a,"\n")

		a.B = 30
		err = slf.Call("TestService1.RPC_Sum",&a,&b)
		fmt.Print("Call",err,a,b,"\n")*/
	}

	fmt.Printf("call :%d\n\n",time.Now().Sub(now).Microseconds())
}

func (slf *TestService2) OnCron(){
	fmt.Printf("A minute passed!\n")
}


func (slf *TestService2) OnRelease(){
	fmt.Print("OnRelease\n")
}