package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

func main() {
	nsqAddress := "122.152.220.31:4150"
	err := initProducer(nsqAddress)
	if err != nil {
		fmt.Printf("init producer failed,err", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read string failed,err:", err)
			return
		}
		data = strings.TrimSpace(data)
		if data == "stop" {
			break
		}
		err = producer.Publish("order_queue", []byte(data))
		if err != nil {
			fmt.Printf("publish message failed,err:", err)
			return
		}
		fmt.Printf("publish data:%s succ\n", data)
	}
}

func initProducer(str string) error {
	var err error
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(str, config)
	if err != nil {
		return err
	}
	return nil
}
