基于 origin 游戏服务器引擎 搭建的DEMO项目
=========================

[origin](https://github.com/duanhf2012/origin) 是一个由 Go 语言（golang）编写的分布式开源游戏服务器引擎，适用于各类游戏服务器的开发，包括 H5（HTML5）游戏服务器。

本项目主要是用于展示如何使用 origin 进行相关的功能开发，帮助开发者更好的理解 origin 各个模块功能和特性。

快速开始
------
1. 首先安装Git，安装GO的运行和开发环境，这里就不细说了，网上文章很多，IDE推荐用Goland
2. 如果你在国内，需要配置GOPROXY翻墙，常见的代理地址：
    >* 阿里： https://mirrors.aliyun.com/goproxy/
    >* 官方： https://goproxy.io/
    >* 中国：https://goproxy.cn

    以阿里云为例，增加环境变量GOPROXY=https://mirrors.aliyun.com/goproxy/

    可以用```go env```看看是否生效
3. 下载本项目，解压
4. 进入到解压的目录，运行```go mod tidy```安装相关依赖的包（模块）
5. 在目录下运行```go run main.go -start nodeid=1```即可启动服务


