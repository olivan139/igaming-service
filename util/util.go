package util

import (
	"igaming-service/errs"
	"igaming-service/models"
)

func GetPayoff(conf models.Configurations) (models.PayoutResponse, error) {
	var payoutResponse models.PayoutResponse

	for i := range conf.Lines {
		var line []string

		for j := range conf.Lines[i].Positions {
			col, row := conf.Lines[i].GetPos(j)
			symbol, err := symbolInReels(conf.Reels, row, col)

			if err != nil {
				return models.PayoutResponse{}, err
			}

			line = append(line, symbol)
		}

		payout, err := getLinePayoff(line, conf.Payouts)

		if err != nil {
			return models.PayoutResponse{}, err
		}

		payoutResponse.Lines = append(payoutResponse.Lines, models.ResponseUnit{
			Line:   conf.Lines[i].LineVal,
			Payout: payout,
		})

		payoutResponse.Total += payout
	}

	return payoutResponse, nil
}

func symbolInReels(reels models.Reels, row int, col int) (string, error) {
	if row >= len(reels) || col >= len(reels[row]) {
		return "", errs.ErrOutOfRange
	}

	return reels[row][col], nil
}

func getLinePayoff(line []string, payouts models.Payouts) (int, error) {
	if len(line) == 0 {
		return 0, errs.ErrLineNotFound
	}

	symbol := line[0]
	i := 1

	for ; i < len(line) && symbol == line[i]; i++ {
	}

	payout, exists := payouts.GetVal(symbol)

	if !exists {
		return 0, errs.ErrUndefinedSymbol
	}

	if i > len(line) {
		return 0, errs.ErrOutOfRange
	}

	return payout[i-1], nil
}
