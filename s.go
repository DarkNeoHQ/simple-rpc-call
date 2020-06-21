package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// TASK
type Task int

// Void type
type Void struct{}

var a int

// Task方法
func (t *Task) Increment(_ Void, reply *int) error {
	a++
	log.Println("the num on sever is ",a)
	*reply = a
	return nil
}

func runServer() {
	task := new(Task)
	// 发布接收方法
	err := rpc.Register(task)
	if err != nil {
		log.Fatal("service Task error. ", err)
	}
	// 注册http句柄
	rpc.HandleHTTP()
	// 监听1234端口
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error: ", e)
	}
	log.Printf("RPC server on port %d", 1234)
	// 开始接受http
	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}

func main() {
	a = 0
	fmt.Println("server on ")
	runServer()

}
