package indicator_dao

type IndicatorEntity struct {
	Id         int64  `gorm:"column:id" json:"id"`
	Code       int64  `gorm:"column:code" json:"code"`
	Name       int64  `gorm:"column:name" json:"name"`
	Type       bool   `gorm:"column:type" json:"type"`
	LeftChild  int64  `gorm:"column:left_child" json:"left_child"`
	RightChild int64  `gorm:"column:right_child" json:"right_child"`
	Op         string `gorm:"column:op" json:"op"`
	Expr       string `gorm:"column:expr" json:"expr"`
	TimeRange  int64  `gorm:"column:time_range" json:"time_range"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
	IsDeleted  bool   `gorm:"column:is_deleted" json:"is_deleted"`
}
