package task

import (
	"Alert_demo/core/dal/task_dao"
	"Alert_demo/core/dto"
	i "Alert_demo/core/interface"
	"context"
	"log"
)

type TaskServiceImpl struct {
}

func NewTaskServiceImpl() i.TaskService {
	return &TaskServiceImpl{}
}

var td = task_dao.NewTaskDaoImpl()

func (t TaskServiceImpl) SelectTaskById(ctx context.Context, id int64) (task *dto.Task, err error) {
	taskEntity, err := td.SelectTaskById(ctx, id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	task = dalToDto(taskEntity)
	return
}

func (t TaskServiceImpl) SelectTaskByRoomId(ctx context.Context, roomId int64) (tasks []*dto.Task, err error) {
	taskEntitys, err := td.SelectTaskByRoomId(ctx, roomId)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, entity := range taskEntitys {
		var task *dto.Task
		task = dalToDto(entity)
		if err != nil {
			log.Fatal(err)
			return
		}
		tasks = append(tasks, task)
	}
	return
}

func (t TaskServiceImpl) AddTask(ctx context.Context, name string, roomId int64, RuleId int64, frequency int64) (taskId int64, err error) {
	params := &task_dao.TaskEntityParams{
		Name:      name,
		RoomId:    roomId,
		RuleId:    RuleId,
		Frequency: frequency,
	}
	taskId, err = td.AddTask(ctx, *params)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	return
}

func (t TaskServiceImpl) UpdateTask(ctx context.Context, id int64, roomId int64, RuleId int64, frequency int64) (taskId int64, err error) {
	params := &task_dao.TaskEntityParams{
		RoomId:    roomId,
		RuleId:    RuleId,
		Frequency: frequency,
	}
	taskId, err = td.UpdateTask(ctx, id, *params)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	return
}

func (t TaskServiceImpl) DeleteTask(ctx context.Context, id int64) (taskId int64, err error) {
	taskId, err = td.DeleteTask(ctx, id)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func (t TaskServiceImpl) UpdateStatus(ctx context.Context, id int64, status *dto.Status) (dec string, err error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskServiceImpl) IsReady(ctx context.Context) (taskIds []int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskServiceImpl) ExecuteTask(ctx context.Context, id int64) (result *dto.Result, err error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskServiceImpl) Alert(ctx context.Context, taskId int64, result *dto.Result) (err error) {
	//TODO implement me
	panic("implement me")
}

func dalToDto(entity task_dao.TaskEntity) (task *dto.Task) {
	return &dto.Task{
		Name:      entity.Name,
		RoomId:    entity.RoomId,
		RuleId:    entity.RuleId,
		Frequency: entity.Frequency,
		LastTime:  entity.LastTime,
	}
}
