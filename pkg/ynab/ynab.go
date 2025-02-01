package ynab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://api.ynab.com/v1"

type Client struct {
	ynabToken  string
	httpClient *http.Client
}

func NewClient(ynabToken string) *Client {
	return &Client{
		ynabToken:  ynabToken,
		httpClient: &http.Client{},
	}
}

func (c *Client) CreateTransaction(budgetID string, transaction *Transaction) error {
	url := baseURL + fmt.Sprintf("/budgets/%s/transactions", budgetID)
	method := "POST"

	transactionReq := CreateTransactionReq{Transaction: *transaction}
	body, err := json.Marshal(transactionReq)
	if err != nil {
		return fmt.Errorf("failed to marshal transaction request body: %w", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ynabToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create request: %d (%s)", resp.StatusCode, string(body))
	}

	// no need for answer for now

	return nil
}
