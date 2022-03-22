package schedule

import (
	"Alert_demo/core/dal"
	"Alert_demo/core/dal/rule_dao"
	"Alert_demo/core/dal/task_dao"
	"context"
	"fmt"
	"testing"
)

var scheduleServiceImpl = NewScheduleServiceImpl()

func initMysql() {
	err := dal.InitMySQL()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect mysql succeed!")
	}
	dal.DB.AutoMigrate(&rule_dao.RuleEntity{}, &task_dao.TaskEntity{})
}

func TestScheduleServiceImpl_Work(t *testing.T) {
	initMysql()
	defer dal.DB.Close()
	ctx := context.Background()
	scheduleServiceImpl.Work(ctx, 5)
}

//func TestSearchAndExecute(t *testing.T) {
//	initMysql()
//	defer dal.DB.Close()
//	scheduleServiceImpl.SearchAndExecute()
//}
