package simple_tcp

import (
	"fmt"
	"github.com/duanhf2012/origin/network/processor"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
	"github.com/duanhf2012/origin/sysservice"
	"github.com/golang/protobuf/proto"
	"originserver/simple_tcp/msgpb"
)

func init(){
	node.Setup(&sysservice.TcpService{})
	node.Setup(&TestTcpService{})
}

//新建自定义服务TestService1
type TestTcpService struct {
	service.Service
	processor *processor.PBProcessor
	tcpService *sysservice.TcpService
}

func (slf *TestTcpService) OnInit() error {
	//获取安装好了的TcpService对象
	slf.tcpService =  node.GetService("TcpService").(*sysservice.TcpService)

	//新建内置的protobuf处理器，您也可以自定义路由器，比如json，后续会补充
	slf.processor = processor.NewPBProcessor()

	//注册监听客户连接断开事件
	slf.processor.RegisterDisConnected(slf.OnDisconnected)
	//注册监听客户连接事件
	slf.processor.RegisterConnected(slf.OnConnected)
	//注册监听消息类型MsgType_MsgReq，并注册回调
	slf.processor.Register(uint16(msgpb.MsgType_MsgReq),&msgpb.Req{},slf.OnRequest)
	//将protobuf消息处理器设置到TcpService服务中
	slf.tcpService.SetProcessor(slf.processor,slf.GetEventHandler())

	return nil
}


func (slf *TestTcpService) OnConnected(clientid uint64){
	fmt.Printf("client id %d connected\n",clientid)
}


func (slf *TestTcpService) OnDisconnected(clientid uint64){
	fmt.Printf("client id %d disconnected\n",clientid)
}

func (slf *TestTcpService) OnRequest (clientid uint64,msg proto.Message){
	//解析客户端发过来的数据
	pReq := msg.(*msgpb.Req)
	//发送数据给客户端
	err := slf.tcpService.SendMsg(clientid,&msgpb.Req{
		Msg: proto.String(pReq.GetMsg()),
	})
	if err != nil {
		fmt.Printf("send msg is fail %+v!",err)
	}
}