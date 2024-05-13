package pow

import (
	"atp/payment/pkg/utils/domain"
	"log"
	"strings"
)

func (r repository) validator3(validator string, nonce uint, prevHash string, data []*domain.Data) bool {
	guesBlock := &domain.Block{
		Header: &domain.Header{
			Nonce:    nonce,
			Time:     0,
			PrevHash: prevHash,
		},
		Data: data,
	}
	guesHashstr := guesBlock.Hash()
	log.Printf("validator3:%s", guesHashstr)
	valid := false
	if strings.Contains(guesHashstr, validator) {
		valid = true
	}

	return valid
}

func (r repository) PoW3(validator string, bc *domain.Blockchain, prevHash string) uint {
	data := r.copyData(bc)
	nonce := 0
	for !r.validator3(validator, uint(nonce), prevHash, data) {
		nonce++
	}
	return uint(nonce)
}
