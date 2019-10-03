package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Consumer struct {
}

func (*Consumer) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

func main() {
	err := initConsumer("order_queue", "first", "122.152.220.31:4161")
	if err != nil {
		fmt.Printf("init consumer failed,err:%v\n", err)
		return
	}
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	<-c
}

func initConsumer(topic string, channel string, address string) error {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = 15 * time.Second
	c, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		return err
	}
	consumer := &Consumer{}
	c.AddHandler(consumer)
	if err := c.ConnectToNSQLookupd(address); err != nil {
		return err
	}
	return nil
}
