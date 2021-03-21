package main 

import (
	"os"
	"fmt"
)

func main(){
	// fmt.Println("Hello Worold")
	
	if len(os.Args) > 1{
		fmt.Println("Hello World" ,os.Args[1])
	}
	// return 1 无返回值
	os.Exit(-1)
}

// 应用程序入口必须是 package main
// 应用程序入口必须是 main 方法

// the master has failed more times than the beginner has tried
