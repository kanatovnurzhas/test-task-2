package kafka

import (
	"context"
	"fmt"
	"sync"

	"github.com/segmentio/kafka-go"
)

type Messaging interface {
	Read(topic string) error
	Write(topic string, key, value []byte) error
}

type kafkaClient struct {
	host      string
	ctx       context.Context
	mu        *sync.Mutex
	consumers map[string]*kafka.Reader
	producers map[string]*kafka.Writer
	chName    chan []byte
	chAnswer  chan []byte
}

func NewKafkaClient(host string, topics []string, ctx context.Context, chName, chAnswer chan []byte) Messaging {
	consumers := make(map[string]*kafka.Reader)

	for _, topic := range topics {
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{host},
			Topic:   topic,
		})
		consumers[topic] = reader
	}

	producers := make(map[string]*kafka.Writer)
	for _, topic := range topics {
		writer := kafka.NewWriter(kafka.WriterConfig{
			Brokers: []string{host},
			Topic:   topic,
		})
		producers[topic] = writer
	}

	return &kafkaClient{
		host:      host,
		ctx:       ctx,
		mu:        &sync.Mutex{},
		consumers: consumers,
		producers: producers,
		chName:    chName,
		chAnswer:  chAnswer,
	}
}

func (k *kafkaClient) Read(topic string) error {
	reader, ok := k.consumers[topic]
	if !ok {
		return fmt.Errorf("Unknown topic: %s", topic)
	}

	for {
		m, err := reader.ReadMessage(k.ctx)
		if err != nil {
			return err
		}

		// Обработка полученного сообщения
		fmt.Printf("Received message: %s\n", string(m.Value))
		// Дополнительная логика обработки сообщения...
		if topic == "stud-to-course" || topic == "course-to-stud" {
			fmt.Println("Read function: ", m.Value)
			k.chName <- m.Value
		} else if topic == "answer-for-stud" || topic == "answer-for-course" {
			fmt.Println("Stud topic")
			k.chAnswer <- m.Value
		}

	}
}

func (k *kafkaClient) Write(topic string, key, value []byte) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	writer, ok := k.producers[topic]
	if !ok {
		return fmt.Errorf("Unknown topic: %s", topic)
	}

	message := kafka.Message{
		Key:   key,
		Value: value,
	}

	err := writer.WriteMessages(k.ctx, message)
	if err != nil {
		return err
	}

	return nil
}
