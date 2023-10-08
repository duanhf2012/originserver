package simple_tcp

import (
	"fmt"
	"github.com/duanhf2012/origin/network"
	"github.com/duanhf2012/origin/network/processor"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
	"github.com/duanhf2012/origin/sysservice/tcpservice"
	"github.com/duanhf2012/origin/util/timer"
	"google.golang.org/protobuf/proto"
	msgpb "originserver/common/proto/msg"
	"time"
)

func init() {
	//因为与gateway中使用的TcpService不允许重复,所以这里使用自定义服务名称
	tcpService := &tcpservice.TcpService{}
	//设置服务的别名
	tcpService.SetName("MyTcpService")

	//安装MyTcpService服务
	node.Setup(tcpService)

	//安装TestTcpService服务
	node.Setup(&TestTcpService{})
}

// 新建自定义服务TestService1
type TestTcpService struct {
	service.Service
	processor  *processor.PBProcessor
	tcpService *tcpservice.TcpService

	client Client
}

func (slf *TestTcpService) OnInit() error {
	//1.先获取MyTcpService服务指针(它在init函数中被安装)。注意需要在cluster.json中将MyTcpService和TestTcpService一起配置。
	slf.tcpService = node.GetService("MyTcpService").(*tcpservice.TcpService)

	//2.新建内置的protobuf处理器，您也可以自定义处理器，参数返回的PBProcessor
	//该处理器负责解析消息，具体请阅读以下PBProcessor结构体实现代码。一般会根据自己的项目重新改写自己的Processor
	slf.processor = processor.NewPBProcessor()

	//3.向处理器注册监听客户连接断开事件
	slf.processor.RegisterDisConnected(slf.OnDisconnected)
	//4.向处理器注册监听客户连接事件
	slf.processor.RegisterConnected(slf.OnConnected)
	//5.向处理器注册监听消息类型MsgType_MsgReq，并注册回调
	slf.processor.Register(uint16(msgpb.MsgType_MsgReq), &msgpb.Req{}, slf.OnRequest)
	//6.向处理器将消息处理器设置到TcpService服务中，当收到网络消息时，MyTcpService会发送解析好的消息到Event管理中来。然后在当前MyTcpService服务中处理回调(例如：slf.OnRequest)。
	slf.tcpService.SetProcessor(slf.processor, slf.GetEventHandler())

	return nil
}

func (slf *TestTcpService) OnStart() {
	//为了演示，打开并发器，在其他协程中模拟客户端与服务器建立连接，并发送和接收消息。
	slf.OpenConcurrentByNumCPU(1)
	//在一个协程中运行testTcp函数
	slf.AsyncDo(slf.testTcp, nil)
}

type Client struct {
	network.TCPClient
	conn *network.TCPConn

	pbProcessor *processor.PBProcessor
}

func (slf *TestTcpService) testTcp() bool {
	//开启客户端连接
	slf.client.Addr = ":9930"          //因为在service.json中MyTcpService.ListenAddr配置的是0.0.0.0:9930
	slf.client.NewAgent = slf.newAgent //建议连接成功时会新建一个Agent，它即是slf.Client对象
	slf.client.pbProcessor = processor.NewPBProcessor()
	slf.client.pbProcessor.Register(uint16(msgpb.MsgType_MsgRes), &msgpb.Res{}, slf.client.MsgHandle) //注册消息回调
	slf.client.Start()

	//加一个定时器，定时发送消息
	slf.NewTicker(time.Second*1, func(ticker *timer.Ticker) {
		var req msgpb.Req
		req.Msg = "request"

		//生成一个pbInfo结构对象
		pbInfo := slf.client.pbProcessor.MakeMsg(uint16(msgpb.MsgType_MsgReq), &req)

		//将对象序列化成一个buf
		buf, _ := slf.client.pbProcessor.Marshal(0, pbInfo)

		//发送出去
		slf.client.Write(slf.client.conn, buf)
	})

	//返回值为true，表示继续执行回调，否则不执行回调
	return false
}

func (slf *TestTcpService) newAgent(c *network.TCPConn) network.Agent {
	slf.client.conn = c
	return &slf.client
}

// MsgHandle 消息处理，调用slf.pbProcessor.Register注册
func (slf *Client) MsgHandle(clientId uint64, msg proto.Message) {
	fmt.Println(msg)
}

func (slf *Client) Run() {
	for {
		//读取网络字节流，注意，在底层已经处理了粘包问题，是通过size+字节流的格式处理粘包。返回的内容是除size之外的后续字节流
		msgBuff, err := slf.conn.ReadMsg()
		if err != nil {
			break
		}

		//让解析器去解析，如何解析可以查看Unmarshal函数的实现。客户端可以模拟类似的方式去解析
		packInfo, err := slf.pbProcessor.Unmarshal(0, msgBuff)
		if err != nil {
			fmt.Println(err)
			continue
		}

		//直接路由回调
		slf.pbProcessor.MsgRoute(0, packInfo)
	}
}

func (slf *Client) OnClose() {
}

func (slf *TestTcpService) OnConnected(clientid uint64) {
	fmt.Printf("client id %d connected\n", clientid)
}

func (slf *TestTcpService) OnDisconnected(clientid uint64) {
	fmt.Printf("client id %d disconnected\n", clientid)
}

func (slf *TestTcpService) OnRequest(clientid uint64, msg proto.Message) {
	//解析客户端发过来的数据
	msgReq := msg.(*msgpb.Req)
	fmt.Println(msgReq)

	var res msgpb.Res
	res.Msg = "response"

	//发送数据给客户端
	slf.SendMsg(clientid, msgpb.MsgType_MsgRes, &res)
}

func (slf *TestTcpService) SendMsg(clientid uint64, msgtype msgpb.MsgType, msg proto.Message) error {
	//生成PBPackInfo结构
	msgData := slf.processor.MakeMsg(uint16(msgtype), msg)

	//将PBPackInfo结构发送至网络层，它会先在processor.PBProcessor进行序列化，具体序列化方式参照PBProcessor的Marshal函数
	return slf.tcpService.SendMsg(clientid, msgData)
}
