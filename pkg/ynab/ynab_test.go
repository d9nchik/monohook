package ynab

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func newTestClient(t *testing.T, handler http.HandlerFunc) (*Client, *httptest.Server) {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	return &Client{
		baseURL:    server.URL,
		ynabToken:  "test-token",
		httpClient: server.Client(),
	}, server
}

func TestClient_CreateTransaction(t *testing.T) {
	var receivedReq CreateTransactionReq
	var receivedAuth string

	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		receivedAuth = r.Header.Get("Authorization")
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/budgets/budget-1/transactions", r.URL.Path)

		err := json.NewDecoder(r.Body).Decode(&receivedReq)
		require.NoError(t, err)
		w.WriteHeader(http.StatusCreated)
	})

	memo := "test memo"
	payee := "test payee"
	importID := "import-123"

	tx := &Transaction{
		AccountID: "acc-1",
		Date:      "2025-01-01",
		Amount:    -10000,
		Cleared:   TransactionCleared,
		PayeeName: &payee,
		Memo:      &memo,
		ImportID:  &importID,
	}

	err := client.CreateTransaction("budget-1", tx)
	require.NoError(t, err)
	require.Equal(t, "Bearer test-token", receivedAuth)
	require.Equal(t, "acc-1", receivedReq.Transaction.AccountID)
	require.Equal(t, int64(-10000), receivedReq.Transaction.Amount)
}

func TestClient_CreateTransaction_ServerError(t *testing.T) {
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad request"}`))
	})

	err := client.CreateTransaction("budget-1", &Transaction{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "400")
}
