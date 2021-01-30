package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

const (
	topic          = "mysqlServer.mysqldb.Sensor"
	broker1Address = "kafka:9092" //localhost
	groupId        = "localgroup"
)

func main() {
	ctx := context.Background()
	consumeTopic(ctx)
}

func consumeTopic(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
		GroupID: groupId,
	})
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println("received: ", string(msg.Value))
	}
}
