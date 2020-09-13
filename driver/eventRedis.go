package driver

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

type Subscriber struct {
	pubsub  *redis.PubSub
	channel string
	//callback processFunc
}

type PubSub struct {
	Client *redis.Client
}

var Service *PubSub

func GetPubSub() *PubSub {
	var client *redis.Client

	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 10,
	})
	Service = &PubSub{Client: client}
	return Service
}

func (s *PubSub) Publish(channel string, message interface{}) {
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	messageString := string(jsonBytes)
	pub := s.Client.Publish(channel, messageString)

	if err = pub.Err(); err != nil {
		log.Print("PublishString() error", err)
	}
}

//func (s *PubSub) PublishString(channel, message string) *redis.IntCmd {
//	return s.client.Publish(channel, message)
//}

func (s *PubSub) Subscribe(channel string, chanOrCallBack interface{}) {
	var err error

	subscriber := Subscriber{
		pubsub:  Service.Client.Subscribe(),
		channel: channel,
		//callback: fn,
	}

	err = subscriber.pubsub.Subscribe(channel)
	if err != nil {
		log.Println("Error subscribing to channel.")
	}

	go subscriber.listen()
}

func (s *Subscriber) listen() error {
	var channel string
	var payload string

	for {
		msg, _ := s.pubsub.ReceiveTimeout(time.Second)
		//if err != nil {
		//	if reflect.TypeOf(err) == reflect.TypeOf(&net.OpError{}) && reflect.TypeOf(err.(*net.OpError).Err).String() == "*net.timeoutError" {
		//		// Timeout, ignore
		//		continue
		//	}
		//	// Actual error
		//	log.Print("Error in ReceiveTimeout()", err)
		//}

		channel = ""
		payload = ""

		switch m := msg.(type) {
		case *redis.Subscription:
			log.Printf("Subscription Message: %v to channel '%v'. %v total subscriptions.", m.Kind, m.Channel, m.Count)
			continue
		case *redis.Message:
			channel = m.Channel
			payload = m.Payload
		}

		if len(channel) == 0 {
			continue
		}

		fmt.Printf("channel: %s Topic: %s\n", channel, payload)

		//go s.callback(channel, payload)
	}
}
