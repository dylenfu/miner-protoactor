package dao_test

import (
	"fmt"
	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"testing"
	"time"
)

type Hello struct{ Who string }
type Ask struct{ Who string }
type SetBehaviorActor struct{}

func (state *SetBehaviorActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("Hello %v\n", msg.Who)
		context.SetBehavior(state.Other)

	case Ask:
		fmt.Printf("%v, u ask me what?\n", msg.Who)
		context.Respond("answer is " + msg.Who)
	}
}

func (state *SetBehaviorActor) Other(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("%v, we are now handling messages in another behavior\n", msg.Who)
	}
}

func NewSetBehaviorActor() actor.Actor {
	return &SetBehaviorActor{}
}

func Test_ProtoActorTell(t *testing.T) {
	props := actor.FromProducer(NewSetBehaviorActor)
	pid := actor.Spawn(props)
	for i := 0; i < 5; i++ {
		pid.Tell(Hello{Who: "Roger"})
	}
	console.ReadLine()
}

func Test_ProtoActorRequest(t *testing.T) {
	props := actor.FromProducer(NewSetBehaviorActor)
	pid := actor.Spawn(props)
	pid.Request(Ask{Who: "Tom"}, pid)
	console.ReadLine()
}

func Test_ProtoActorRequestFuture(t *testing.T) {
	props := actor.FromProducer(NewSetBehaviorActor)
	pid := actor.Spawn(props)
	for i := 0; i < 500; i++ {
		result, _ := pid.RequestFuture(Ask{Who: fmt.Sprintf("Mike %d", i)}, 1*time.Second).Result()
		t.Log(result)
	}
	console.ReadLine()
}
