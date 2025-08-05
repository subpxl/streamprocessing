package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

const workerCOunt = 5

var specialCharRegex = regexp.MustCompile(`[^a-zA-Z0-9\s]+`)

type messageData struct {
	offset int64
	key    string
	value  string
}

func main() {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"kafka:9092"},
		Topic:     "streamprocessing",
		GroupID:   "go-consumer-group",
		Partition: 0,
		MinBytes:  1,
		MaxBytes:  10e6,
	})
	defer reader.Close()

	fmt.Println("listenting for messagess")

	jobs := make(chan messageData, 100)

	var wg sync.WaitGroup

	for i := 0; i < workerCOunt; i++ {
		wg.Add(1)
		go worker(jobs, &wg)
	}

	// consume message from kafka
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		m, err := reader.ReadMessage(ctx)
		cancel()

		if err != nil {
			log.Println("error reading message:", err)
			continue
		}
		jobs <- messageData{
			offset: m.Offset,
			key:    string(m.Key),
			value:  string(m.Value),
		}
	}

	// close(jobs)
	// wg.Wait()
// 
}

func worker(jobs <-chan messageData, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range jobs {
		processedValue := processData(msg.value)
		logMsg := fmt.Sprintf("received message at offset %d: key %s, value %s",
			msg.offset, msg.key, processedValue)
		fmt.Println(logMsg)
		sendLogToAnalytics(logMsg)
	}
}

func processData(input string) string {
	cleaned := specialCharRegex.ReplaceAllString(input, "")
	return stringUpper(cleaned)
}

func stringUpper(s string) string {
	return strings.ToUpper(s)

}
func sendLogToAnalytics(logMsg string) {
	http.Post(
		"http://analytics-service:8000/log",
		"application/text",
		bytes.NewBufferString(logMsg),
	)
}
