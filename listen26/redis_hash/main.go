package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "192.168.56.101:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer c.Close()
	a, err := c.Do("HSet", "books", "abc", 101)
	fmt.Println(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(c.Do("HGet", "books", "abc"))
	if err != nil {
		fmt.Println("get abc failed", err)
		return
	}
	fmt.Println(r)
}
