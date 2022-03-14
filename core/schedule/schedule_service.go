package schedule

import (
	i "Alert_demo/core/interface"
	"Alert_demo/core/task"
	"context"
	"github.com/robfig/cron"
	"log"
	"strconv"
	"sync"
)

var taskServiceImpl = task.NewTaskServiceImpl()

type ScheduleServiceImpl struct {
}

func NewScheduleServiceImpl() i.ScheduleService {
	return &ScheduleServiceImpl{}
}

func (s ScheduleServiceImpl) Work(ctx context.Context, timeRange int64) (err error) {
	c := cron.New()
	c.AddFunc("*/"+strconv.FormatInt(timeRange, 10)+" * * * * *", searchAndExecute)
	c.Start()
	defer c.Stop()
	select {}
}

func searchAndExecute() {
	log.Println("start...")
	ctx := context.Background()
	taskIds, err := taskServiceImpl.IsReady(ctx)
	if err != nil {
		log.Println(err)
	}
	var wg sync.WaitGroup
	for i := 0; i < len(taskIds); i++ {
		wg.Add(1)
		go taskServiceImpl.ExecuteTask(ctx, wg, taskIds[i])
	}
	wg.Wait()
}
