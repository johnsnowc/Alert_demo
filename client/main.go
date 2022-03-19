package main

import (
	"Alert_demo/kitex_gen/api"
	"Alert_demo/kitex_gen/api/alert"
	"context"
	"github.com/cloudwego/kitex/client"
	"log"
	"strconv"
	"time"
)

func main() {
	c, err := alert.NewClient("example", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	addTaskReq := &api.AddTaskRequest{
		Name:      "testTask",
		RoomId:    2,
		RuleCode:  "amount>=200__2",
		Frequency: 40,
	}
	addResp, err := c.AddTask(context.Background(), addTaskReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(addResp)

	id, _ := strconv.Atoi(addResp.Data)
	selectTaskReq := &api.SelectTaskByIdRequest{Id: int64(id)}
	selectTaskResp, err := c.SelectTaskById(context.Background(), selectTaskReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(selectTaskResp)

	time.Sleep(3 * time.Second)
	updateReq := &api.UpdateTaskRequest{
		Id:        int64(id),
		RoomId:    2,
		RuleCode:  "amount>=200__2",
		Frequency: 60,
	}
	updateResp, err := c.UpdateTask(context.Background(), updateReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(updateResp)

	deleteReq := &api.DeleteTaskRequest{Id: int64(id)}
	deleteResp, err := c.DeleteTask(context.Background(), deleteReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(deleteResp)
}
