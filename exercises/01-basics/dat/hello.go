package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(getTime())
	fmt.Println(getVersion())
	getTime()
	getVersion()
}

func getTime() string {
	return time.DateTime
}

func getVersion() string {
	return runtime.Version()
}

// 深入思考
//1. 为什么Go强调代码格式的一致性？
//回答: 这样开发者就不需要花费时间在代码格式上了
//2. Go的包管理机制有什么优势？
//回答: 管理机制从GOPATH --> GO Vendor --> Go mod
//从任意项目都在固定的GOPATH
//到可以每个项目都在自己的目录 依赖存储到Vendor目录
//再后来有一个Go Modules
//让项目能够自己目录管理 并且依赖包能够有了版本控制  更加方便快捷准确的管理依赖 并且能够有一系列的go mod  go get命令能够将依赖下载记录在go.mod文件中
//3. 如何组织一个规范的Go项目结构？
//在我的了解中 Go项目一般有cmd用于管理主程序入口   pkg用于管理通用的组件   service用来管理服务相关的代码  model用来管理结构体
//经过查询  官方推荐cmd   internal pkg 这样的结构
