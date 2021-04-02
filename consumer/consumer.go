package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	ch := make(chan kafka.Message)

	go Consume(ch)

	for msg := range ch {
		fmt.Println(a...:"Receiving message:  ", string(msg.Value))
	}

}

func Consume(ch chan kafka.Message)  {
	configMap := kafka.ConfigMap {
		"bootstrap.servers":"host.docker.internal:9894"
		"group.id":"mygroup"
	}

	c, err := kafka.NewConsumer(&configMap)

	if err != nil {
		log.Fatal(err.Error())
	}

	consumer.Subscribe(topic:"ex2", rebalanceCb: nil)

	for {
		msg, err := consumer.ReadMessage(timeout: -1)

		if err != nil {
			ch <- *msg
		}
	}

}

