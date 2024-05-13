package blockchain

import (
	"atp/payment/pkg/adapter/model"
	"atp/payment/pkg/utils/domain"
	"atp/payment/pkg/utils/echos/util"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (bc *blockchain) insertDB(ctx context.Context, info *domain.Block, nowHash string) error {
	errTx := bc.transRepo.DoTransaction(ctx, func(ctxWithTx context.Context, dbt *gorm.DB) error {
		js, _ := json.Marshal(info)
		data := model.Transaction{
			Key:   nowHash,
			Value: string(js),
		}

		_, err := bc.transRepo.FindTransaction(ctx, data.Key)
		if errors.Is(err, util.ErrorNotFound) {

			prev, err := bc.LatestBlock(ctx)
			if errors.Is(err, util.ErrorNotFound) {
				errTx := data.Create(dbt)
				if errTx != nil {
					errN := errors.New("failed create when ErrorNotFound " + fmt.Sprintf("%s :%s", data.Key, errTx.Error()))
					return errN
				}
			} else if err != nil {
				return util.CustomError{
					ErrorType: util.ErrBadRequest,
					Message:   "The Last Block",
					Cause:     "failed get last block",
				}
			} else {
				var decode domain.Block
				_ = json.Unmarshal([]byte(prev.Value), &decode)
				if (decode.Header.PrevHash != info.Header.PrevHash) && (info.Header.PrevHash != "0000000000000000000000000000000000000000000000000000000000000000") {
					errTx := data.Create(dbt)
					if errTx != nil {
						errN := errors.New("failed create when PrevHash " + fmt.Sprintf("%s :%s", data.Key, errTx.Error()))
						return errN
					}
				} else {
					errN := errors.New("failed PrevHash is same")
					return errN
				}
			}

		} else if err != nil {
			return err
		} else {
			errN := errors.New("[warning] key " + fmt.Sprintf("%s : is EXIST", data.Key))
			return errN
		}

		return nil
	})
	if errTx != nil {
		return errTx
	}
	return nil
}
