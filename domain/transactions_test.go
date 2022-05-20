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

func TestGroupTrx(t *testing.T) {
	output := make(map[string]map[string][]Basket)

	GroupTrx(output, GroupA, INLINE_A_DATA)
	GroupTrx(output, GroupA, INLINE_B_DATA)

	assert.Len(t, output, 2)
	assert.Contains(t, output, "ID_ALFAMART")
	assert.Contains(t, output, "ID_INDOMARET")

	assert.Len(t, output["ID_ALFAMART"], 9)
	assert.Len(t, output["ID_INDOMARET"], 12)
}

func TestCompareTrx(t *testing.T) {
	output := CompareTrx(INLINE_A_DATA, INLINE_B_DATA)

	expected := map[string]string{
		"3f48ac25-4964-4e5f-8703-82508999dbbd":   MISMATCH_TRANSACTION,
		"55735054-9873-4ff1-a65b-df9cb4ecb0b7":   MISMATCH_TRANSACTION,
		"819759c3-cd2a-431a-a5f8-66e22ced49a1":   MISMATCH_TRANSACTION,
		"8w84wjduw-b540-4d67-9f78-9ae7f7430c30":  MISSING_IN_B_DATA,
		"d897c533-494912-4d67-9f78-9ae7f7430c30": MISSING_IN_A_DATA,
		"f59cf91b-d6b2-4def-bc14-0c4cc4befa63":   MISMATCH_TRANSACTION,
	}

	assert.Len(t, output, len(expected))
	for k, v := range expected {
		assert.Contains(t, output, k)
		assert.Equal(t, output[k], v)
	}
}

var INLINE_A_DATA = []Transaction{
	{
		PaymentRefID: "0780433d-67f7-460c-8b34-39542aab4f7f",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST4759",
		Amount:       12282.1,
	},
	{
		PaymentRefID: "819759c3-cd2a-431a-a5f8-66e22ced49a1",
		Channel:      "ID_ALFAMART",
		PaymentCode:  "TEST33723",
		Amount:       40839.14,
	},
	{
		PaymentRefID: "268d743d-48bf-4ef1-a959-02cd822d1b28",
		Channel:      "ID_ALFAMART",
		PaymentCode:  "TEST4901",
		Amount:       25291.44,
	},
	{
		PaymentRefID: "c7a902e2-01af-437f-a6a6-894d79f96b00",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST3627",
		Amount:       46957.63,
	},
	{
		PaymentRefID: "f59cf91b-d6b2-4def-bc14-0c4cc4befa63",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST4491",
		Amount:       16333.93,
	},
	{
		PaymentRefID: "1d7e9004-d0a0-4d1e-97be-bb41a75c79f8",
		Channel:      "ID_ALFAMART",
		PaymentCode:  "TEST4335",
		Amount:       16025.97,
	},
	{
		PaymentRefID: "cc159e81-d531-41bb-8491-a956df8e7cb8",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST3505",
		Amount:       45158.05,
	},
	{
		PaymentRefID: "3f48ac25-4964-4e5f-8703-82508999dbbd",
		Channel:      "ID_ALFAMART",
		PaymentCode:  "TEST3471",
		Amount:       13345.4,
	},
	{
		PaymentRefID: "55735054-9873-4ff1-a65b-df9cb4ecb0b7",
		Channel:      "ID_ALFAMART",
		PaymentCode:  "TEST4641",
		Amount:       31971.27,
	},
	{
		PaymentRefID: "55fd4e3e-0c93-4f9d-aaef-89d7b13e963c",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST3523",
		Amount:       39247.7,
	},
	{
		PaymentRefID: "8870f4e3-689a-48bd-8531-68cb5ff75b15",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST3466",
		Amount:       32587.54,
	},
	{
		PaymentRefID: "8w84wjduw-b540-4d67-9f78-9ae7f7430c30",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST84783",
		Amount:       4916.39,
	},
	{
		PaymentRefID: "64d0f980-c8df-4939-9fdc-5aeb67429b1b",
		Channel:      "ID_ALFAMART",
		PaymentCode:  "TEST4540",
		Amount:       33016.57,
	},
	{
		PaymentRefID: "ff7c86d1-352e-47e3-a9e2-f83815a90870",
		Channel:      "ID_ALFAMART",
		PaymentCode:  "TEST3866",
		Amount:       24139.05,
	},
	{
		PaymentRefID: "8345869e-1bb2-4c90-a72a-aef53c33bda8",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST3348",
		Amount:       26149.49,
	},
	{
		PaymentRefID: "78aff9c3-1ac5-4c58-8f50-19dfd7cc64c9",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST4194",
		Amount:       2222.45,
	},
	{
		PaymentRefID: "5369e1ac-dfaf-43ac-a340-75cdeb2f7f53",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST3590",
		Amount:       38659.22,
	},
	{
		PaymentRefID: "fae61b32-e941-4909-8df5-ed8e0a85953b",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST4359",
		Amount:       20693.87,
	},
	{
		PaymentRefID: "c5d31907-ca17-4ca6-9897-29a70e48a573",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST4017",
		Amount:       33856.77,
	},
	{
		PaymentRefID: "7bc214ab-2274-4bc8-a6ab-25c4e88f8cee",
		Channel:      "ID_ALFAMART",
		PaymentCode:  "TEST4511",
		Amount:       3056.42,
	},
}

var INLINE_B_DATA = []Transaction{
	{
		PaymentRefID: "fae61b32-e941-4909-8df5-ed8e0a85953b",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST4359",
		Amount:       20693.87,
	},
	{
		PaymentRefID: "c5d31907-ca17-4ca6-9897-29a70e48a573",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST4017",
		Amount:       33856.77,
	},
	{
		PaymentRefID: "3f48ac25-4964-4e5f-8703-82508999dbbd",
		Channel:      "ID_ALFAMART",
		PaymentCode:  "TEST3472",
		Amount:       13345.4,
	},
	{
		PaymentRefID: "55fd4e3e-0c93-4f9d-aaef-89d7b13e963c",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST3523",
		Amount:       39247.7,
	},
	{
		PaymentRefID: "819759c3-cd2a-431a-a5f8-66e22ced49a1",
		Channel:      "ID_ALFAMART",
		PaymentCode:  "TEST33729",
		Amount:       40839.14,
	},
	{
		PaymentRefID: "8870f4e3-689a-48bd-8531-68cb5ff75b15",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST3466",
		Amount:       32587.54,
	},
	{
		PaymentRefID: "ff7c86d1-352e-47e3-a9e2-f83815a90870",
		Channel:      "ID_ALFAMART",
		PaymentCode:  "TEST3866",
		Amount:       24139.05,
	},
	{
		PaymentRefID: "268d743d-48bf-4ef1-a959-02cd822d1b28",
		Channel:      "ID_ALFAMART",
		PaymentCode:  "TEST4901",
		Amount:       25291.44,
	},
	{
		PaymentRefID: "0780433d-67f7-460c-8b34-39542aab4f7f",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST4759",
		Amount:       12282.1,
	},
	{
		PaymentRefID: "55735054-9873-4ff1-a65b-df9cb4ecb0b7",
		Channel:      "ID_ALFAMART",
		PaymentCode:  "TEST4641",
		Amount:       31971.37,
	},
	{
		PaymentRefID: "78aff9c3-1ac5-4c58-8f50-19dfd7cc64c9",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST4194",
		Amount:       2222.45,
	},
	{
		PaymentRefID: "cc159e81-d531-41bb-8491-a956df8e7cb8",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST3505",
		Amount:       45158.05,
	},
	{
		PaymentRefID: "5369e1ac-dfaf-43ac-a340-75cdeb2f7f53",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST3590",
		Amount:       38659.22,
	},
	{
		PaymentRefID: "8345869e-1bb2-4c90-a72a-aef53c33bda8",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST3348",
		Amount:       26149.49,
	},
	{
		PaymentRefID: "64d0f980-c8df-4939-9fdc-5aeb67429b1b",
		Channel:      "ID_ALFAMART",
		PaymentCode:  "TEST4540",
		Amount:       33016.57,
	},
	{
		PaymentRefID: "1d7e9004-d0a0-4d1e-97be-bb41a75c79f8",
		Channel:      "ID_ALFAMART",
		PaymentCode:  "TEST4335",
		Amount:       16025.97,
	},
	{
		PaymentRefID: "c7a902e2-01af-437f-a6a6-894d79f96b00",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST3627",
		Amount:       46957.63,
	},
	{
		PaymentRefID: "d897c533-494912-4d67-9f78-9ae7f7430c30",
		Channel:      "ID_ALFAMART",
		PaymentCode:  "TEST48138",
		Amount:       42586.39,
	},
	{
		PaymentRefID: "f59cf91b-d6b2-4def-bc14-0c4cc4befa63",
		Channel:      "ID_INDOMARET",
		PaymentCode:  "TEST4491",
		Amount:       18554.93,
	},
	{
		PaymentRefID: "7bc214ab-2274-4bc8-a6ab-25c4e88f8cee",
		Channel:      "ID_ALFAMART",
		PaymentCode:  "TEST4511",
		Amount:       3056.42,
	},
}
