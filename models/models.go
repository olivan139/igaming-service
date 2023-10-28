package models

type Reels [][]string
type Lines []Line
type Payouts []Payout

type Configurations struct {
	Reels   Reels   `json:"reels"`
	Lines   Lines   `json:"lines"`
	Payouts Payouts `json:"payouts"`
}
type Payout struct {
	Symbol    string `json:"symbol"`
	PayoutArr []int  `json:"payout"`
}
type ResponseUnit struct {
	Line   int `json:"line"`
	Payout int `json:"payout"`
}

type PayoutResponse struct {
	Lines []ResponseUnit `json:"lines"`
	Total int            `json:"total"`
}

type Position struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

type Line struct {
	LineVal   int        `json:"line"`
	Positions []Position `json:"positions"`
}

func (l Line) GetPos(index int) (int, int) {
	return l.Positions[index].Col, l.Positions[index].Row
}

func (p Payouts) GetVal(symbol string) ([]int, bool) {
	for i := range p {
		if p[i].Symbol == symbol {
			return p[i].PayoutArr, true
		}
	}

	return []int{}, false
}
