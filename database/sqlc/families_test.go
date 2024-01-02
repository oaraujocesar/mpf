package database

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/oaraujocesar/mpf/util"
	"github.com/stretchr/testify/require"
)

func createRandomFamily(t *testing.T) Family {
	user := createRandomUser(t)

	arg := CreateFamilyParams{
		Name:   util.RandomName(),
		UserID: user.ID,
	}

	family, err := testQueries.CreateFamily(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, family)
	require.Equal(t, arg.Name, family.Name)
	require.Equal(t, arg.UserID, family.UserID)
	require.NotZero(t, family.ID)

	return family
}

func TestCreateFamily(t *testing.T) {
	setupTest(migrations)
	createRandomFamily(t)
	teardownTest(migrations)
}

func TestGetFamilyById(t *testing.T) {
	setupTest(migrations)
	family := createRandomFamily(t)

	family2, err := testQueries.GetFamilyById(context.Background(), family.ID)
	require.NoError(t, err)
	require.NotEmpty(t, family2)
	require.Equal(t, family.ID, family2.ID)
	require.Equal(t, family.Name, family2.Name)
	require.Equal(t, family.UserID, family2.UserID)
	teardownTest(migrations)
}

func TestUpdateFamily(t *testing.T) {
	setupTest(migrations)
	family := createRandomFamily(t)

	arg := UpdateFamilyParams{
		ID:   family.ID,
		Name: util.RandomName(),
	}

	family2, err := testQueries.UpdateFamily(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, family2)
	require.Equal(t, arg.ID, family2.ID)
	require.Equal(t, arg.Name, family2.Name)
	require.Equal(t, family.CreatedAt, family2.CreatedAt, time.Second)
	require.WithinDuration(t, family.UpdatedAt, family2.UpdatedAt, time.Second)
	teardownTest(migrations)
}

func TestListFamilies(t *testing.T) {
	setupTest(migrations)
	for i := 0; i < 10; i++ {
		createRandomFamily(t)
	}

	arg := ListFamiliesParams{
		Limit:  10,
		Offset: 0,
	}

	families, err := testQueries.ListFamilies(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, families, 10)

	for _, family := range families {
		require.NotEmpty(t, family)
		require.NotZero(t, family.ID)
		require.NotZero(t, family.UserID)
		require.NotZero(t, family.CreatedAt)
		require.NotZero(t, family.UpdatedAt)
	}

	teardownTest(migrations)
}

func TestDeleteFamily(t *testing.T) {
	setupTest(migrations)
	family := createRandomFamily(t)

	err := testQueries.DeleteFamily(context.Background(), family.ID)
	require.NoError(t, err)

	family2, err := testQueries.GetFamilyById(context.Background(), family.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, family2)
	teardownTest(migrations)
}