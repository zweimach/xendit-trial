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

func (t *Transaction) PartialEq(o *Transaction) bool {
	return t.PaymentRefID == o.PaymentRefID && t.Channel == o.Channel
}

func (t *Transaction) Eq(o *Transaction) bool {
	return t.PartialEq(o) && t.PaymentCode == o.PaymentCode && t.Amount == o.Amount
}
