package dto

type Rule struct {
	Rules        []*Rule    `json:"rules"`
	Name         string     `json:"name"`
	RoomId       int64      `json:"room_id"`
	Logic        string     `json:"logic"`         //与或逻辑
	RelationalOp string     `json:"relational_op"` //关系运算符(>/</==/!=)
	Indicator    *Indicator `json:"indicator"`
	Value        float64    `json:"value"`
}