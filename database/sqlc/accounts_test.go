package database

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/oaraujocesar/mpf/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T, tx *sql.Tx) Account {
	user := createRandomUser(t, tx)

	arg := CreateAccountParams{
		Balance: float64(util.RandomMoney()),
		Name:    util.RandomName(),
		UserID:  user.ID,
	}

	account, err := testStore.WithTx(tx).CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Name, account.Name)
	require.Equal(t, arg.UserID, account.UserID)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.UpdatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)
	account := createRandomAccount(t, tx)

	require.NotEmpty(t, account)

	tx.Rollback()
}

func TestGetAccountById(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)
	account := createRandomAccount(t, tx)

	account2, err := testStore.WithTx(tx).GetAccountById(context.Background(), account.ID)
	if err != nil {
		t.Fatal(err)
	}

	require.NotEmpty(t, account2)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Balance, account2.Balance)
	require.Equal(t, account.Name, account2.Name)
	require.Equal(t, account.UserID, account2.UserID)
	tx.Rollback()
}

func TestSoftDeleteAccount(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)
	account := createRandomAccount(t, tx)

	err := testStore.WithTx(tx).DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account2, err := testStore.WithTx(tx).GetAccountById(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
	require.NotNil(t, account2.DeletedAt)
	tx.Rollback()
}

func TestListAccounts(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)

	for i := 0; i < 10; i++ {
		createRandomAccount(t, tx)
	}

	arg := ListAccountsParams{
		Limit:  10,
		Offset: 0,
	}

	accounts, err := testStore.WithTx(tx).ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 10)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.NotZero(t, account.ID)
		require.NotZero(t, account.Balance)
		require.NotEmpty(t, account.Name)
	}
	tx.Rollback()
}

func TestUpdateAccount(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)
	account := createRandomAccount(t, tx)

	arg := UpdateAccountParams{
		Balance: float64(util.RandomMoney()),
		Name:    util.RandomName(),
		ID:      account.ID,
	}

	account2, err := testStore.WithTx(tx).UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, arg.Name, account2.Name)
	require.Equal(t, account.UserID, account2.UserID)
	require.NotEqual(t, account.Balance, account2.Balance)
	require.WithinDuration(t, account.UpdatedAt, account2.UpdatedAt, time.Second)
	tx.Rollback()
}

func TestUpdateBalance(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)
	account := createRandomAccount(t, tx)

	arg := UpdateBalanceParams{
		Balance: account.Balance + float64(util.RandomMoney()),
		ID:      account.ID,
	}

	account2, err := testStore.WithTx(tx).UpdateBalance(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account.Name, account2.Name)
	require.Equal(t, account.UserID, account2.UserID)
	require.NotEqual(t, account.Balance, account2.Balance)
	require.WithinDuration(t, account.UpdatedAt, account2.UpdatedAt, time.Second)
	tx.Rollback()
}
