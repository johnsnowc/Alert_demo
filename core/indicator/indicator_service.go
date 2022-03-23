package indicator

import (
	"Alert_demo/core/dal/indicator_dao"
	"Alert_demo/core/dto"
	i "Alert_demo/core/interface"
	"context"
	"errors"
	"log"
)

type Entity = indicator_dao.IndicatorEntity

type Indicator = dto.Indicator

type IndicatorServiceImpl struct {
}

var Indicators = make(map[string]Entity, 0)

var IndicatorTemps []Indicator

var impl indicator_dao.IndicatorDaoImpl

func NewIndicatorServiceImpl() i.IndicatorService {
	return &IndicatorServiceImpl{}
}

//func init() {
//	if len(Indicators) == 0 {
//		dal.InitMySQL()
//		FindAll(nil)
//	}
//}
func FindAll(ctx context.Context) {
	if Indicators == nil {
		Indicators = make(map[string]Entity)
	}
	temps, _ := impl.SelectAllIndicators(ctx)
	Indicators = temps
}

func Find(ctx context.Context, code string) (entity Entity, err error) {
	if entity, err = impl.SelectIndicator(ctx, code); err != nil {
		return Entity{}, err
	}
	return
}

func simpleConvert(ctx context.Context, entity Entity) (indicator Indicator) {
	if entity.Type == true {
		return
	}
	object := dto.ObjectType{Name: entity.Name, Expr: entity.Expr}
	indicator = Indicator{
		Name:         entity.Name,
		Code:         entity.Code,
		ValueType:    entity.Type,
		ObjectType:   object,
		ArithmeticOp: entity.Op,
		TimeRange:    entity.TimeRange,
	}
	return
}

func computeConvert(ctx context.Context, entity Entity) (indicator Indicator) {
	if entity.Type == false {
		return
	}

	leftTemp := Indicators[entity.LeftChild]
	rightTemp := Indicators[entity.RightChild]

	if leftTemp.Type == false {
		IndicatorTemps = append(IndicatorTemps, simpleConvert(ctx, leftTemp))
	} else {
		IndicatorTemps = append(IndicatorTemps, computeConvert(ctx, leftTemp))
	}
	if rightTemp.Type == false {
		IndicatorTemps = append(IndicatorTemps, simpleConvert(ctx, rightTemp))
	} else {
		IndicatorTemps = append(IndicatorTemps, computeConvert(ctx, rightTemp))
	}
	indicator = Indicator{
		entity.Name,
		entity.Code,
		&IndicatorTemps[len(IndicatorTemps)-2],
		&IndicatorTemps[len(IndicatorTemps)-1],
		entity.Type,
		dto.ObjectType{},
		entity.Op,
		indicator.TimeRange,
	}
	return
}

func convert(ctx context.Context, entity Entity) (indicator Indicator) {
	if entity.Type == true {
		indicator = computeConvert(ctx, entity)
	} else {
		indicator = simpleConvert(ctx, entity)
	}
	return
}

func (i *IndicatorServiceImpl) SelectIndicator(ctx context.Context, code string) (indicator Indicator, err error) {
	entity := Entity{}
	if len(Indicators) == 0 {
		FindAll(nil)
	}
	if entity, err = impl.SelectIndicator(ctx, code); err != nil {
		return Indicator{}, err
	}
	indicator = convert(ctx, entity)
	return
}

func (i *IndicatorServiceImpl) AddSimpleIndicator(ctx context.Context, code string, name string, expr string, timeRange int64) (id int64, err error) {
	params := indicator_dao.IndicatorEntityParams{
		Code:      code,
		Name:      name,
		Type:      false,
		Expr:      expr,
		TimeRange: timeRange,
	}
	if id, err = impl.AddIndicator(ctx, params); err != nil {
		log.Fatal(err)
	}
	Indicators[code] = Entity(params)
	return
}

func (i *IndicatorServiceImpl) AddCompleteIndicator(ctx context.Context, code string, name string, left string, right string, op string, timeRange int64) (id int64, err error) {
	params := indicator_dao.IndicatorEntityParams{
		Code:       code,
		Name:       name,
		Op:         op,
		Type:       true,
		LeftChild:  left,
		RightChild: right,
		TimeRange:  timeRange,
	}
	if id, err = impl.AddIndicator(ctx, params); err != nil {
		log.Fatal(err)
	}
	Indicators[code] = Entity(params)
	return
}

func (i *IndicatorServiceImpl) UpdateIndicator(ctx context.Context, timeRange int64, code, name, left, right, op, expr string) (id int64, err error) {
	params, _ := impl.SelectIndicator(ctx, code)
	if params.Type == true {
		params.LeftChild = left
		params.RightChild = right
		params.Op = op
	} else {
		params.Expr = expr
	}
	params.Name = name
	params.TimeRange = timeRange
	id, err = impl.UpdateIndicator(ctx, params.Id, indicator_dao.IndicatorEntityParams(params))
	if err != nil {
		log.Fatal(err)
	}
	Indicators[code] = params
	return
}

func (i *IndicatorServiceImpl) DeleteIndicator(ctx context.Context, code string) (id int64, err error) {
	if id, err = impl.DeleteIndicator(ctx, code); err != nil {
		log.Fatal(err)
		return
	}
	delete(Indicators, code)
	return
}

func (i *IndicatorServiceImpl) QueryData(ctx context.Context, code string, roomId int64) (data float64, err error) {
	temp, _ := impl.SelectIndicator(ctx, code)
	if temp.Type == false {
		data, err = impl.QueryData(ctx, code, roomId)
	} else {
		left, _ := impl.SelectIndicator(ctx, temp.LeftChild)
		right, _ := impl.SelectIndicator(ctx, temp.RightChild)
		dataLeft, _ := i.QueryData(ctx, left.Code, roomId)
		dataRight, _ := i.QueryData(ctx, right.Code, roomId)
		switch temp.Op {
		case "*":
			data = dataRight * dataLeft
		case "/":
			if dataRight == 0 {
				err = errors.New("division by zero")
				log.Println(err)
			} else {
				data = dataLeft / dataRight
			}
		case "+":
			data = dataRight + dataLeft
		case "-":
			data = dataLeft - dataRight
		}
	}
	return
}
