package blockchain

import (
	"atp/payment/pkg/utils/domain"
	"context"
	"time"
)

func (bc blockchain) CreateBlockchain(ctx context.Context, nonce uint, prevHash string) *domain.Blockchain {
	nbc := new(domain.Blockchain)
	ts := time.Now().Unix()
	bc.CreateBlock(ctx, ts, nonce, nbc, prevHash)
	return nbc
}
