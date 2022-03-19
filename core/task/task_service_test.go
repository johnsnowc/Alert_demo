package task

import (
	"Alert_demo/core/dal"
	"Alert_demo/core/dal/rule_dao"
	"Alert_demo/core/dal/task_dao"
	"Alert_demo/core/dto"
	"context"
	"fmt"
	"log"
	"testing"
)

var taskServiceImpl = NewTaskServiceImpl()

func initMysql() {
	err := dal.InitMySQL()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect mysql succeed!")
	}
	dal.DB.AutoMigrate(&rule_dao.RuleEntity{}, &task_dao.TaskEntity{})
}

func TestTaskServiceImpl_AddTask(t *testing.T) {
	initMysql()
	defer dal.DB.Close()
	ctx := context.Background()
	taskId, err := taskServiceImpl.AddTask(ctx, "taskTest", 1, "simpleRule5__2", 60)
	if taskId == -1 || err != nil {
		t.Errorf("add normal task failed")
	}
	log.Println(taskId)
}

func TestTaskServiceImpl_SelectTaskById(t *testing.T) {
	initMysql()
	defer dal.DB.Close()
	ctx := context.Background()
	if task, err := taskServiceImpl.SelectTaskById(ctx, 3); task == nil || err != nil {
		t.Errorf("select task failed")
	}
	if task, err := taskServiceImpl.SelectTaskById(ctx, 4); task != nil || err == nil {
		t.Errorf("select nonexistent task expected to be failed,but succeed")
	}
}

func TestTaskServiceImpl_SelectTaskByRoomId(t *testing.T) {
	initMysql()
	defer dal.DB.Close()
	ctx := context.Background()
	if tasks, err := taskServiceImpl.SelectTaskByRoomId(ctx, 1); tasks == nil || err != nil || len(tasks) != 2 {
		t.Errorf("select tasks failed")
	}
	if tasks, _ := taskServiceImpl.SelectTaskByRoomId(ctx, 3); len(tasks) > 0 {
		log.Println(tasks)
		t.Errorf("select nonexistent tasks expected to be failed,but succeed")
	}
}

func TestTaskServiceImpl_UpdateTask(t *testing.T) {
	initMysql()
	defer dal.DB.Close()
	ctx := context.Background()
	if taskId, err := taskServiceImpl.UpdateTask(ctx, 1, 1, "", 120); taskId == -1 || err != nil {
		t.Errorf("update task failed")
	}
	if taskId, err := taskServiceImpl.UpdateTask(ctx, 1, 1, "", 120); taskId > 0 || err == nil {
		t.Errorf("use wrong params to update task,expected to be failed,but succeed")
	}
}

func TestTaskServiceImpl_DeleteTask(t *testing.T) {
	initMysql()
	defer dal.DB.Close()
	ctx := context.Background()
	if taskId, err := taskServiceImpl.DeleteTask(ctx, 1); taskId != 1 || err != nil {
		t.Errorf("delete task failed")
	}
}

func TestTaskServiceImpl_IsReady(t *testing.T) {
	initMysql()
	defer dal.DB.Close()
	ctx := context.Background()
	if taskIds, err := taskServiceImpl.IsReady(ctx); len(taskIds) != 2 || err != nil {
		t.Errorf("search ready task failed")
	}
}

func TestTaskServiceImpl_UpdateStatus(t *testing.T) {
	initMysql()
	defer dal.DB.Close()
	ctx := context.Background()
	status := &dto.Status{
		Code: "200",
		Desc: "success",
	}
	if taskId, err := taskServiceImpl.UpdateStatus(ctx, 3, status); taskId != 3 || err != nil {
		t.Errorf("update status failed")
	}
}

func TestTaskServiceImpl_ExecuteTask(t *testing.T) {
	initMysql()
	defer dal.DB.Close()
	ctx := context.Background()
	result, err := taskServiceImpl.ExecuteTask(ctx, 3)
	log.Println(result)
	log.Println(result.Status)
	if err != nil {
		log.Println(err)
		t.Errorf("execute task failed")
	}
}
