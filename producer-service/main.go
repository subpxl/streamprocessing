package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"kafka:9092"},
		Topic:    "streamprocessing",
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close()

	// mock data generation
	// in real world this will come from a api like stock , telemetry etc

	mockStrings := []string{
		"hello#world!",
		"foo123bar",
		"!@#nonsense$%^",
		"kafka-msg-001",
		"Te$t%Data&",
		"random_string_!@#",
		"123_@!ABC",
		"cleanTHISstring",
		"filter*&out!!",
		"UPPER_lower MIX",
	}
	for i, msg := range mockStrings {
		// msg := fmt.Sprintf("hello kafka %d", i)
		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(fmt.Sprintf("Key-%d", i)),
				Value: []byte(msg),
			})
		if err != nil {
			log.Fatal("failed to write message : ", err)
		}
		fmt.Println("produced: ", msg)
		sendAnalytics(fmt.Sprintf("produced: %s", msg))
		time.Sleep(1 * time.Second)
	}

}

func sendAnalytics(logMsg string) {
	http.Post(
		"http://analytics-service:8000/log",
		"application/text",
		bytes.NewBufferString(logMsg))
}
