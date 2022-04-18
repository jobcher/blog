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
## 数据类型
|类型|表示|备注|
|:----|:----|:----|
|字符串|string|只能用一对双引号（""）或反引号（``）括起来定义，不能用单引号（''）定义！|
|布尔|bool|只有 true 和 false，默认为 false。|
|整型|int8 uint8 int16 uint16 int32 uint32 int64 uint64 int uint|具体长度取决于 CPU 位数。|
|浮点型|float32 float64||

## 常量声明
`常量`，在程序编译阶段就确定下来的值，而程序在运行时无法改变该值。  
### 1. 单个常量声明  
第一种：const 变量名称 数据类型 = 变量值  
如果不赋值，使用的是该数据类型的默认值。  
第二种：const 变量名称 = 变量值  
根据变量值，自行判断数据类型。  
### 2. 多个常量声明
第一种：const 变量名称,变量名称 ... ,数据类型 = 变量值,变量值 ...  
第二种：const 变量名称,变量名称 ... = 变量值,变量值 ...  
### 3. 代码
```go
//demo_1.go
package main

import (
	"fmt"
)

func main() {
	const name string = "Tom"
	fmt.Println(name)

	const age = 30
	fmt.Println(age)

	const name_1, name_2 string = "Tom", "Jay"
	fmt.Println(name_1, name_2)

	const name_3, age_1 = "Tom", 30
	fmt.Println(name_3, age_1)
}
```