package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

func main() {
	dial, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}
	channel, err := dial.Channel()

	exchangeName := "ex_one"
	err = channel.ExchangeDeclare(exchangeName, "fanout", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	declareOne, err := channel.QueueDeclare("", false, false, false, false, nil)
	declareTwo, err := channel.QueueDeclare("", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = channel.QueueBind(declareOne.Name, "", exchangeName, false, nil)
	if err != nil {
		panic(err)
	}
	err = channel.QueueBind(declareTwo.Name, "", exchangeName, false, nil)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		channel.Publish(exchangeName, "", false, false, amqp.Publishing{
			Body: []byte(fmt.Sprintf("%d", i)),
		})
	}

	go consumer(channel, declareOne.Name, "consumer_one")
	go consumer(channel, declareTwo.Name, "consumer_two")

	time.Sleep(20 * time.Second)
}

func consumer(channel *amqp.Channel, queueName, consumer string) {
	consume, err := channel.Consume(queueName, consumer, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	for msg := range consume {
		time.Sleep(time.Second * 1)
		fmt.Println("consumerName:", consumer, msg.Body)
	}
}
