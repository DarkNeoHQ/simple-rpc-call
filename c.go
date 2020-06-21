package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

// Task for use with RPC
type Task int

// Void type
type Void struct{}

var a int


func runClient() {
	var err error
	var reply int
//47.102.196.129
	// 创建链接
	client, err := rpc.DialHTTP("tcp", "47.102.196.129:1234")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	for {
		client.Call("Task.Increment", struct{}{}, &reply)//rpccall
		fmt.Printf("Called Task.Increment recv: %d\n", reply)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	a = 0

	fmt.Println("client on")
	runClient()

}
