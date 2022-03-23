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
	c.AddFunc("@every "+strconv.FormatInt(timeRange, 10)+"s", searchAndExecute)
	c.Start()
	defer c.Stop()
	select {}
}

func (s ScheduleServiceImpl) SearchAndExecute() {
	searchAndExecute()
}

func searchAndExecute() {
	log.Println("start...")
	ctx := context.Background()
	taskIds, err := taskServiceImpl.IsReady(ctx)
	if err != nil {
		log.Println(err)
	}
	var wg sync.WaitGroup
	for j := 0; j < len(taskIds); j++ {
		wg.Add(1)
		go func(index int) {
			taskServiceImpl.ExecuteTask(ctx, taskIds[index])
			wg.Done()
		}(j)
	}
	wg.Wait()
	log.Println("done...")
}
