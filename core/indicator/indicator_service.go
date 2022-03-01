package indicator

import (
	"Alert_demo/core/dal/indicator_dao"
	"Alert_demo/core/dto"
	"context"
	"log"
)

type Entity = indicator_dao.IndicatorEntity
type Indicator = dto.Indicator
type IndicatorDaoImpl struct {
}

var Indicators map[string]Entity
var IndicatorTemps []Indicator
var impl indicator_dao.IndicatorDaoImpl

func FindAll(ctx context.Context) {
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

func (i *IndicatorDaoImpl) SelectIndicator(ctx context.Context, code string) (indicator Indicator, err error) {
	entity := Entity{}
	if entity, err = impl.SelectIndicator(ctx, code); err != nil {
		return Indicator{}, err
	}
	indicator = convert(ctx, entity)
	return
}

func (i *IndicatorDaoImpl) AddSimpleIndicator(ctx context.Context, code string, name string, op string, expr string, timeRange int64) (id int64, err error) {
	params := indicator_dao.IndicatorEntityParams{
		Code:      code,
		Name:      name,
		Op:        op,
		Expr:      expr,
		TimeRange: timeRange,
	}
	id, err = impl.AddIndicator(ctx, params)
	return
}

func (i *IndicatorDaoImpl) AddCompleteIndicator(ctx context.Context, code string, name string, left string, right string, op string, timeRange int64) (id int64, err error) {
	params := indicator_dao.IndicatorEntityParams{
		Code:       code,
		Name:       name,
		Op:         op,
		LeftChild:  left,
		RightChild: right,
		TimeRange:  timeRange}
	id, err = impl.AddIndicator(ctx, params)
	return
}

func (i *IndicatorDaoImpl) UpdateIndicator(ctx context.Context, code string, expr string) (id int64, err error) {
	params, _ := impl.SelectIndicator(ctx, code)
	params.Expr = expr
	id, err = impl.UpdateIndicator(ctx, params.Id, indicator_dao.IndicatorEntityParams(params))
	return
}

func (i *IndicatorDaoImpl) DeleteIndicator(ctx context.Context, code string) (id int64, err error) {
	if id, err = impl.DeleteIndicator(ctx, code); err != nil {
		log.Fatal(err)
		return
	}
	return
}
