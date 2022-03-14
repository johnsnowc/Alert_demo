package handlers

import (
	"Alert_demo/core/schedule"
	"Alert_demo/kitex_gen/api"
	"context"
)

var scheduleService = schedule.NewScheduleServiceImpl()

func Work(ctx context.Context, req *api.WorkRequest) (resp *api.Response, err error) {
	scheduleService.Work(ctx, req.TimeRange)
	select {}
}
