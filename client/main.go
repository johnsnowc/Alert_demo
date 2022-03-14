package main

import (
	"Alert_demo/kitex_gen/api"
	"Alert_demo/kitex_gen/api/alert"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"log"
	"time"
)

func main() {
	c, err := alert.NewClient("example", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := &api.SelectIndicatorByCodeRequest{Code: "test-2022-03-07"}
	deleteResp, err := c.SelectIndicator(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(deleteResp)
}
