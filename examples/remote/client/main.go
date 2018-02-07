package main

import (
	"fmt"
	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/dylenfu/miner-protoactor/messages"
	"time"
)

const (
	clientAddress = "127.0.0.1:9091"
	serverAddress = "127.0.0.1:9090"
)

func main() {
	remote.Start(clientAddress)
	timeout := 1 * time.Second
	pidResp, _ := remote.SpawnNamed(serverAddress, "remote", "hello", timeout)
	pid := pidResp.Pid
	for i := 0; i < 1000; i++ {
		res, _ := pid.RequestFuture(&messages.HelloRequest{Who: fmt.Sprintf("grade no %d ", i)}, timeout).Result()
		response := res.(*messages.HelloResponse)
		fmt.Printf("Response from remote %v\n", response.Message)
	}

	console.ReadLine()

}
