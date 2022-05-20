# xendit-trial (Transaction Audit)

## Overview

This is an application that will compare 2 sets of transaction data and generate the comparison results.

The app will accept an input of 2 sets of transactions with the following data:

| Field name       | Data type |
| ---------------- | --------- |
| `payment_ref_id` | String    |
| `channel`        | String    |
| `payment_code`   | String    |
| `amount`         | Decimal   |
| `timestamp`      | String    |
| `payer_name`     | String    |

And it will produce a comparison result like the following example data:

| `payment_ref_id`                         | `audit_result`         |
| ---------------------------------------- | ---------------------- |
| `8w84wjduw-b540-4d67-9f78-9ae7f7430c30`  | `MISSING_IN_B_DATA`    |
| `d897c533-494912-4d67-9f78-9ae7f7430c30` | `MISSING_IN_A_DATA`    |
| `f59cf91b-d6b2-4def-bc14-0c4cc4befa63`   | `MISMATCH_TRANSACTION` |
| `55735054-9873-4ff1-a65b-df9cb4ecb0b7`   | `MISMATCH_TRANSACTION` |
| `819759c3-cd2a-431a-a5f8-66e22ced49a1`   | `MISMATCH_TRANSACTION` |
| `3f48ac25-4964-4e5f-8703-82508999dbbd`   | `MISMATCH_TRANSACTION` |

The app will do the following if it gets 2 set of input, A and B:

- Compare a transaction in A and B when a transaction with the following fields are found on both data set:
  - `payment_ref_id`
  - `channel`
- Consider a transaction as _matching_ and won't produce an output if a transaction in A and B have **all of the following 4 fields matching**[^1]:
  - `payment_ref_id`
  - `channel`
  - `payment_code`
  - `amount`
- Output a `MISMATCH_TRANSACTION` result if a transaction exists in A and B but **any of the following 2 fields are not matching**:
  - `payment_code`
  - `amount`
- Output a `MISSING_IN_A_DATA` result if a transaction exists in A but not in B
- Output a `MISSING_IN_B_DATA` result if a transaction exists in B but not in A

[^1]: No output means that the app founds no mismatch or missing transactions.

This is the flow of how the app works:

```
 input --> split the task by the channel field --> compare the sets -> output
```

## Technical Spec

To access the API, you only need to send a `POST` request.

The input will be the following JSON object:

```json
{
  "a": [
    {
      "payment_ref_id": "string",
      "channel": "string",
      "payment_code": "string",
      "amount": 0,
      "timestamp": "string",
      "payer_name": "string"
    }
  ],
  "b": [
    {
      "payment_ref_id": "string",
      "channel": "string",
      "payment_code": "string",
      "amount": 0,
      "timestamp": "string",
      "payer_name": "string"
    }
  ]
}
```

And the result will be a JSON object with the following schema:

```json
[
  {
    "payment_ref_id": "string",
    "audit_result": "string"
  }
]
```
