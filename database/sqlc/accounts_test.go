package database

import (
	"context"
	"database/sql"
	"testing"

	"github.com/oaraujocesar/mpf/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Balance: float64(util.RandomMoney()),
		Name:    util.RandomName(),
		UserID:  user.ID,
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	return account
}

func TestCreateAccount(t *testing.T) {
	account := createRandomAccount(t)

	require.NotEmpty(t, account)

	t.Cleanup(func() {
		err := testQueries.DeleteAccount(context.Background(), account.ID)
		require.NoError(t, err)
	})
}

func TestGetAccountById(t *testing.T) {
	account := createRandomAccount(t)

	account2, err := testQueries.GetAccountById(context.Background(), account.ID)
	if err != nil {
		t.Fatal(err)
	}

	require.NotEmpty(t, account2)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Balance, account2.Balance)
	require.Equal(t, account.Name, account2.Name)
	require.Equal(t, account.UserID, account2.UserID)

	t.Cleanup(func() {
		err := testQueries.DeleteAccount(context.Background(), account.ID)
		require.NoError(t, err)
	})
}

func TestSoftDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccountById(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
	require.NotNil(t, account2.DeletedAt)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  10,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 10)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.NotZero(t, account.ID)
		require.NotZero(t, account.Balance)
		require.NotEmpty(t, account.Name)
	}
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	arg := UpdateAccountParams{
		Balance: float64(util.RandomMoney()),
		Name:    util.RandomName(),
		ID:      account.ID,
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, arg.Name, account2.Name)
	require.Equal(t, account.UserID, account2.UserID)
	require.NotEqual(t, account.Balance, account2.Balance)
	require.NotEqual(t, account.UpdatedAt, account2.UpdatedAt)
}

func TestUpdateBalance(t *testing.T) {
	account := createRandomAccount(t)

	arg := UpdateBalanceParams{
		Balance: account.Balance + float64(util.RandomMoney()),
		ID:      account.ID,
	}

	account2, err := testQueries.UpdateBalance(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account.Name, account2.Name)
	require.Equal(t, account.UserID, account2.UserID)
	require.NotEqual(t, account.Balance, account2.Balance)
	require.NotEqual(t, account.UpdatedAt, account2.UpdatedAt)
}
