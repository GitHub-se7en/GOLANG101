package main

import (
	_func "GOLANG101/101/func"
	"GOLANG101/101/runner"
	"log"
	_ "os"
	"time"
)

const timeout = 10 * time.Second

func main() {
	//chapter3.Hello()
	//zhizhen.GetAddress()
	//getnotvalue.Value()
	//notify.Notify()
	//words.Words()
	//println("I am ", os.Args[1])
	//args.Args()
	//test.TestDownload()
	log.Println("开始工作")
	//为本次执行分配时间
	//r := runner.New(timeout)
	//r.Add(createTask(),createTask(),createTask())
	//
	// err := r.Start()
	//if err != nil {
	//	switch err {
	//	case runner.ErrorTimeOut:
	//		log.Println("时间超时，结束程序")
	//		os.Exit(1)
	//	case runner.ErrorInterrupt:
	//		log.Println("用户终止，结束程序")
	//		os.Exit(2)
	//
	//	}
	//}
	_func.Function()

}
func createTask() func(int) {
	return func(i int) {
		log.Printf("执行者#%d.", i)
		time.Sleep(time.Duration(i) * time.Second)
	}
}
