package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "streamprocessing",
		GroupID:   "go-consumer-group",
		Partition: 0,
		MinBytes:  1,
		MaxBytes:  10e6,
	})
	defer reader.Close()

	fmt.Println("listenting for messagess")

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		m, err := reader.ReadMessage(ctx)
		cancel()

		if err != nil {
			log.Println("error reading message", err)
			continue
		}
		fmt.Printf("recieved message at offer %f: key %d , value %s\n",
			m.Offset, string(m.Key), string(m.Value))
	}

}
