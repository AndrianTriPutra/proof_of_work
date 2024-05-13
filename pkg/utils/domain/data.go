package domain

type Data struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	IDR  float64 `json:"idr"`
}

func NewData(data Data) *Data {
	m := new(Data)
	m.From = data.From
	m.To = data.To
	m.IDR = data.IDR
	return m
}

func (bc *Blockchain) GiveData(data Data) {
	m := NewData(data)
	bc.Pool = append(bc.Pool, m)
}
