package task_dao

import (
	"Alert_demo/core/dal"
	"context"
	"log"
	"time"
)

type TaskDaoImpl struct {
}

func NewTaskDaoImpl() TaskDao {
	return &TaskDaoImpl{}
}

func (t TaskDaoImpl) SelectTaskById(ctx context.Context, id int64) (task TaskEntity, err error) {
	if err = dal.DB.Debug().Where("id = ? AND is_deleted = ?", id, 0).First(&task).Error; err != nil {
		log.Println(err)
		return TaskEntity{}, err
	}
	return
}

func (t TaskDaoImpl) SelectTaskByRoomId(ctx context.Context, roomId int64) (tasks []TaskEntity, err error) {
	if err = dal.DB.Debug().Where("room_id = ? AND is_deleted = ?", roomId, 0).Find(&tasks).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

func (t TaskDaoImpl) AddTask(ctx context.Context, params TaskEntityParams) (taskId int64, err error) {
	taskEntity := TaskEntity(params)
	if err = dal.DB.Debug().Create(&taskEntity).Error; err != nil {
		log.Println(err)
		return -1, err
	}
	var ids []int64
	dal.DB.Raw("select LAST_INSERT_ID() as id").Pluck("id", &ids)
	return ids[0], nil
}

func (t TaskDaoImpl) UpdateTask(ctx context.Context, id int64, params TaskEntityParams) (taskId int64, err error) {
	taskEntity := TaskEntity(params)
	if err = dal.DB.Debug().Model(&TaskEntity{}).Where("id = ? AND is_deleted = ?", id, 0).Updates(taskEntity).Error; err != nil {
		log.Println(err)
		return id, err
	}
	return id, nil
}

func (t TaskDaoImpl) DeleteTask(ctx context.Context, id int64) (taskId int64, err error) {
	if err = dal.DB.Debug().Model(&TaskEntity{}).Where("id = ?", id).Update("is_deleted", true).Error; err != nil {
		log.Println(err)
		return id, err
	}
	return id, nil
}

func (t TaskDaoImpl) IsReady(ctx context.Context) (tasks []TaskEntity, err error) {
	if err = dal.DB.Debug().Select("id").Where("last_time + frequency <= ? AND is_deleted = ?", time.Now().Unix(), 0).Find(&tasks).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return
}
