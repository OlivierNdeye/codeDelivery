package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluent/confluent-kafka-go/kafka"
)

type kafkaConsumer struct {
	msgChan chan *ckafka.Message
}

func (k *kafkaConsumer) consumer() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
		"group.id":          os.Getenv("kafkaConsumerGroupId"),
	}
	c, err := ckafka.NewCosumer(configMap)
	if err != nil {
		log.Fatal("error consuming kafka message" + err.Error())
	}
	topics := []string{os.Getenv("kafkaTopics")}
	c.SubscribeTopics(topics, nil)
	fmt.Println("kafka consumer has been started...")
}
