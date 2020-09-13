package main

import (
	"github.com/razorpay/pubSub/driver"
	"time"
)

func redisEvent() {
	var redisPubSub = driver.GetPubSub()

	// Create a subscriber
	redisPubSub.Subscribe("topic1", nil)
	redisPubSub.Subscribe("topic2", nil)

	//log.Print("Subscriptions done. Publishing...")

	redisPubSub.Publish("topic1", "topic 1 data")
	redisPubSub.Publish("topic2", "topic 2 data")
	redisPubSub.Publish("topic1", "topic 1 again data")
	redisPubSub.Publish("topic2", "topic 2 again data")
}

func inMemoryEvent() {
	eb := driver.GetEventBus()

	ch1 := make(driver.DataChannel)
	ch2 := make(driver.DataChannel)

	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch2)

	eb.Publish("topic1", "topic 1 data ")
	eb.Publish("topic2", "topic 2 data ")

	go driver.Background(ch1, ch2)
}

func main() {
	redisEvent()
	inMemoryEvent()

	for {
		time.Sleep(time.Second)
	}
}
