package database

import (
	"context"
	"database/sql"
	"testing"

	"github.com/oaraujocesar/mpf/util"
	"github.com/stretchr/testify/require"
)

func createRandomCard(t *testing.T, user User, family Family, tx *sql.Tx) Card {
	arg := CreateCardParams{
		Name:      util.RandomName(),
		CardLimit: util.RandomMoney(),
		DueDate:   int32(util.RandomInt(1, 28)),
		FamilyID:  sql.NullInt64{Int64: family.ID, Valid: true},
		UserID:    user.ID,
	}

	card, err := testStore.WithTx(tx).CreateCard(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, card)

	require.Equal(t, arg.Name, card.Name)
	require.Equal(t, arg.CardLimit, card.CardLimit)
	require.Equal(t, arg.DueDate, card.DueDate)
	require.Equal(t, arg.UserID, card.UserID)

	require.NotZero(t, card.ID)
	require.NotZero(t, card.CreatedAt)
	require.NotZero(t, card.UpdatedAt)
	require.Zero(t, card.DeletedAt)

	require.Equal(t, arg.Name, card.Name)
	require.Equal(t, user.ID, card.UserID)
	require.Equal(t, family.ID, card.FamilyID.Int64)
	require.InDelta(t, 1.00, card.CardLimit, 5000.00)
	require.InDelta(t, 1, card.DueDate, 28)

	return card
}

func TestCreateCard(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)

	user := createRandomUser(t, tx)
	family := createRandomFamily(t, tx)

	createRandomCard(t, user, family, tx)

	tx.Rollback()
}

func TestGetCard(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)

	user := createRandomUser(t, tx)
	family := createRandomFamily(t, tx)

	card1 := createRandomCard(t, user, family, tx)
	card2, err := testStore.WithTx(tx).GetCardById(context.Background(), card1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, card2)

	require.Equal(t, card1.ID, card2.ID)
	require.Equal(t, card1.Name, card2.Name)
	require.Equal(t, card1.UserID, card2.UserID)
	require.Equal(t, card1.FamilyID, card2.FamilyID)
	require.Equal(t, card1.CardLimit, card2.CardLimit)
	require.Equal(t, card1.DueDate, card2.DueDate)

	tx.Rollback()
}

func TestListCards(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)

	user := createRandomUser(t, tx)
	family := createRandomFamily(t, tx)

	for i := 0; i < 10; i++ {
		createRandomCard(t, user, family, tx)
	}

	arg := ListCardsParams{
		Limit:  5,
		Offset: 5,
	}

	cards, err := testStore.WithTx(tx).ListCards(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, cards, 5)

	for _, card := range cards {
		require.NotEmpty(t, card)
		require.NotZero(t, card.ID)
		require.NotZero(t, card.CreatedAt)
		require.NotZero(t, card.UpdatedAt)
		require.NotZero(t, card.DueDate)
	}

	tx.Rollback()
}

func TestUpdateCard(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)

	user := createRandomUser(t, tx)
	family := createRandomFamily(t, tx)

	card1 := createRandomCard(t, user, family, tx)

	arg := UpdateCardParams{
		ID:        card1.ID,
		Name:      util.RandomName(),
		CardLimit: util.RandomMoney(),
		DueDate:   int32(util.RandomInt(1, 28)),
	}

	card2, err := testStore.WithTx(tx).UpdateCard(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, card2)

	require.Equal(t, card1.ID, card2.ID)
	require.Equal(t, arg.Name, card2.Name)
	require.Equal(t, arg.CardLimit, card2.CardLimit)
	require.Equal(t, arg.DueDate, card2.DueDate)

	tx.Rollback()
}

func TestSoftDeleteCard(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)

	user := createRandomUser(t, tx)
	family := createRandomFamily(t, tx)

	card := createRandomCard(t, user, family, tx)

	err := testStore.WithTx(tx).DeleteCard(context.Background(), card.ID)
	require.NoError(t, err)

	card, err = testStore.WithTx(tx).GetCardById(context.Background(), card.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, card)

	tx.Rollback()
}
