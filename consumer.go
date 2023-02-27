package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("消费者启动")
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	defer func(channel *amqp.Channel) {
		err := channel.Close()
		if err != nil {
			panic(err)
		}
	}(channel)

	msg, err := channel.Consume("Test", "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	result := make(chan bool)
	go func() {
		for i := range msg {
			fmt.Printf("收到消息: %s\n", i.Body)
		}
	}()

	fmt.Println("成功连接至MQ")
	fmt.Println("[*]等待消息中")
	<-result
}
