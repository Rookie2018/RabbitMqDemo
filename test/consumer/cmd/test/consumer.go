package main

import (
	"encoding/json"
	"fmt"
	"log"
	"rabbitmq-demo/test/queque"
)

func main() {

	Username := "guest"
	Password := "guest"
	Host := "127.0.0.1:5673"

	mqUrl := fmt.Sprintf("amqp://%s:%s@%s/", Username, Password, Host)
	rabbit := queque.NewRabbitMq(mqUrl)

	msgChan, err := rabbit.Channel.Consume(
		"test_queue",
		"test_normal",
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
	type mesage struct {
		TestId  int    `json:"testId"`
		TestStr string `json:"testStr"`
	}

	go func() {
		for d := range msgChan {
			log.Printf("收到消息: %s", d.Body)

			msg := mesage{}
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				fmt.Println("解析失败:", err)
			}
			if msg.TestId == 31 {
				d.Nack(true, true)
				continue
			}

			//取消未支付订单的操作，使用grpc发送消息
			d.Ack(true)

			// d.Nack(true, true)
		}
	}()

	log.Printf("wait message")
	<-forever
}
