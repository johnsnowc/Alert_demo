package task_dao

type TaskEntity struct {
	Id         int64 `gorm:"column:id" json:"id"`
	Name       int64 `gorm:"column:name" json:"name"`
	RoomId     int64 `gorm:"column:room_id" json:"room_id"`
	RuleId     int64 `gorm:"column:rule_id" json:"rule_id"`
	Frequency  int64 `gorm:"column:frequency" json:"frequency"`
	LastTime   int64 `gorm:"column:last_time" json:"last_time"`
	LastStatus bool  `gorm:"column:last_status" json:"last_status"`
	CreateTime int64 `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64 `gorm:"column:update_time" json:"update_time"`
	IsDeleted  bool  `gorm:"column:is_deleted" json:"is_deleted"`
}

func (TaskEntity) TableName() string {
	return "task"
}

type TaskEntityParams struct {
	Id         int64 `gorm:"column:id" json:"id"`
	Name       int64 `gorm:"column:name" json:"name"`
	RoomId     int64 `gorm:"column:room_id" json:"room_id"`
	RuleId     int64 `gorm:"column:rule_id" json:"rule_id"`
	Frequency  int64 `gorm:"column:frequency" json:"frequency"`
	LastTime   int64 `gorm:"column:last_time" json:"last_time"`
	LastStatus bool  `gorm:"column:last_status" json:"last_status"`
	CreateTime int64 `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64 `gorm:"column:update_time" json:"update_time"`
	IsDeleted  bool  `gorm:"column:is_deleted" json:"is_deleted"`
}
