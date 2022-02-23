package _interface

import (
	"Alert_demo/core/dto"
	"context"
)

type TaskService interface {
	SelectTaskById(ctx context.Context, id int64) (Task *dto.Task, err error)

	SelectTaskByRoomId(ctx context.Context, roomId int64) (Tasks []*dto.Task, err error)

	AddTask(ctx context.Context, name string, roomId int64, RuleId int64, frequency int64) (TaskId int64, err error)

	UpdateTask(ctx context.Context, id int64, roomId int64, RuleId int64, frequency int64) (TaskId int64, err error)

	DeleteTask(ctx context.Context, id int64) (TaskId int64, err error)

	UpdateStatus(ctx context.Context, id int64, status *dto.Status) (dec string, err error)

	IsReady(ctx context.Context) (taskIds []int64, err error)

	ExecuteTask(ctx context.Context, id int64) (result *dto.Result, err error)

	Alert(ctx context.Context, taskId int64, result *dto.Result) (err error)
}
