package ynab

//	{
//	   "transaction": {
//	       "account_id": "06f76742-0bd9-4cfc-ade8-c5e7212317cb",
//	       "date": "2025-02-01",
//	       "amount": -350000,
//	       "payee_name": "non now",
//	       "memo": "123",
//	       "cleared": "cleared",
//	       "import_id": "2qaz"
//	   }
//	}

type TransactionClearedType string

const (
	TransactionCleared    = "cleared"
	TransactionUncleared  = "uncleared"
	TransactionReconciled = "reconciled"
)

type Transaction struct {
	AccountID string                 `json:"account_id"`
	Date      string                 `json:"date"`
	Amount    int64                  `json:"amount"`
	PayeeName *string                `json:"payee_name,omitempty"`
	Memo      *string                `json:"memo,omitempty"`
	Cleared   TransactionClearedType `json:"cleared"`
	ImportID  *string                `json:"import_id,omitempty"`
}

type CreateTransactionReq struct {
	Transaction Transaction `json:"transaction"`
}
