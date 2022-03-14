package handlers

import (
	"Alert_demo/core/dto"
	"Alert_demo/core/task"
	"Alert_demo/kitex_gen/api"
	"context"
	"log"
)

func SelectTaskById(ctx context.Context, req *api.SelectTaskIdRequest) (resp *api.SelectTaskResponse, err error) {
	taskService := task.NewTaskServiceImpl()
	var temp *dto.Task
	if temp, err = taskService.SelectTaskById(ctx, req.Id); err != nil {
		log.Fatal(err)
	}
	resp = new(api.SelectTaskResponse)
	resp.Name = temp.Name
	resp.Id = temp.Id
	resp.RoomId = temp.RoomId
	resp.RuleId = temp.RuleId
	return
}

func SelectTaskByRoomId(ctx context.Context, req *api.SelectTaskRoomIdRequest) (resp *api.SelectTasksResponse, err error) {
	taskService := task.NewTaskServiceImpl()
	var tasks []*dto.Task
	if tasks, err = taskService.SelectTaskByRoomId(ctx, req.RoomId); err != nil {
		log.Fatal(err)
	}
	resp = new(api.SelectTasksResponse)
	for _, each := range tasks {
		temp := &api.SelectTaskResponse{each.Id, each.RoomId, each.Name, each.RuleId}
		resp.Tasks = append(resp.Tasks, temp)
	}
	return
}

func AddTask(ctx context.Context, req *api.AddTaskRequest) (resp *api.Response, err error) {
	taskService := task.NewTaskServiceImpl()
	resp = new(api.Response)
	if _, err = taskService.AddTask(ctx, req.Name, req.RoomId, req.RuleId, req.Frequency); err != nil {
		resp.State = "Failed"
		log.Fatal(err)
	}
	resp.State = "Succeed"
	return
}

func UpdateTask(ctx context.Context, req *api.UpdateTaskRequest) (resp *api.Response, err error) {
	taskService := task.NewTaskServiceImpl()
	resp = new(api.Response)
	if _, err = taskService.UpdateTask(ctx, req.Id, req.RoomId, req.RuleId, req.Frequency); err != nil {
		resp.State = "Failed"
		log.Fatal(err)
	}
	resp.State = "Succeed"
	return
}

func DeleteTask(ctx context.Context, req *api.DeleteTaskRequest) (resp *api.Response, err error) {
	taskService := task.NewTaskServiceImpl()
	resp = new(api.Response)
	if _, err = taskService.DeleteTask(ctx, req.Id); err != nil {
		resp.State = "Failed"
		log.Fatal(err)
	}
	resp.State = "Succeed"
	return
}
