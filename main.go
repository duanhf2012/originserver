
package main

import (
	"github.com/duanhf2012/origin/node"
	"time"

	_ "orginserver/simple_module"
	//导入simple_service模块
	_ "orginserver/simple_service"
)

func main(){
	//打开性能分析报告功能，并设置10秒汇报一次
	node.OpenProfilerReport(time.Second*10)
	node.Start()
}
