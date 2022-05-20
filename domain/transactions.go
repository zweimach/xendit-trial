package domain

const (
	MISMATCH_TRANSACTION = "MISMATCH_TRANSACTION"
	MISSING_IN_A_DATA    = "MISSING_IN_A_DATA"
	MISSING_IN_B_DATA    = "MISSING_IN_B_DATA"
)

type Transaction struct {
	PaymentRefID string  `json:"payment_ref_id"`
	Channel      string  `json:"channel"`
	PaymentCode  string  `json:"payment_code"`
	Amount       float64 `json:"amount"`
	Timestamp    string  `json:"timestamp"`
	PayerName    string  `json:"payer_name"`
}

func (t Transaction) Eq(o Transaction) bool {
	return t.PaymentCode == o.PaymentCode && t.Amount == o.Amount
}

const (
	GroupA = iota
	GroupB
)

type Basket struct {
	Group int `json:"group"`
	Idx   int `json:"idx"`
}

func getTrx(t Basket, xs []Transaction, ys []Transaction) Transaction {
	if t.Group == GroupA {
		return xs[t.Idx]
	}
	return ys[t.Idx]
}

func GroupTrx(
	result map[string]map[string][]Basket,
	group int,
	xs []Transaction,
) {
	for i, t := range xs {
		if _, ok := result[t.Channel]; !ok {
			result[t.Channel] = make(map[string][]Basket)
		}
		if _, ok := result[t.Channel][t.PaymentRefID]; !ok {
			result[t.Channel][t.PaymentRefID] = make([]Basket, 0)
		}
		result[t.Channel][t.PaymentRefID] = append(result[t.Channel][t.PaymentRefID], Basket{group, i})
	}
}

func CompareTrx(xs []Transaction, ys []Transaction) map[string]string {
	result := make(map[string]string)

	// Group transactions by channel and id
	group := make(map[string]map[string][]Basket)
	GroupTrx(group, GroupA, xs)
	GroupTrx(group, GroupB, ys)

	for _, v := range group {
		for id, trx := range v {
			// If there is only one basket
			if len(trx) < 2 {
				if trx[0].Group == GroupA {
					result[id] = MISSING_IN_B_DATA
				} else {
					result[id] = MISSING_IN_A_DATA
				}
				continue
			}

			trx1, trx2 := getTrx(trx[0], xs, ys), getTrx(trx[1], xs, ys)

			// If both transactions are not equal
			if !trx1.Eq(trx2) {
				result[id] = MISMATCH_TRANSACTION
				continue
			}
		}
	}

	return result
}
