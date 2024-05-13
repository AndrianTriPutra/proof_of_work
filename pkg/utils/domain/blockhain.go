package domain

type Blockchain struct {
	Pool  []*Data  `json:"pool"`
	Chain []*Block `json:"chain"`
}
