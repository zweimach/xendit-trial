package domain

type Transaction struct {
	PaymentRefID string  `json:"payment_ref_id"`
	Channel      string  `json:"channel"`
	PaymentCode  string  `json:"payment_code"`
	Amount       float64 `json:"amount"`
	Timestamp    string  `json:"timestamp"`
	PayerName    string  `json:"payer_name"`
}

const (
	MATCHING_TRANSACTION = "MATCHING_TRANSACTION"
	MISMATCH_TRANSACTION = "MISMATCH_TRANSACTION"
	MISSING_IN_A_DATA    = "MISSING_IN_A_DATA"
	MISSING_IN_B_DATA    = "MISSING_IN_B_DATA"
)

func CompareTransaction(a Transaction, b Transaction) string {
	if a.PaymentCode != b.PaymentCode || a.Amount != b.Amount {
		return MISMATCH_TRANSACTION
	}
	return MATCHING_TRANSACTION
}
