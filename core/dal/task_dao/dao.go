package task_dao

import (
	"context"
)

type TaskDao interface {
	IsReady(ctx context.Context) (tasks []TaskEntity, err error)

	SelectTaskById(ctx context.Context, id int64) (Task TaskEntity, err error)

	SelectTaskByRoomId(ctx context.Context, roomId int64) (Tasks []TaskEntity, err error)

	AddTask(ctx context.Context, params TaskEntityParams) (TaskId int64, err error)

	UpdateTask(ctx context.Context, id int64, params TaskEntityParams) (TaskId int64, err error)

	DeleteTask(ctx context.Context, id int64) (TaskId int64, err error)
}
