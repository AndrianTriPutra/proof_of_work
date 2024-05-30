package pow

import (
	"atp/payment/pkg/utils/domain"
	"math/big"
)

type ProofOfWork struct {
	Block  *domain.Blockchain
	Target *big.Int
}

type Setting struct {
	Difficult int
}

type repository struct {
	setting Setting
}

func NewRepository(setting Setting) RepositoryI {
	return repository{
		setting,
	}
}

type RepositoryI interface {
	PoW1(bc *domain.Blockchain, ts int64, prevHash string) uint

	NewProof(b *domain.Blockchain) *ProofOfWork
	Run(pOw *ProofOfWork, ts int64, prevHash string, data []*domain.Data) (uint, []byte)
	Validate(pOw *ProofOfWork, nonce uint, ts int64, prevHash string, data []*domain.Data) bool

	PoW3(validator string, ts int64, bc *domain.Blockchain, prevHash string) uint
}
