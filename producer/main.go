package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "streamprocessing",
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close()

	for i := 1; i < 100; i++ {
		msg := fmt.Sprintf("hello kafka %d", i)
		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(fmt.Sprintf("Key-%d", i)),
				Value: []byte(msg),
			})
		if err != nil {
			log.Fatal("failed to write message : ", err)
		}
		fmt.Println("produced: ", msg)
		time.Sleep(1 * time.Second)
	}

}
