package database

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomInvoice(t *testing.T, tx *sql.Tx, card Card, account Account) Invoice {
	arg := CreateInvoiceParams{
		CardID:    card.ID,
		Amount:    100.00,
		AccountID: account.ID,
		CloseAt:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		DueAt:     time.Date(2023, 1, 7, 0, 0, 0, 0, time.UTC),
	}

	invoice, err := testStore.WithTx(tx).CreateInvoice(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, invoice)

	require.Equal(t, arg.CardID, invoice.CardID)
	require.Equal(t, arg.Amount, invoice.Amount)
	require.Equal(t, arg.AccountID, invoice.AccountID)
	require.WithinDuration(t, arg.CloseAt, invoice.CloseAt, time.Second)
	require.WithinDuration(t, arg.DueAt, invoice.DueAt, time.Second)

	return invoice
}

func TestCreateInvoice(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()

	user := createRandomUser(t, tx)
	family := createRandomFamily(t, tx)
	account := createRandomAccount(t, tx)
	card := createRandomCard(t, user, family, tx)

	createRandomInvoice(t, tx, card, account)
}

func TestGetInvoiceById(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()

	user := createRandomUser(t, tx)
	family := createRandomFamily(t, tx)
	account := createRandomAccount(t, tx)
	card := createRandomCard(t, user, family, tx)

	invoice := createRandomInvoice(t, tx, card, account)

	invoiceFound, err := testStore.WithTx(tx).GetInvoiceById(context.Background(), invoice.ID)
	require.NoError(t, err)
	require.NotEmpty(t, invoiceFound)

	require.Equal(t, invoice.ID, invoiceFound.ID)
	require.Equal(t, invoice.CardID, invoiceFound.CardID)
	require.Equal(t, invoice.Amount, invoiceFound.Amount)
	require.Equal(t, invoice.AccountID, invoiceFound.AccountID)
	require.WithinDuration(t, invoice.CloseAt, invoiceFound.CloseAt, time.Second)
	require.WithinDuration(t, invoice.DueAt, invoiceFound.DueAt, time.Second)
}

func TestListInvoices(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()

	user := createRandomUser(t, tx)
	family := createRandomFamily(t, tx)
	account := createRandomAccount(t, tx)
	card := createRandomCard(t, user, family, tx)

	for i := 0; i < 10; i++ {
		createRandomInvoice(t, tx, card, account)
	}

	arg := ListInvoicesParams{
		Limit:  5,
		Offset: 5,
	}

	invoices, err := testStore.WithTx(tx).ListInvoices(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, invoices, 5)

	for _, invoice := range invoices {
		require.NotEmpty(t, invoice)
	}
}

func TestUpdateInvoiceAmount(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()

	user := createRandomUser(t, tx)
	family := createRandomFamily(t, tx)
	account := createRandomAccount(t, tx)
	card := createRandomCard(t, user, family, tx)

	invoice := createRandomInvoice(t, tx, card, account)

	arg := UpdateInvoiceAmountParams{
		ID:     invoice.ID,
		Amount: 200.00,
	}

	invoice2, err := testStore.WithTx(tx).UpdateInvoiceAmount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, invoice2)

	require.Equal(t, invoice.ID, invoice2.ID)
	require.Equal(t, arg.Amount, invoice2.Amount)
	require.Equal(t, invoice.CardID, invoice2.CardID)
	require.Equal(t, invoice.AccountID, invoice2.AccountID)
	require.WithinDuration(t, invoice.CloseAt, invoice2.CloseAt, time.Second)
	require.WithinDuration(t, invoice.DueAt, invoice2.DueAt, time.Second)
	require.WithinDuration(t, invoice.UpdatedAt, invoice2.UpdatedAt, time.Second)
}

// TODO: continue invoice tests
