package _interface

import "context"

type ScheduleService interface {
	Work(ctx context.Context, timeRange int64) (taskId int64, err error)
}
