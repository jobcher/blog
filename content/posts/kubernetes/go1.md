---
title: "go 基础知识"
date: 2022-04-18
draft: true
author: "jobcher"
tags: ["golang"]
categories: ["k8s"]
series: ["k8s入门系列"]
---
# go 基础知识
## 目录结构
```sh
├─ code  -- 代码根目录
│  ├─ bin
│  ├─ pkg
│  ├─ src
│     ├── hello
│         ├── hello.go
```
- bin 存放编译后可执行的文件。
- pkg 存放编译后的应用包。
- src 存放应用源代码。
  
Hello World 代码  
```go
//在 hello 目录下创建 hello.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
```

## 基础命令

```sh
go build hello
#在src目录或hello目录下执行 go build hello，只在对应当前目录下生成文件。
go install hello
#在src目录或hello目录下执行 go install hello，会把编译好的结果移动到 $GOPATH/bin。
go run hello
#在src目录或hello目录下执行 go run hello，不生成任何文件只运行程序。
go fmt hello
#在src目录或hello目录下执行 go run hello，格式化代码，将代码修改成标准格式。
```
