# Day 1: 环境搭建与Go基础回顾

## 学习目标
1. 确保Go开发环境正确配置
2. 回顾Go基础语法
3. 理解Go项目结构
4. 掌握Go工具链使用

## 概念学习

### Go环境变量
- `GOPATH`: 工作空间路径（Go 1.11后可选）
- `GOROOT`: Go安装路径
- `GO111MODULE`: 模块管理开关
- `GOPROXY`: 模块代理设置

### Go工具链
- `go run`: 编译并运行
- `go build`: 编译
- `go fmt`: 格式化代码
- `go test`: 运行测试
- `go mod`: 模块管理

## 实践任务

### 任务1：环境检查
1. 在终端运行以下命令，确认Go已正确安装：
```bash
go version
go env
```

2. 记录你的Go版本和重要环境变量

### 任务2：创建第一个程序
1. 在 `exercises/01-basics/day01/` 目录下创建 `hello.go`
2. 实现一个程序，要求：
   - 打印Go版本信息
   - 打印当前时间
   - 打印一个自定义的问候语

提示：
- 使用 `runtime` 包获取Go版本
- 使用 `time` 包处理时间
- 思考如何组织代码结构

### 任务3：探索Go工具
1. 使用 `go fmt` 格式化你的代码
2. 使用 `go build` 编译程序
3. 创建 `go.mod` 文件初始化模块

### 任务4：编写测试
1. 创建 `hello_test.go`
2. 为你的函数编写至少一个测试用例
3. 运行测试并确保通过

## 深入思考
1. 为什么Go强调代码格式的一致性？
2. Go的包管理机制有什么优势？
3. 如何组织一个规范的Go项目结构？

## 提交要求
1. 完成所有练习代码
2. 在 `docs/progress.md` 中更新今日学习记录
3. Git提交信息格式：`feat: complete day01 setup and basics`

## 扩展阅读
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

## 明日预告
Day 2: 切片(Slice)深入理解 - 探索Go中最重要的数据结构之一