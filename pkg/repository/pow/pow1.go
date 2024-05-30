package pow

import (
	"atp/payment/pkg/utils/domain"
	"log"
	"strings"
	"time"
)

func (r repository) validator1(nonce uint, ts int64, prevHash string, data []*domain.Data) (string, bool) {
	prefixExpected := strings.Repeat("0", r.setting.Difficult)

	guesBlock := &domain.Block{
		Header: &domain.Header{
			Nonce:    nonce,
			Time:     ts,
			PrevHash: prevHash,
		},
		Data: data,
	}

	guesHashstr := guesBlock.Hash()
	//log.Printf("guesHashstr:%s", guesHashstr)

	return guesHashstr, guesHashstr[:r.setting.Difficult] == prefixExpected
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

func (r repository) PoW1(bc *domain.Blockchain, ts int64, prevHash string) uint {
	data := r.copyData(bc)

	nonce := 0
	valid := false
	pow := ""

	start := time.Now()
	elapsed := 1 * time.Nanosecond

	for !valid {
		pow, valid = r.validator1(uint(nonce), ts, prevHash, data)
		elapsed = time.Since(start)
		nonce++
	}
	log.Printf("PoW :%s", pow)
	log.Printf("PoW Process:%v", elapsed)
	log.Printf("valid ? %v", valid)

	return uint(nonce)
}
