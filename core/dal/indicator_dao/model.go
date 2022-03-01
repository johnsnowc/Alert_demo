package indicator_dao

type IndicatorEntity struct {
	Id         int64  `gorm:"column:id" json:"id"`
	Code       string `gorm:"column:code" json:"code"`
	Name       string `gorm:"column:name" json:"name"`
	Type       bool   `gorm:"column:type" json:"type"`
	LeftChild  string `gorm:"column:left_child" json:"left_child"`
	RightChild string `gorm:"column:right_child" json:"right_child"`
	Op         string `gorm:"column:op" json:"op"`
	Expr       string `gorm:"column:expr" json:"expr"`
	TimeRange  int64  `gorm:"column:time_range" json:"time_range"`
	CreateTime string `gorm:"column:create_time" json:"create_time"`
	UpdateTime string `gorm:"column:update_time" json:"update_time"`
	IsDeleted  bool   `gorm:"column:is_deleted" json:"is_deleted"`
}

func (IndicatorEntity) TableName() string {
	return "indicator"
}

type IndicatorEntityParams struct {
	Id         int64  `gorm:"column:id" json:"id"`
	Code       string `gorm:"column:code" json:"code"`
	Name       string `gorm:"column:name" json:"name"`
	Type       bool   `gorm:"column:type" json:"type"`
	LeftChild  string `gorm:"column:left_child" json:"left_child"`
	RightChild string `gorm:"column:right_child" json:"right_child"`
	Op         string `gorm:"column:op" json:"op"`
	Expr       string `gorm:"column:expr" json:"expr"`
	TimeRange  int64  `gorm:"column:time_range" json:"time_range"`
	CreateTime string `gorm:"column:create_time" json:"create_time"`
	UpdateTime string `gorm:"column:update_time" json:"update_time"`
	IsDeleted  bool   `gorm:"column:is_deleted" json:"is_deleted"`
}
