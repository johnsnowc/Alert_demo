package main

import (
	"Alert_demo/core/dal"
	api "Alert_demo/kitex_gen/api/alert"
	"log"
)

func main() {
	svr := api.NewServer(new(AlertImpl))
	if err := dal.InitMySQL(); err != nil {
		log.Fatal(err)
	}
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
