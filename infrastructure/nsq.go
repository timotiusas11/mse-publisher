package infrastructure

import (
	"github.com/nsqio/go-nsq"
)

func NewNsqConn() (*nsq.Producer, error) {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}
