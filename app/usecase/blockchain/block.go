package blockchain

import (
	"atp/payment/pkg/adapter/model"
	"atp/payment/pkg/utils/domain"
	"context"
	"log"
	"time"
)

func (bc blockchain) NewBlock(ctx context.Context, nonce uint, prevHash string, data []*domain.Data) *domain.Block {
	b := new(domain.Block)
	b.Data = data

	b.Header = &domain.Header{
		Nonce:    nonce,
		Time:     time.Now().Unix(),
		PrevHash: prevHash,
	}
	nowHash := b.Hash()

	err := bc.insertDB(ctx, b, nowHash)
	if err != nil {
		log.Printf("FAILED NewBlock insertDB:" + err.Error())
	}

	return b
}

func (bc *blockchain) CreateBlock(ctx context.Context, nonce uint, a *domain.Blockchain, prevHash string) *domain.Block {
	b := bc.NewBlock(ctx, nonce, prevHash, a.Pool)
	a.Chain = append(a.Chain, b)
	a.Pool = []*domain.Data{}
	return b
}

func (bc *blockchain) LatestBlock(ctx context.Context) (model.Transaction, error) {
	last, err := bc.transRepo.GetLast(ctx)
	return last, err
}
