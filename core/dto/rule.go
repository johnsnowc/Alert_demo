package dto

type Rule struct {
	Id            int64       `json:"id,omitempty"`
	Code          string      `json:"code"`
	Rules         []*Rule     `json:"rules"`
	Name          string      `json:"name"`
	RoomId        int64       `json:"room_id"`
	Logic         string      `json:"logic"`         //与或逻辑
	RelationalOp  string      `json:"relational_op"` //关系运算符(>/</==/!=)
	IndicatorCode string      `json:"indicator_code"`
	Value         interface{} `json:"value"`
}
