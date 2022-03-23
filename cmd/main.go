package main

import (
	"Alert_demo/core/dal"
	"Alert_demo/core/dal/indicator_dao"
	"Alert_demo/core/dal/rule_dao"
	"Alert_demo/core/dal/task_dao"
	"Alert_demo/core/schedule"
	"context"
	"fmt"
)

func main() {
	err := dal.InitMySQL()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect mysql succeed!")
	}
	dal.DB.AutoMigrate(&indicator_dao.IndicatorEntity{}, &rule_dao.RuleEntity{}, &task_dao.TaskEntity{})
	scheduleServiceImpl := schedule.NewScheduleServiceImpl()
	scheduleServiceImpl.Work(context.Background(), 60*60*2)
}
