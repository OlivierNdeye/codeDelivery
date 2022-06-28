package kafka

import ckafka "github.com/confluent/confluent-kafka-go/kafka"

type kafkaConsumer struct {
	msgChan chan *
}