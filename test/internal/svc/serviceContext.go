package svc

import (
	"fmt"
	"rabbitmq-demo/test/internal/config"
	"rabbitmq-demo/test/queque"
)

type ServiceContext struct {
	Config   config.Config
	RabbitMq queque.RabbitMq
}

func NewServiceContext(c config.Config) *ServiceContext {
	addr := fmt.Sprintf("amqp://%s:%s@%s/", c.RabbitMq.Username, c.RabbitMq.Password, c.RabbitMq.Host)

	rabbit := queque.NewRabbitMq(addr)
	rabbit.InitDelay("test_delay_queue", "test_delay_exchange", "test_delay")
	rabbit.InitQueue("test_queue", "test_exchange", "test_normal")
	return &ServiceContext{
		Config:   c,
		RabbitMq: *rabbit,
	}
}
