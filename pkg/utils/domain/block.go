package domain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type Block struct {
	Header *Header `json:"header"`
	Data   []*Data `json:"data"`
}

func (b *Block) Hash() string {
	m, _ := json.Marshal(b)
	return fmt.Sprintf("%x", sha256.Sum256(m))
}
