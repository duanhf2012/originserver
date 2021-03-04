package main

import (
	"github.com/duanhf2012/origin/node"
	"time"

	_ "originserver/simple_event"
	_ "originserver/simple_gogorpc"
	_ "originserver/simple_http"
	_ "originserver/simple_module"
	_ "originserver/simple_pbrpc"
	_ "originserver/simple_rpc"
	//导入simple_service模块
	_ "originserver/simple_service"
	_ "originserver/simple_tcp"
)


func main() {
	//打开性能分析报告功能，并设置10秒汇报一次
	node.OpenProfilerReport(time.Second * 10)
	node.Start()
}
