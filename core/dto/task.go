package dto

// Task 任务实体
type Task struct {
	Id         int64  `json:"id,omitempty"`
	Name       string `json:"name"`
	RoomId     int64  `json:"room_id"`
	RuleCode   string `json:"rule_code"`
	Frequency  int64  `json:"frequency"`
	LastTime   int64  `json:"last_time"`
	LastStatus Status `json:"last_status"`
}

// Status 运行状态
type Status struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}
