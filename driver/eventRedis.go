package driver

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
	"time"
)

type Subscriber struct {
	pubsub   *redis.PubSub
	channel  string
	callback processFunc
}

type processFunc func(string, string)

type PubSub struct {
	Client *redis.Client
	sub    Subscriber
}

var Service *PubSub

func GetRedisEvent() *PubSub {
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

func (s *PubSub) Subscribe(channel string, chanOrCallBack interface{}) {
	var err error

	subscriber := Subscriber{
		pubsub:   Service.Client.Subscribe(),
		channel:  channel,
		callback: chanOrCallBack.(func(string, string)),
	}

	err = subscriber.pubsub.Subscribe(channel)
	if err != nil {
		log.Println("Error subscribing to channel.")
	}

	s.sub = subscriber

	go subscriber.listen()
}

func (s *PubSub) UnSubscribe(topic string) {
	err := s.sub.pubsub.Unsubscribe(topic)
	if err != nil {
		log.Printf("Error Unsubscribing to topic %v.\n", topic)
	}
}

func (s *Subscriber) listen() error {
	var channel string
	var payload string

	for {
		msg, _ := s.pubsub.ReceiveTimeout(time.Second)

		channel = ""
		payload = ""

		switch m := msg.(type) {
		case *redis.Subscription:
			log.Printf("Subscription Message: %v to channel '%v'. %v total subscriptions.\n", m.Kind, m.Channel, m.Count)
			continue
		case *redis.Message:
			channel = m.Channel
			payload = m.Payload
		}

		if len(channel) == 0 {
			continue
		}

		//log.Printf("channel: %s Topic: %s\n", channel, payload)

		go s.callback(channel, payload)
	}
}
