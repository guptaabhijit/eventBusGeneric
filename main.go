package main

import (
	"github.com/razorpay/eventBusGeneric/driver"
	"log"
	"time"
)

func inMemoryPubSub() {
	// In memory driver for event bus
	pubSub := driver.GetInMemoryEventBus()

	eventBus := driver.New(pubSub)

	eventBus.SubscribeMessage("topic1", make(driver.DataChannel))
	eventBus.SubscribeMessage("topic2", make(driver.DataChannel))

	eventBus.PublishMessage("topic1", "topic 1 data")
	eventBus.PublishMessage("topic2", "topic 2 data")

	eventBus.UnSubscribeMessage("topic2")

	eventBus.PublishMessage("topic1", "new topic 1 data")
	eventBus.PublishMessage("topic2", "new topic 2 data")

}
func redisPubSub() {
	pubSub := driver.GetRedisEvent()

	eventBus := driver.New(pubSub)

	eventBus.SubscribeMessage("topic1", callBackFunc)
	eventBus.SubscribeMessage("topic2", callBackFunc)

	eventBus.PublishMessage("topic1", "topic 1 data")
	eventBus.PublishMessage("topic2", "topic 2 data")

	eventBus.UnSubscribeMessage("topic2")

	eventBus.PublishMessage("topic1", "new topic 1 data")
	eventBus.PublishMessage("topic2", "new topic 2 data")
}

func main() {

	//redisPubSub()
	inMemoryPubSub()

	for {
		time.Sleep(time.Second)
	}
}

func callBackFunc(channel string, payload string) {
	log.Printf("Channel: %v  :: Payload: %v.\n", channel, payload)

}
