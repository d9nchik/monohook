package monobank

// ClientInfo - client/user info
// Personal API - https://api.monobank.ua/docs/#/definitions/UserInfo
type ClientInfo struct {
	ID         string   `json:"clientId"`
	Name       string   `json:"name"`
	WebHookURL string   `json:"webHookUrl"`
	Accounts   Accounts `json:"accounts"`
	Jars       Jars     `json:"jars"`
}

type Account struct {
	AccountID    string   `json:"id"`
	SendID       string   `json:"sendId"`
	Balance      int64    `json:"balance"`
	CreditLimit  int64    `json:"creditLimit"`
	CurrencyCode int      `json:"currencyCode"`
	CashbackType string   `json:"cashbackType"` // enum: None, UAH, Miles
	CardMasks    []string `json:"maskedPan"`    // card number masks
	Type         CardType `json:"type"`
	IBAN         string   `json:"iban"`
}

type Jar struct {
	ID           string `json:"id"`
	SendID       string `json:"sendId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	CurrencyCode int    `json:"currencyCode"`
	Balance      int64  `json:"balance"`
	Goal         int64  `json:"goal"`
}

type CardType string

const (
	Black    CardType = "black"    //
	White    CardType = "white"    //
	Platinum CardType = "platinum" //
	FOP      CardType = "fop"      // ФОП
	EAid     CardType = "eAid"     // єПідтримка
)

type Accounts []Account

type Jars []Jar

// Transaction - bank account statement
type Transaction struct {
	ID          string `json:"id"`
	Time        int64  `json:"time"`
	Description string `json:"description"`
	MCC         int32  `json:"mcc"`
	OriginalMCC int32  `json:"originalMcc"`
	Hold        bool   `json:"hold"`
	// Amount in the account currency
	Amount int64 `json:"amount"`
	// OperationAmount in the transaction currency or amount after double conversion
	OperationAmount int64 `json:"operationAmount"`
	// ISO 4217 numeric code
	CurrencyCode   int32  `json:"currencyCode"`
	CommissionRate int64  `json:"commissionRate"`
	CashbackAmount int64  `json:"cashbackAmount"`
	Balance        int64  `json:"balance"`
	Comment        string `json:"comment"`
	// For withdrawal only.
	ReceiptID string `json:"receiptId"`
	// For fop(ФОП) accounts only.
	InvoiceID string `json:"invoiceId"`
	// For fop(ФОП) accounts only.
	EDRPOU string `json:"counterEdrpou"`
	// For fop(ФОП) accounts only.
	IBAN string `json:"counterIban"`
}

// Transactions - transactions
type Transactions []Transaction

type Currency struct {
	CurrencyCodeA int     `json:"currencyCodeA"`
	CurrencyCodeB int     `json:"currencyCodeB"`
	Date          int64   `json:"date"`
	RateSell      float64 `json:"rateSell"`
	RateBuy       float64 `json:"rateBuy"`
	RateCross     float64 `json:"rateCross"`
}

type Currencies []Currency

type WebHookRequest struct {
	WebHookURL string `json:"webHookUrl"`
}

type WebHookResponse struct {
	Type string      `json:"type"` // "StatementItem"
	Data WebHookData `json:"data"`
}

type WebHookData struct {
	AccountID   string      `json:"account"`
	Transaction Transaction `json:"statementItem"`
}
