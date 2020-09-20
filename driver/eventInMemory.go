package driver

import (
	"fmt"

	"sync"
)

type DataEvent struct {
	Data  interface{}
	Topic string
}

type DataChannel chan DataEvent

type DataChannelSlice []DataChannel

type EventBus struct {
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

func GetInMemoryEventBus() *EventBus {
	var eb = &EventBus{
		subscribers: map[string]DataChannelSlice{},
	}
	return eb
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.RLock()

	if chans, found := eb.subscribers[topic]; found {
		channels := append(DataChannelSlice{}, chans...)
		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
	eb.rm.RUnlock()
}

func (eb *EventBus) Subscribe(topic string, ch interface{}) {
	eb.rm.Lock()

	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch.(DataChannel))
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch.(DataChannel))
	}

	go Background(ch.(DataChannel))

	eb.rm.Unlock()
}

func (eb *EventBus) UnSubscribe(topic string) {
	eb.rm.Lock()

	delete(eb.subscribers, topic)
	eb.rm.Unlock()
}

func PrintDataEvent(ch string, data DataEvent) {
	fmt.Printf("Channel: %s; Topic: %s; DataEvent: %v\n", ch, data.Topic, data.Data)
}

func Background(ch1 DataChannel) {
	for {
		select {
		case d := <-ch1:
			go PrintDataEvent("ch1", d)
		}
	}
}
