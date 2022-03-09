package task

import (
	"Alert_demo/core/dal/task_dao"
	"Alert_demo/core/dto"
	i "Alert_demo/core/interface"
	"Alert_demo/core/rule"
	"context"
	"log"
	"strconv"
	"sync"
	"time"
)

type TaskServiceImpl struct {
}

func NewTaskServiceImpl() i.TaskService {
	return &TaskServiceImpl{}
}

var taskDao = task_dao.NewTaskDaoImpl()
var ruleService = rule.NewRuleServiceImpl()

func (t TaskServiceImpl) SelectTaskById(ctx context.Context, id int64) (task *dto.Task, err error) {
	taskEntity, err := taskDao.SelectTaskById(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	task = dalToDto(taskEntity)
	return
}

func (t TaskServiceImpl) SelectTaskByRoomId(ctx context.Context, roomId int64) (tasks []*dto.Task, err error) {
	taskEntitys, err := taskDao.SelectTaskByRoomId(ctx, roomId)
	if err != nil {
		log.Println(err)
		return
	}
	for _, entity := range taskEntitys {
		var task *dto.Task
		task = dalToDto(entity)
		if err != nil {
			log.Println(err)
			return
		}
		tasks = append(tasks, task)
	}
	return
}

func (t TaskServiceImpl) AddTask(ctx context.Context, name string, roomId int64, ruleCode string, frequency int64) (taskId int64, err error) {
	if _, err = ruleService.SelectRuleByCode(ctx, ruleCode); err != nil {
		log.Println(err)
		return -1, err
	}
	params := &task_dao.TaskEntityParams{
		Name:      name,
		RoomId:    roomId,
		RuleCode:  ruleCode,
		Frequency: frequency,
	}
	log.Println("task service next add task")
	taskId, err = taskDao.AddTask(ctx, *params)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return
}

func (t TaskServiceImpl) UpdateTask(ctx context.Context, id int64, roomId int64, ruleCode string, frequency int64) (taskId int64, err error) {
	if _, err = ruleService.SelectRuleByCode(ctx, ruleCode); err != nil {
		log.Println(err)
		return -1, err
	}
	params := &task_dao.TaskEntityParams{
		RoomId:    roomId,
		RuleCode:  ruleCode,
		Frequency: frequency,
	}
	taskId, err = taskDao.UpdateTask(ctx, id, *params)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return
}

func (t TaskServiceImpl) DeleteTask(ctx context.Context, id int64) (taskId int64, err error) {
	taskId, err = taskDao.DeleteTask(ctx, id)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (t TaskServiceImpl) UpdateStatus(ctx context.Context, id int64, status *dto.Status) (taskId int64, err error) {
	params := &task_dao.TaskEntityParams{
		LastTime: time.Now().Unix(),
	}
	if status.Code == "200" {
		params.LastStatus = true
	} else {
		params.LastStatus = false
	}
	taskId, err = taskDao.UpdateTask(ctx, id, *params)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (t TaskServiceImpl) IsReady(ctx context.Context) (taskIds []int64, err error) {
	tasks, err := taskDao.IsReady(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < len(tasks); i++ {
		taskIds = append(taskIds, tasks[i].Id)
	}
	return
}

func (t TaskServiceImpl) ExecuteTask(ctx context.Context, wg sync.WaitGroup, id int64) (result *dto.Result, err error) {
	defer wg.Done()
	taskEntity, err := taskDao.SelectTaskById(ctx, id)
	if err != nil {
		log.Println(err)
		return
	}
	result = &dto.Result{
		RuleCode:  taskEntity.RuleCode,
		RoomId:    taskEntity.RoomId,
		CheckTime: time.Now().Unix(),
	}
	isCorrect, failedCode, err := ruleService.Validate(ctx, taskEntity.RuleCode)
	if err != nil {
		result.Status = &dto.Status{
			Code: "401",
			Desc: "任务运行出现异常，错误为" + err.Error(),
		}
		result.Error = err
		err = t.Alert(ctx, id, result)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = t.UpdateStatus(ctx, id, result.Status)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	if isCorrect == true {
		result.Status = &dto.Status{
			Code: "200",
			Desc: "任务运行成功，校验成功",
		}
		_, err = t.UpdateStatus(ctx, id, result.Status)
		if err != nil {
			log.Println(err)
			return
		}
		return
	} else {
		result.Status = &dto.Status{
			Code: "400",
			Desc: "任务运行成功，校验失败，失败规则code为" + failedCode,
		}
		err = t.Alert(ctx, id, result)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = t.UpdateStatus(ctx, id, result.Status)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
}

func (t TaskServiceImpl) Alert(ctx context.Context, taskId int64, result *dto.Result) (err error) {
	log.Println("任务id" + strconv.FormatInt(taskId, 10) + "发出告警，校验结果为：")
	log.Println(result)
	return nil
}

func dalToDto(entity task_dao.TaskEntity) (task *dto.Task) {
	return &dto.Task{
		Name:      entity.Name,
		RoomId:    entity.RoomId,
		RuleCode:  entity.RuleCode,
		Frequency: entity.Frequency,
		LastTime:  entity.LastTime,
	}
}
