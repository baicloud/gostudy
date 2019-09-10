package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func testCli1() {
	app := cli.NewApp()
	app.Action = func(c *cli.Context) error {
		fmt.Printf("hello %q\n", c.Args().Get(0))
		return nil
	}
	app.Run(os.Args)
}
func testCli2() {
	var language string
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port,p",
			Value: 8000,
			Usage: "listening port",
		},
		cli.StringFlag{
			Name:        "lang,l",
			Value:       "englist",
			Usage:       "read from `FILE`",
			Destination: &language,
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Println(c.String("lang"), c.Int("port"))
		fmt.Println(language)
		return nil
	}
	app.Run(os.Args)
}
func main() {
	// testCli1()
	testCli2()
}
