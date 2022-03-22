package dto

type Result struct {
	RuleCode  string  `json:"rule_code"`
	RoomId    int64   `json:"room_id"`
	Status    *Status `json:"status"`
	CheckTime int64   `json:"check_time"`
	Error     error   `json:"error"`
}

type Alert struct {
	Id        int64  `json:"id"`
	RuleCode  string `json:"rule_code"`
	RoomID    int64  `json:"room_id"`
	Info      []Info `json:"info"`
	CheckTime int64  `json:"check_time"`
}

type Info struct {
	RuleCode       string  `json:"rule_code"`
	Op             string  `json:"op"`
	IndicatorCode  string  `json:"indicator_code"`
	ActualValue    float64 `json:"actual_value"`
	ThresholdValue float64 `json:"threshold_value"`
}
