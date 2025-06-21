package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("=== Go环境信息 ===")
	fmt.Printf("Go版本: %s\n", getVersion())
	fmt.Printf("当前时间: %s\n", getTimeToTime())
	fmt.Println(getTimeToString())
	fmt.Println(getTimeToTime())
	fmt.Println(getVersion())
	getTimeToString()
	Hello()
	getVersion()
}

//func getTime() string {
//	return time.DateTime
//}

func getTimeToString() string {
	return time.Now().Format(time.RFC850)
}

func getTimeToTime() time.Time {
	return time.Now()
}

func getVersion() string {
	return runtime.Version()
}

func Hello() {
	hour := time.Now().Hour()
	if hour < 12 {
		fmt.Println("Good morning!")
	} else if hour < 17 {
		fmt.Println("Good afternoon!")
	} else {
		fmt.Println("Good evening!")
	}
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
