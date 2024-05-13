package transaction

import (
	"atp/payment/pkg/adapter/model"
	"atp/payment/pkg/adapter/sqLite"
	"atp/payment/pkg/utils/echos/util"
	"context"

	"gorm.io/gorm"
)

type repository struct {
	provider sqLite.DatabaseI
}

func NewRepository(provider sqLite.DatabaseI) RepositoryI {
	return repository{
		provider,
	}
}

type RepositoryI interface {
	DoTransaction(ctx context.Context, fn func(ctxWithTx context.Context, dbt *gorm.DB) error) error
	FindTransaction(ctx context.Context, key string) (model.Transaction, error)
	GetLast(ctx context.Context) (model.Transaction, error)
	GetALLTrans(ctx context.Context) ([]model.Transaction, error)
}

func (r repository) DoTransaction(ctx context.Context, fn func(ctxWithTx context.Context, dbt *gorm.DB) error) error {
	return r.provider.WithTransaction(ctx, fn)
}

func (r repository) FindTransaction(ctx context.Context, key string) (model.Transaction, error) {
	var data model.Transaction
	db := r.provider.Db(ctx).(*gorm.DB)
	err := db.Table(model.Transaction{}.TableName()).Where("key = ?", key).First(&data)
	if err.RowsAffected == 0 || err.Error == gorm.ErrRecordNotFound {
		return data, util.ErrorNotFound
	}
	if err.Error != nil {
		return data, err.Error
	}

	return data, nil
}

func (r repository) GetLast(ctx context.Context) (model.Transaction, error) {
	var data model.Transaction
	db := r.provider.Db(ctx).(*gorm.DB)
	err := db.Table(model.Transaction{}.TableName()).Last(&data)
	if err.RowsAffected == 0 || err.Error == gorm.ErrRecordNotFound {
		return data, util.ErrorNotFound
	}
	if err.Error != nil {
		return data, err.Error
	}

	return data, nil
}

func (r repository) GetALLTrans(ctx context.Context) ([]model.Transaction, error) {
	var data []model.Transaction
	db := r.provider.Db(ctx).(*gorm.DB)
	err := db.Table(model.Transaction{}.TableName()).Find(&data)
	if err.RowsAffected == 0 || err.Error == gorm.ErrRecordNotFound {
		return data, util.ErrorNotFound
	}
	if err.Error != nil {
		return data, err.Error
	}

	return data, nil
}
