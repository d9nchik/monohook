package cpu

import (
	"fmt"
	"monoHook/pkg/monobank"
	"strings"
	"time"

	"github.com/brunomvsouza/ynab.go"
	"github.com/brunomvsouza/ynab.go/api"
	"github.com/brunomvsouza/ynab.go/api/transaction"
)

type CPU struct {
	ynabClient ynab.ClientServicer
	budgetId   string
	accountId  string
}

func NewCPU(ynabClient ynab.ClientServicer, budgetId, accountId string) *CPU {
	return &CPU{ynabClient: ynabClient, budgetId: budgetId, accountId: accountId}
}

func (c *CPU) AddTransaction(t monobank.Transaction) error {
	importId := stringPtr(shortenString(t.ID, 36))

	memo, payeeName := getMemoAndPayeeName(t.Description, t.Comment)

	_, err := c.ynabClient.Transaction().CreateTransaction(c.budgetId, transaction.PayloadTransaction{
		AccountID:  c.accountId,
		Date:       api.Date{Time: time.Now().UTC()},
		Amount:     t.Amount * 10,
		Cleared:    transaction.ClearingStatusCleared,
		Approved:   false,
		PayeeID:    nil,
		PayeeName:  payeeName,
		CategoryID: nil,
		Memo:       memo,
		FlagColor:  nil,
		ImportID:   importId,
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
	if len(str) > maxLen {
		return str[:maxLen]
	}
	return str
}

func stringPtr(str string) *string {
	return &str
}
