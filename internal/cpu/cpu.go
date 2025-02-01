package cpu

import (
	"fmt"
	"monoHook/pkg/monobank"
	"monoHook/pkg/ynab"
	"strings"
	"time"
)

type CPU struct {
	ynabClient *ynab.Client
	budgetId   string
	accountId  string
}

func NewCPU(ynabClient *ynab.Client, budgetId, accountId string) *CPU {
	return &CPU{ynabClient: ynabClient, budgetId: budgetId, accountId: accountId}
}

func (c *CPU) AddTransaction(t monobank.Transaction) error {
	importId := stringPtr(shortenString(t.ID, 36))

	memo, payeeName := getMemoAndPayeeName(t.Description, t.Comment)

	err := c.ynabClient.CreateTransaction(c.budgetId, &ynab.Transaction{
		AccountID: c.accountId,
		Date:      time.Now().Format(time.DateOnly),
		Amount:    t.Amount * 10,
		Cleared:   ynab.TransactionCleared,
		PayeeName: payeeName,
		Memo:      memo,
		ImportID:  importId,
	})
	if err != nil {
		return fmt.Errorf("couldn't create transaction, %w", err)
	}

	return nil
}

func getMemoAndPayeeName(description, comment string) (memo, payeeName *string) {
	descriptionParts := strings.Split(description, ":")
	switch {
	case len(descriptionParts) == 2:
		payeeName = stringPtr(strings.TrimSpace(shortenString(descriptionParts[1], 50)))
		if comment != "" {
			memo = stringPtr(strings.TrimSpace(shortenString(comment, 200)))
		}
		return
	case comment != "":
		memo = stringPtr(strings.TrimSpace(shortenString(fmt.Sprintf("%s Comment: %s", description, comment), 200)))
		return
	default:
		memo = stringPtr(strings.TrimSpace(shortenString(description, 200)))
		return
	}
}

func shortenString(str string, maxLen int) string {
	var result string
	counter := 0
	for _, symbol := range str {
		if counter > maxLen {
			break
		}
		result += string(symbol)
		counter++
	}
	return str
}

func stringPtr(str string) *string {
	return &str
}
