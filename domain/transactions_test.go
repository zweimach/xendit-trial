package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionMismatch(t *testing.T) {
	inputA := Transaction{
		PaymentRefID: "fae61b32-e941-4909-8df5-ed8e0a85953b",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST4359",
		Amount:       206930,
	}
	inputB := Transaction{
		PaymentRefID: "fae61b32-e941-4909-8df5-ed8e0a85953b",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST4359",
		Amount:       0,
	}
	assert.False(t, inputA.Eq(&inputB))
}

func TestTransactionMatch(t *testing.T) {
	inputA := Transaction{
		PaymentRefID: "fae61b32-e941-4909-8df5-ed8e0a85953b",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST4359",
		Amount:       206930,
	}
	inputB := Transaction{
		PaymentRefID: "fae61b32-e941-4909-8df5-ed8e0a85953b",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST4359",
		Amount:       206930,
	}
	assert.True(t, inputA.Eq(&inputB))
}
