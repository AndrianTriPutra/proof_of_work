package pow

import (
	"atp/payment/pkg/utils/domain"
	"log"
	"strings"
)

func (r repository) validator1(nonce uint, prevHash string, data []*domain.Data) bool {
	prefixExpected := strings.Repeat("0", r.setting.Difficult)

	guesBlock := &domain.Block{
		Header: &domain.Header{
			Nonce:    nonce,
			Time:     0,
			PrevHash: prevHash,
		},
		Data: data,
	}

	guesHashstr := guesBlock.Hash()
	log.Printf("guesHashstr:%s", guesHashstr)

	return guesHashstr[:r.setting.Difficult] == prefixExpected
}

func (r repository) copyData(bc *domain.Blockchain) []*domain.Data {
	data := make([]*domain.Data, 0)
	for _, v := range bc.Pool {
		vdata := domain.Data{
			From: v.From,
			To:   v.To,
			IDR:  v.IDR,
		}
		data = append(data, domain.NewData(vdata))
	}
	return data
}

func (r repository) PoW1(bc *domain.Blockchain, prevHash string) uint {
	data := r.copyData(bc)
	nonce := 0
	for !r.validator1(uint(nonce), prevHash, data) {
		nonce++
	}
	return uint(nonce)
}
