package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("new server")

	app := fiber.New()

	app.Get("/api", api)

	app.Listen(":8002")
}

func init() {
	// hystrix.DefaultTimeout = 500`
	hystrix.ConfigureCommand("hysapiconfig", hystrix.CommandConfig{
		Timeout: 500,
	})
}



//this should be client though
func api(c *fiber.Ctx) error {

	hystrix.Go("hysapiconfig", func() error {

		resp, err := http.Get("http://localhost:8001/")

		if err != nil {
			return err
		}

		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		msg := string(data)
		fmt.Println(msg)

		return nil
	}, func(err error) error {
		fmt.Println(err)
		return nil
	})


	return nil

	// return c.SendString(msg)
}