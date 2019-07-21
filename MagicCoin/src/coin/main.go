package main

import (
	"core"
	"fmt"
)

func main() {
	//初始化CLI对象
	cli := core.CLI{}
	fmt.Println("====返回结果======")
	//执行Run命令调用对应参数对应的方法
	cli.Run()
}