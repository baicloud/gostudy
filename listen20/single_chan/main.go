package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func readData(sendch <-chan int) {
	//sendch<-10
	data := <-sendch
	fmt.Println(data)
}

func main() {
	chnl := make(chan int)
	go sendData(chnl)
	readData(chnl)
}
