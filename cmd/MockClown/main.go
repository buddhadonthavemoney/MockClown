package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/buddhadonthavemoney/MockClown/pkg/server"
	"github.com/buddhadonthavemoney/MockClown/pkg/storage"
)

func main() {
	app := *cli.NewApp()
	app.Name = "MockClown"
	app.Usage = "MockClown is a simple HTTP server that mocks a response"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "port, p",
			Value: "8080",
			Usage: "Port to listen on",
		},
		cli.StringFlag{
			Name:  "path, pt",
			Value: "/",
			Usage: "Path to listen on",
		},
		cli.StringFlag{
			Name:     "datafile, f",
			Value:    "mock.json",
			Usage:    "File to read data from",
			Required: false,
		},
		cli.StringFlag{
			Name:     "data, d",
			Value:    "",
			Usage:    "Data to return",
			Required: false,
		},
	}
	app.Action = func(c *cli.Context) error {
		port := c.String("port")
		path := c.String("path")
		data := c.String("data")
		datafile := c.String("datafile")

		var jsonData map[string]interface{}
		var err error

		if data == "" {
			jsonData, err = storage.GetJsonFromFile(datafile)
			if err != nil {
				return err
			}
		} else {
			jsonData, err = storage.GetJsonFromString([]byte(data))
			if err != nil {
				return err
			}

		}
		return server.StartServer(port, path, jsonData)
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

