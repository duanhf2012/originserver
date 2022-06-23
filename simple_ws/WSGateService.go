package simple_ws

import (
	"github.com/duanhf2012/origin/log"
	"github.com/duanhf2012/origin/network/processor"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
	"github.com/duanhf2012/origin/sysservice/wsservice"
)

type WsGateService struct {
	service.Service
	processor *processor.JsonProcessor
	wsService *wsservice.WSService
}

func init() {
	//因为与gateway中使用的TcpService不允许重复,所以这里使用自定义服务名称
	node.Setup(&wsservice.WSService{})
	node.Setup(&WsGateService{})
}

type MessageType = uint16

const (
	MessageId1 MessageType = 1
)

func (slf *WsGateService) OnInit() error {
	//获取安装好了的TcpService对象
	slf.wsService = node.GetService("WSService").(*wsservice.WSService)

	slf.processor = processor.NewJsonProcessor()
	//注册监听客户连接断开事件
	slf.processor.RegisterDisConnected(slf.OnDisconnected)
	//注册监听客户连接事件
	slf.processor.RegisterConnected(slf.OnConnected)

	//测试用的json：{"typ":1,"UserName":"username...张","Passwd":"ksdfjwef8"}
	slf.processor.Register(MessageId1, &MsgStruct{}, slf.ProcessMessage)
	slf.wsService.SetProcessor(slf.processor, slf.GetEventHandler())
	//默认消息类型是：websocket.TextMessage
	//slf.wsService.SetMessageType(websocket.BinaryMessage)
	return nil
}

type MsgStruct struct {
	UserName string
	Passwd   string
}

func (slf *WsGateService) ProcessMessage(clientid uint64, msg interface{}) {
	//解析客户端发过来的数据
	msgStruct := msg.(*MsgStruct)

	log.Debug("recv:%+v", msgStruct)

	//发送数据给客户端
	err := slf.wsService.SendMsg(clientid, msgStruct)
	if err != nil {
		log.Warning("send msg is fail %+v!", err)
	}
}

func (slf *WsGateService) OnConnected(clientid uint64) {
	log.SDebug("client id ", clientid, " is connected")
}

func (slf *WsGateService) OnDisconnected(clientid uint64) {
	log.SDebug("client id ", clientid, " is disconnected\n")
}
