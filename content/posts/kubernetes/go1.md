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

## 变量声明
### 单个变量声明
第一种：var 变量名称 数据类型 = 变量值  
如果不赋值，使用的是该数据类型的默认值。  
第二种：var 变量名称 = 变量值  
根据变量值，自行判断数据类型。  
第三种：变量名称 := 变量值  
省略了 var 和数据类型，变量名称一定要是未声明过的。  
### 多个变量声明
第一种：var 变量名称,变量名称 ... ,数据类型 = 变量值,变量值 ...  
第二种：var 变量名称,变量名称 ... = 变量值,变量值 ...  
第三种：变量名称,变量名称 ... := 变量值,变量值 ...  
### 代码
```go
//demo_2.go
package main

import (
	"fmt"
)

func main() {
	var age_1 uint8 = 31
	var age_2 = 32
	age_3 := 33
	fmt.Println(age_1, age_2, age_3)

	var age_4, age_5, age_6 int = 31, 32, 33
	fmt.Println(age_4, age_5, age_6)

	var name_1, age_7 = "Tom", 30
	fmt.Println(name_1, age_7)

	name_2, is_boy, height := "Jay", true, 180.66
	fmt.Println(name_2, is_boy, height)
}
```

## 输出方法
fmt.Print：输出到控制台（仅只是输出）  
  
fmt.Println：输出到控制台并换行  
  
fmt.Printf：仅输出格式化的字符串和字符串变量（整型和整型变量不可以）  
  
fmt.Sprintf：格式化并返回一个字符串，不输出。  
### 代码
```go
//demo_3.go
package main

import (
	"fmt"
)

func main() {
	fmt.Print("输出到控制台不换行")
	fmt.Println("---")
	fmt.Println("输出到控制台并换行")
	fmt.Printf("name=%s,age=%d\n", "Tom", 30)
	fmt.Printf("name=%s,age=%d,height=%v\n", "Tom", 30, fmt.Sprintf("%.2f", 180.567))
}
```