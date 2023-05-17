package main

import (
	"flag"
	"fmt"
	"log"
	"rabbitmq-demo/test/consumer/config"
	"rabbitmq-demo/test/queque"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "../../etc/consumer.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	Username := "guest"
	Password := "guest"
	Host := "127.0.0.1:5673"

	mqUrl := fmt.Sprintf("amqp://%s:%s@%s/", Username, Password, Host)
	rabbit := queque.NewRabbitMq(mqUrl)

	msgChan, err := rabbit.Channel.Consume(
		"test_delay_queue",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	defer rabbit.Conn.Close()
	defer rabbit.Channel.Close()

	if err != nil {
		panic(err.Error())
	}

	forever := make(chan bool)

	go func() {
		for d := range msgChan {
			log.Printf("收到消息: %s", d.Body)
			//取消未支付订单的操作，使用grpc发送消息

			d.Ack(false)
		}
	}()

	log.Printf("wait message")
	<-forever
}
