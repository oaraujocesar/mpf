package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomMember(t *testing.T, family Family, user User) Member {
	arg := CreateMemberParams{
		FamilyID: family.ID,
		UserID:   user.ID,
	}

	member, err := testQueries.CreateMember(context.Background(), arg)
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
	setupTest(migrations)

	user := createRandomUser(t)
	family := createRandomFamily(t)

	member, err := testQueries.CreateMember(context.Background(), CreateMemberParams{FamilyID: family.ID, UserID: user.ID})
	require.NoError(t, err)
	require.NotEmpty(t, member)
	require.Equal(t, user.ID, member.UserID)
	require.NotZero(t, family.ID, member.FamilyID)
	require.NotZero(t, member.CreatedAt)
	require.NotZero(t, member.UpdatedAt)

	teardownTest(migrations)
}

func TestDeleteMember(t *testing.T) {
	setupTest(migrations)

	user := createRandomUser(t)
	family := createRandomFamily(t)

	member, err := testQueries.CreateMember(context.Background(), CreateMemberParams{FamilyID: family.ID, UserID: user.ID})
	require.NoError(t, err)

	err = testQueries.DeleteMember(context.Background(), member.ID)
	require.NoError(t, err)

	member, err = testQueries.GetMemberById(context.Background(), member.ID)
	require.Error(t, err)
	require.Empty(t, member)

	teardownTest(migrations)
}

func TestListMembers(t *testing.T) {
	setupTest(migrations)

	user := createRandomUser(t)
	family := createRandomFamily(t)

	for i := 0; i < 10; i++ {
		createRandomMember(t, family, user)
	}

	arg := ListMembersParams{
		Limit:  5,
		Offset: 5,
	}

	members, err := testQueries.ListMembers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, members, 5)

	for _, member := range members {
		require.NotEmpty(t, member)
	}

	teardownTest(migrations)
}
