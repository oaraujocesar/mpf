package database

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomMember(t *testing.T, family Family, user User, tx *sql.Tx) Member {
	arg := CreateMemberParams{
		FamilyID: family.ID,
		UserID:   user.ID,
	}

	member, err := testStore.WithTx(tx).CreateMember(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, member)

	require.Equal(t, arg.FamilyID, member.FamilyID)
	require.Equal(t, arg.UserID, member.UserID)

	require.NotZero(t, member.ID)
	require.NotZero(t, member.CreatedAt)
	require.NotZero(t, member.UpdatedAt)
	require.Zero(t, member.DeletedAt)

	return member
}

func TestCreateFamilyMember(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()

	user := createRandomUser(t, tx)
	user2 := createRandomUser(t, tx)
	family := createRandomFamily(t, tx, user2)

	member, err := testStore.WithTx(tx).CreateMember(context.Background(), CreateMemberParams{FamilyID: family.ID, UserID: user.ID})
	require.NoError(t, err)
	require.NotEmpty(t, member)
	require.Equal(t, user.ID, member.UserID)
	require.Equal(t, family.ID, member.FamilyID)
	require.NotZero(t, member.CreatedAt)
	require.NotZero(t, member.UpdatedAt)
}

func TestDeleteMember(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()

	user := createRandomUser(t, tx)
	family := createRandomFamily(t, tx, user)

	member := createRandomMember(t, family, createRandomUser(t, tx), tx)

	err := testStore.WithTx(tx).DeleteMember(context.Background(), member.ID)
	require.NoError(t, err)

	member, err = testStore.WithTx(tx).GetMemberById(context.Background(), member.ID)
	require.Error(t, err)
	require.Empty(t, member)
}

func TestListAppMembers(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()

	user := createRandomUser(t, tx)
	family := createRandomFamily(t, tx, user)

	for i := 0; i < 10; i++ {
		createRandomMember(t, family, createRandomUser(t, tx), tx)
	}

	arg := ListAppMembersParams{
		Limit:  5,
		Offset: 5,
	}

	members, err := testStore.WithTx(tx).ListAppMembers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, members, 5)

	for _, member := range members {
		require.NotEmpty(t, member)
		require.NotZero(t, member.UserID)
		require.Equal(t, family.ID, member.FamilyID)
		require.NotZero(t, member.CreatedAt)
		require.NotZero(t, member.UpdatedAt)
	}
}

func TestListFamilyMembers(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)
	defer tx.Rollback()

	user := createRandomUser(t, tx)
	family := createRandomFamily(t, tx, user)

	for i := 0; i < 10; i++ {
		createRandomMember(t, family, createRandomUser(t, tx), tx)
	}

	arg := ListFamilyMembersParams{
		FamilyID: family.ID,
		Limit:    5,
		Offset:   5,
	}

	members, err := testStore.WithTx(tx).ListFamilyMembers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, members, 5)

	for _, member := range members {
		require.NotEmpty(t, member)
		require.NotZero(t, member.UserID)
		require.Equal(t, family.ID, member.FamilyID)
		require.NotZero(t, member.CreatedAt)
		require.NotZero(t, member.UpdatedAt)
	}
}
