package main

import (
	api "Alert_demo/kitex_gen/api/alert"
	"log"
)

func main() {
	svr := api.NewServer(new(AlertImpl))
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
