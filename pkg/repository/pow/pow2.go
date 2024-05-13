package pow

import (
	"atp/payment/pkg/utils/domain"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/big"
)

func (r repository) NewProof(b *domain.Blockchain) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-r.setting.Difficult))
	log.Printf("target:%v", target)
	pow := &ProofOfWork{b, target}
	return pow
}

func (r repository) initData(nonce uint, prevHash string, data []*domain.Data) []byte {
	guesBlock := &domain.Block{
		Header: &domain.Header{
			Nonce:    nonce,
			Time:     0,
			PrevHash: prevHash,
		},
		Data: data,
	}

	js, _ := json.Marshal(guesBlock)

	return js
}

func (r repository) Run(pOw *ProofOfWork, prevHash string, data []*domain.Data) (uint, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0
	for nonce < math.MaxInt64 {
		data := r.initData(uint(nonce), prevHash, data)
		hash = sha256.Sum256(data)

		fmt.Printf("\rRun %d:%x", nonce, hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pOw.Target) == -1 {
			break
		} else {
			nonce++
		}

	}
	fmt.Println()

	return uint(nonce), hash[:]
}

func (r repository) Validate(pOw *ProofOfWork, nonce uint, prevHash string, data []*domain.Data) bool {
	var intHash big.Int

	dataX := r.initData(nonce, prevHash, data)
	hash := sha256.Sum256(dataX)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pOw.Target) == -1
}
