package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"sync"
)

type Messaging interface {
	Consume(topic string) error
	Produce(topic string, key, value []byte) error
	Close() error
}

type kafkaClient struct {
	host    string
	readers map[string]*kafka.Reader
	writers map[string]*kafka.Writer
	ctx     context.Context
	mu      *sync.Mutex
}
