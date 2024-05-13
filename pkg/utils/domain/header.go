package domain

type Header struct {
	Nonce    uint
	Time     int64  `json:"time"`
	PrevHash string `json:"prev_hash"`
}
