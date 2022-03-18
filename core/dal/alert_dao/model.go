package alert_dao

type AlertEntity struct {
	Id        int64  `gorm:"column:id" json:"id"`
	RuleCode  string `gorm:"column:rule_code" json:"rule_code"`
	RoomId    int64  `gorm:"column:room_id" json:"room_id"`
	Info      string `gorm:"column:info" json:"info"`
	CheckTime int64  `gorm:"column:check_time" json:"check_time"`
}

func (AlertEntity) TableName() string {
	return "alert"
}

type AlertEntityParams struct {
	Id        int64  `gorm:"column:id" json:"id"`
	RuleCode  string `gorm:"column:rule_code" json:"rule_code"`
	RoomId    int64  `gorm:"column:room_id" json:"room_id"`
	Info      string `gorm:"column:info" json:"info"`
	CheckTime int64  `gorm:"column:check_time" json:"check_time"`
}
