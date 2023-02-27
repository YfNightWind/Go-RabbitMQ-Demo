package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("程序启动")
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	// defer connection.Close()
	defer func(connection *amqp.Connection) {
		err := connection.Close()
		if err != nil {
			panic(err)
		}
	}(connection)
	fmt.Println("连接MQ成功")

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	// defer channel.Close()
	defer func(channel *amqp.Channel) {
		err := channel.Close()
		if err != nil {
			panic(err)
		}
	}(channel)

	// QueueDeclare 声明一个队列来保存消息并传递给消费者。
	// 声明创建一个队列（如果它不存在），或者确保一个现有队列匹配相同的参数。
	queue, err := channel.QueueDeclare("Test", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(queue)

	// 发布消息
	err = channel.Publish("", "Test", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello World"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("成功发送消息到队列当中")
}
