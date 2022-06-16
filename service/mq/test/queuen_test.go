package test

import (
	"fmt"
	"github.com/streadway/amqp"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	dial, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		t.Error(err)
	}
	channel, err := dial.Channel()
	declare, err := channel.QueueDeclare("", false, false, false, false, nil)

	go func() {
		consumerName := "cm1"
		consume, err := channel.Consume(declare.Name, consumerName, true, false, false, false, nil)
		if err != nil {
			t.Error(err)
		}
		for msg := range consume {
			fmt.Println(consumerName, "get message", msg.Body)
		}
	}()

	if err != nil {
		t.Error(err)
	}

	defer func() {
		fmt.Println("overflow")
		//_, err = channel.QueueDelete(declare.Name, false, false, false)
		//if err != nil {
		//	t.Error(err)
		//}
		//err = channel.Close()
		//if err != nil {
		//	t.Error(err)
		//}
	}()

	i := 1
	for {
		time.Sleep(2 * time.Second)

		err = channel.Publish("", declare.Name, false, false, amqp.Publishing{
			Body: []byte(fmt.Sprintf("这里是写入队列中的信息 %d", i)),
		})

		i++

		if err != nil {
			t.Error(err)
		}
	}

}
