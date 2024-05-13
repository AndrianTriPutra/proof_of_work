package blockchain

import (
	"atp/payment/pkg/utils/domain"
	"context"
)

func (bc blockchain) CreateBlockchain(ctx context.Context, nonce uint, prevHash string) *domain.Blockchain {
	nbc := new(domain.Blockchain)
	bc.CreateBlock(ctx, nonce, nbc, prevHash)
	return nbc
}
