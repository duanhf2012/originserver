package simple_asyncdo

import (
	"fmt"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
	"github.com/duanhf2012/origin/util/timer"
	"time"
)

// 模块加载时自动安装TestService1服务
func init() {
	node.Setup(&TestService13{})
}

// 新建自定义服务TestService1
type TestService13 struct {
	service.Service
}

func (slf *TestService13) OnInit() error {
	//以下通过cpu数量来定开启协程并发数量，建议:(1)cpu密集型计算使用1.0  (2)i/o密集型使用2.0或者更高
	slf.OpenConcurrentByNumCPU(1.0)

	//以下通过函数打开并发协程数，以下协程数最小5，最大10，任务管道的cap数量1000000
	//origin会根据任务的数量在最小与最大协程数间动态伸缩
	//slf.OpenConcurrent(5, 10, 1000000)
	return nil
}

func (slf *TestService13) OnStart() {
	slf.AfterFunc(time.Second*5, func(*timer.Timer) {
		slf.testAsyncDo()
	})
}

func (slf *TestService13) testAsyncDo() {
	var context struct {
		data int64
	}

	//1.示例普通使用
	//参数一的函数在其他协程池中执行完成，将执行完成事件放入服务工作协程，
	//参数二的函数在服务协程中执行，是协程安全的。
	slf.AsyncDo(func() bool {
		//该函数回调在协程池中执行
		context.data = 100
		return true
	}, func(err error) {
		//函数将在服务协程中执行
		fmt.Print(context.data) //显示100
	})

	//2.示例按队列顺序
	//参数一传入队列Id,同一个队列Id将在协程池中被排队执行
	//以下进行两次调用，因为两次都传入参数queueId都为1，所以它们会都进入queueId为1的排队执行
	queueId := int64(1)
	for i := 0; i < 2; i++ {
		slf.AsyncDoByQueue(queueId, func() bool {
			//该函数会被2次调用，但是会排队执行
			return true
		}, func(err error) {
			//函数将在服务协程中执行
		})
	}

	//3.函数参数可以某中一个为空
	//参数二函数将被延迟执行
	slf.AsyncDo(nil, func(err error) {
		//将在下
	})

	//参数一函数在协程池中执行，但没有在服务协程中回调
	slf.AsyncDo(func() bool {
		return true
	}, nil)

	//4.函数返回值控制不进行回调
	slf.AsyncDo(func() bool {
		//返回false时，参数二函数将不会被执行; 为true时，则会被执行
		return false
	}, func(err error) {
		//该函数将不会被执行
	})
}
