package main

import (
	"flag"
	"fmt"
	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/dylenfu/miner-protoactor/messages"
	"reflect"
	"runtime"
	"time"
)

var (
	fn         = flag.String("fn", "server", "select server or client")
	serverAddr = flag.String("server", "localhost:9090", "server address")
	clientAddr = flag.String("client", "localhost:9091", "client address")
)

type Handle struct{}

func main() {
	flag.Parse()
	handle := &Handle{}
	reflect.ValueOf(handle).MethodByName(*fn).Call([]reflect.Value{})
}

type helloActor struct{}

func (*helloActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *messages.HelloRequest:
		ctx.Respond(&messages.HelloResponse{
			Message: "Hello from " + msg.Who,
		})
	}
}
func newHelloActor() actor.Actor {
	return &helloActor{}
}

func (h *Handle) Server() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	remote.Register("hello", actor.FromProducer(newHelloActor))
	remote.Start(*serverAddr)
	console.ReadLine()
}

func (h *Handle) Client() {
	remote.Start(*clientAddr)
	timeout := 1 * time.Second
	pidResp, _ := remote.SpawnNamed(*serverAddr, "remote", "hello", timeout)
	pid := pidResp.Pid
	for i := 0; i < 1000; i++ {
		res, _ := pid.RequestFuture(&messages.HelloRequest{Who: fmt.Sprintf("grade no %d ", i)}, timeout).Result()
		response := res.(*messages.HelloResponse)
		fmt.Printf("Response from remote %v\n", response.Message)
	}

	console.ReadLine()
}
