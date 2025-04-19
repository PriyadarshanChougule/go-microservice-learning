package main

import (
	"fmt"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("rabbit mq service")
	rabitConn, err := connect()
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}

}

func connect() (*amqp.Connection, error) {
	var Count int
	var BackOff = 1 * time.Second
	var Connection *amqp.Connection

	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost")

		if err != nil {
			fmt.Println("RabbitMQ not yet ready")
			Count++
		} else {
			Connection = c
			break
		}

		if Count > 5 {
			fmt.Println(err)
			return nil, err
		}

		BackOff = time.Duration(math.Pow(float64(Count), 2)) * time.Second
		fmt.Println("backing off...")
		time.Sleep(BackOff)
		continue
	}

	return Connection, nil
}
