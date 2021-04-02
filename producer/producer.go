package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	forever := make(chan bool)
	
	producer := NewKafkaProducer()
	
	Publish(msg:"Hello from the other side", topic:"ex1", producer)

	<- forever

}

func NewKafkaProducer() *kafka.Producer {

	configMap := &kafka.ConfigMap{
		"bootstrap.servers":"kafka:9894"
	}

	p, err := kafka.NewProducer(configMap)

	if err != nil {
		log.Fatal(err.Error())
	}

	return p

}