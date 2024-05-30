package pow

import (
	"atp/payment/pkg/utils/domain"
	"log"
	"strings"
	"time"
)

func (r repository) validator3(validator string, nonce uint, ts int64, prevHash string, data []*domain.Data) (string, bool) {
	guesBlock := &domain.Block{
		Header: &domain.Header{
			Nonce:    nonce,
			Time:     ts,
			PrevHash: prevHash,
		},
		Data: data,
	}
	guesHashstr := guesBlock.Hash()
	//log.Printf("validator3:%s", guesHashstr)
	valid := false
	if strings.Contains(guesHashstr, validator) {
		valid = true
	}

	return guesHashstr, valid
}

func (r repository) PoW3(validator string, ts int64, bc *domain.Blockchain, prevHash string) uint {
	data := r.copyData(bc)

	nonce := 0
	valid := false
	pow := ""

	start := time.Now()
	elapsed := 1 * time.Nanosecond

	for !valid {
		pow, valid = r.validator3(validator, uint(nonce), ts, prevHash, data)
		elapsed = time.Since(start)
		nonce++
	}
	log.Printf("PoW :%s", pow)
	log.Printf("PoW Process:%v", elapsed)
	log.Printf("valid ? %v", valid)

	return uint(nonce)
}
