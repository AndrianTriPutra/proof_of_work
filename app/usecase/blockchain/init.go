package blockchain

import (
	"atp/payment/pkg/adapter/model"
	"atp/payment/pkg/repository/transaction"
	"atp/payment/pkg/utils/domain"
	"context"
)

type blockchain struct {
	transRepo transaction.RepositoryI
}

func NewBlockChain(transRepo transaction.RepositoryI) Usecase {
	return &blockchain{
		transRepo: transRepo,
	}
}

type Usecase interface {
	CreateBlockchain(ctx context.Context, nonce uint, prevHash string) *domain.Blockchain
	CreateBlock(ctx context.Context, ts int64, nonce uint, a *domain.Blockchain, prevHash string) *domain.Block
	LatestBlock(ctx context.Context) (model.Transaction, error)
}
