package handlers

import (
	"Alert_demo/core/dto"
	"Alert_demo/core/task"
	"Alert_demo/kitex_gen/api"
	"context"
	"encoding/json"
	"log"
	"strconv"
)

var taskService = task.NewTaskServiceImpl()

func SelectTaskById(ctx context.Context, req *api.SelectTaskByIdRequest) (resp *api.Response, err error) {
	var temp *dto.Task
	if temp, err = taskService.SelectTaskById(ctx, req.Id); err != nil {
		resp = &api.Response{
			Code:    400,
			Message: "Failed",
			Data:    err.Error(),
		}
		log.Fatal(err)
	}
	data, _ := json.Marshal(temp)
	resp = &api.Response{
		Code:    200,
		Message: "Succeed",
		Data:    string(data),
	}
	return
}

func SelectTaskByRoomId(ctx context.Context, req *api.SelectTaskByRoomIdRequest) (resp *api.Response, err error) {
	var tasks []*dto.Task
	if tasks, err = taskService.SelectTaskByRoomId(ctx, req.RoomId); err != nil {
		resp = &api.Response{
			Code:    400,
			Message: "Failed",
			Data:    err.Error(),
		}
		log.Fatal(err)
	}
	data, _ := json.Marshal(tasks)
	resp = &api.Response{
		Code:    200,
		Message: "Succeed",
		Data:    string(data),
	}
	return
}

func AddTask(ctx context.Context, req *api.AddTaskRequest) (resp *api.Response, err error) {
	var id int64
	if id, err = taskService.AddTask(ctx, req.Name, req.RoomId, req.RuleCode, req.Frequency); err != nil {
		resp = &api.Response{
			Code:    400,
			Message: "Failed",
			Data:    err.Error(),
		}
		log.Fatal(err)
	}
	resp = &api.Response{
		Code:    200,
		Message: "Succeed",
		Data:    strconv.FormatInt(id, 10),
	}
	return
}

func UpdateTask(ctx context.Context, req *api.UpdateTaskRequest) (resp *api.Response, err error) {
	var id int64
	if id, err = taskService.UpdateTask(ctx, req.Id, req.RoomId, req.RuleCode, req.Frequency); err != nil {
		resp = &api.Response{
			Code:    400,
			Message: "Failed",
			Data:    err.Error(),
		}
		log.Fatal(err)
	}
	resp = &api.Response{
		Code:    200,
		Message: "Succeed",
		Data:    strconv.FormatInt(id, 10),
	}
	return
}

func DeleteTask(ctx context.Context, req *api.DeleteTaskRequest) (resp *api.Response, err error) {
	var id int64
	if id, err = taskService.DeleteTask(ctx, req.Id); err != nil {
		resp = &api.Response{
			Code:    400,
			Message: "Failed",
			Data:    err.Error(),
		}
		log.Fatal(err)
	}
	resp = &api.Response{
		Code:    200,
		Message: "Succeed",
		Data:    strconv.FormatInt(id, 10),
	}
	return
}
