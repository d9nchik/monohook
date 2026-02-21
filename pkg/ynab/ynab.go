package ynab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const defaultBaseURL = "https://api.ynab.com/v1"

type Client struct {
	baseURL    string
	ynabToken  string
	httpClient *http.Client
}

func NewClient(ynabToken string) *Client {
	return &Client{
		baseURL:    defaultBaseURL,
		ynabToken:  ynabToken,
		httpClient: &http.Client{},
	}
}

func (c *Client) CreateTransaction(budgetID string, transaction *Transaction) error {
	url := c.baseURL + fmt.Sprintf("/budgets/%s/transactions", budgetID)

	transactionReq := CreateTransactionReq{Transaction: *transaction}
	body, err := json.Marshal(transactionReq)
	if err != nil {
		return fmt.Errorf("failed to marshal transaction request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ynabToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
