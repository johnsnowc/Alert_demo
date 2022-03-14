package indicator_dao

import (
	"Alert_demo/core/dal"
	"context"
	"log"
)

type IndicatorDaoImpl struct {
}

func NewIndicatorDaoImpl() IndicatorDao {
	return &IndicatorDaoImpl{}
}

func (i *IndicatorDaoImpl) SelectIndicator(ctx context.Context, code string) (indicator IndicatorEntity, err error) {
	if err = dal.DB.Debug().Where("code = ? AND is_deleted = ?", code, 0).Find(&indicator).Error; err != nil {
		log.Println(err)
		return IndicatorEntity{}, err
	}
	return
}

func (i *IndicatorDaoImpl) SelectAllIndicators(ctx context.Context) (indicators map[string]IndicatorEntity, err error) {
	var ins []IndicatorEntity
	if err = dal.DB.Debug().Where(" is_deleted = ?", 0).Find(&ins).Error; err != nil {
		log.Println(err)
		return
	}
	indicators = make(map[string]IndicatorEntity)
	for _, indicator := range ins {
		indicators[indicator.Code] = indicator
	}
	return
}

func (i *IndicatorDaoImpl) AddIndicator(ctx context.Context, params IndicatorEntityParams) (id int64, err error) {
	indicatorEntity := IndicatorEntity(params)
	if err = dal.DB.Debug().Create(&indicatorEntity).Error; err != nil {
		log.Println(err)
		return -1, err
	}
	return params.Id, nil
}

func (i *IndicatorDaoImpl) UpdateIndicator(ctx context.Context, id int64, params IndicatorEntityParams) (indicatorId int64, err error) {
	indicatorEntity := IndicatorEntity(params)
	if err = dal.DB.Debug().Model(&IndicatorEntity{}).Where("id = ? AND is_deleted = ?", id, 0).Updates(indicatorEntity).Error; err != nil {
		log.Println(err)
		return id, err
	}
	return id, nil
}

func (i *IndicatorDaoImpl) DeleteIndicator(ctx context.Context, code string) (id int64, err error) {
	if err = dal.DB.Debug().Model(&IndicatorEntity{}).Where("code = ?", code).Update("is_deleted", true).Error; err != nil {
		log.Println(err)
		return id, err
	}
	return id, nil
}

func (i *IndicatorDaoImpl) QueryData(ctx context.Context, code string) (result float64, err error) {
	entity, _ := i.SelectIndicator(ctx, code)
	row := dal.DB.Debug().Raw(entity.Expr).Row()
	row.Scan(&result)
	return
}
