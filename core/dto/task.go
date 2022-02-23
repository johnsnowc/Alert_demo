package dto

// Task 任务实体
type Task struct {
	Name       string `json:"name"`
	RoomId     int64  `json:"room_id"`
	RuleId     string `json:"rule_id"`
	Frequency  int64  `json:"frequency"`
	LastTime   int64  `json:"last_time"`
	LastStatus Status `json:"last_status"`
}

// Status 运行状态
type Status struct {
	Code   string `json:"code"`
	Status string `json:"status"`
	Desc   string `json:"desc"`
}

// Context 创建上下文
type Context struct {
	Name      string `json:"name"`
	RoomId    int64  `json:"room_id"`
	RuleId    string `json:"rule_id"`
	Frequency int64  `json:"frequency"`
}
