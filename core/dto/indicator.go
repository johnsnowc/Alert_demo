package dto

type Indicator struct {
	Code         string     `json:"code"`
	LeftChild    *Indicator `json:"left_child"`  //若为普通指标，则左右子树为null
	RightChild   *Indicator `json:"right_child"` //计算型指标表示为一棵二叉树
	Name         string     `json:"name"`
	ValueType    bool       `json:"value_type"`    //指标类型（计算型、普通型）
	ObjectType   ObjectType `json:"object_type"`   //实体类型，用于普通型指标，结构中expr为sql语句用于查询指标值
	ArithmeticOp string     `json:"arithmetic_op"` //算术运算符，用于计算型指标，若该指标是普通型则为null(客单价：总价/客户人数)
	TimeRange    int64      `json:"time_range"`    //指标查询时间范围
}

type ObjectType struct {
	Name string `json:"name"`
	Expr string `json:"expr"`
}
