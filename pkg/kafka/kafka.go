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
	host     string
	topic    string
	ctx      context.Context
	mu       *sync.Mutex
	consumer *kafka.Reader
	producer *kafka.Writer
}

func NewKafkaClient(host, topic string, ctx context.Context) Messaging {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{host},
		Topic:   topic,
	})

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{host},
		Topic:   topic,
	})

	return &kafkaClient{
		host:     host,
		topic:    topic,
		ctx:      ctx,
		mu:       &sync.Mutex{},
		consumer: reader,
		producer: writer,
	}
}

func (k *kafkaClient) Read(topic string) error {
	if topic != k.topic {
		return fmt.Errorf("specified topic '%s' does not match the client's topic '%s'", topic, k.topic)
	}

	for {
		select {
		case <-k.ctx.Done():
			// Контекст отменен, выходим из цикла чтения
			return nil
		default:
			k.mu.Lock()

			m, err := k.consumer.ReadMessage(k.ctx)
			if err != nil {
				k.mu.Unlock()
				return err
			}

			// Обработка полученного сообщения
			fmt.Printf("Received message: %s\n", string(m.Value))

			// Дополнительная логика обработки сообщения...

			// Пример: Если получено определенное сообщение, выходим из цикла чтения
			if string(m.Value) == "exit" {
				k.mu.Unlock()
				return nil
			}

			k.mu.Unlock()
		}
	}
}

func (k *kafkaClient) Write(topic string, key, value []byte) error {
	// Проверяем, совпадает ли указанный топик с текущим топиком клиента
	if topic != k.topic {
		return fmt.Errorf("specified topic '%s' does not match the client's topic '%s'", topic, k.topic)
	}

	k.mu.Lock()
	defer k.mu.Unlock()

	message := kafka.Message{
		Key:   key,
		Value: value,
	}

	err := k.producer.WriteMessages(k.ctx, message)
	if err != nil {
		return err
	}

	return nil
}
