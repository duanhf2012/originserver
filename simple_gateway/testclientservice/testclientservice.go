package testclient

import (
	"fmt"
	"github.com/duanhf2012/origin/log"
	"github.com/duanhf2012/origin/network"
	"github.com/duanhf2012/origin/network/processor"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/rpc"
	"github.com/duanhf2012/origin/service"
	"github.com/duanhf2012/origin/util/timer"
	"github.com/golang/protobuf/proto"
	"time"
)

type TestClientService struct {
	service.Service
	tcpClient *network.TCPClient

	client   *Client
	sendCron *timer.Cron
}

func init() {
	node.Setup(&TestClientService{})
}

type Client struct {
	conn        *network.TCPConn
	pbProcessor *processor.PBProcessor
}

func (slf *Client) Run() {

	for {
		bytes, err := slf.conn.ReadMsg()
		if err != nil {
			slf.conn = nil
			log.Debug("read client is error:%+v", err)
			break
		}

		var msg rpc.PBRpcRequestData
		proto.Unmarshal(bytes, &msg)
		fmt.Print("read:", bytes, err, msg.GetSeq(), "\n")
	}

	//slf.conn.WriteMsg()
}

func (slf *Client) OnClose() {
}

func (slf *TestClientService) NewAgent(conn *network.TCPConn) network.Agent {
	slf.client = &Client{}
	slf.client.conn = conn
	slf.client.pbProcessor = processor.NewPBProcessor()

	return slf.client
}

func (slf *TestClientService) OnInit() error {
	slf.AfterFunc(time.Second*3, slf.OnTime)
	cron, _ := timer.NewCronExpr("0,10,20,30,40,50 * * * * *")
	slf.sendCron = slf.CronFunc(cron, slf.SendMessage)

	return nil
}

func (slf *TestClientService) OnTime() {
	slf.tcpClient = &network.TCPClient{}
	slf.tcpClient.Addr = "127.0.0.1:9830"
	slf.tcpClient.AutoReconnect = true
	slf.tcpClient.ConnNum = 1
	slf.tcpClient.NewAgent = slf.NewAgent
	slf.tcpClient.Start()
}

func (slf *TestClientService) SendMessage() {
	protoData := &rpc.PBRpcRequestData{}
	protoData.Seq = proto.Uint64(34)
	if slf.client == nil {
		return
	}
	pbPackInfo := slf.client.pbProcessor.MakeMsg(1001, protoData)
	bData, err := slf.client.pbProcessor.Marshal(pbPackInfo)
	if err == nil {
		err = slf.client.conn.WriteMsg(bData)
		if err != nil {
			slf.sendCron.Stop()
		}
	} else {
		fmt.Printf("error....")
	}
}
