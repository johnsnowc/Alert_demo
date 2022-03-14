package dto

type Result struct {
	RuleCode  string  `json:"rule_code"`
	RoomId    int64   `json:"room_id"`
	Status    *Status `json:"status"`
	CheckTime int64   `json:"check_time"`
	Error     error   `json:"error"`
}
