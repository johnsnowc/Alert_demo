package rule_dao

type RuleEntity struct {
	Id         int64  `gorm:"column:id" json:"id"`
	Code       string `gorm:"column:code" json:"code"`
	Name       string `gorm:"column:name" json:"name"`
	RoomId     int64  `gorm:"column:room_id" json:"room_id"`
	Expr       string `gorm:"column:expr" json:"expr"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
	IsDeleted  bool   `gorm:"column:is_deleted" json:"is_deleted"`
}

func (RuleEntity) TableName() string {
	return "rule"
}

type RuleEntityParams struct {
	Id         int64  `gorm:"column:id" json:"id"`
	Code       string `gorm:"column:code" json:"code"`
	Name       string `gorm:"column:name" json:"name"`
	RoomId     int64  `gorm:"column:room_id" json:"room_id"`
	Expr       string `gorm:"column:expr" json:"expr"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
	IsDeleted  bool   `gorm:"column:is_deleted" json:"is_deleted"`
}
