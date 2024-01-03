package database

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/oaraujocesar/mpf/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T, tx *sql.Tx) User {
	arg := CreateUserParams{
		Name:     util.RandomName(),
		Email:    util.RandomEmail(),
		Password: util.RandomPassword(10),
		Avatar:   sql.NullString{String: "https://cdn.pixabay.com/photo/2016/08/08/09/17/avatar-1577909_1280.png", Valid: true},
	}

	user, err := testStore.WithTx(tx).CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	// TODO: hash password
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Avatar, user.Avatar)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)

	createRandomUser(t, tx)

	tx.Rollback()
}

func TestGetUserByEmail(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)

	user1 := createRandomUser(t, tx)
	user2, err := testStore.WithTx(tx).GetUserByEmail(context.Background(), user1.Email)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	// TODO: hash password
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Avatar, user2.Avatar)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)

	tx.Rollback()
}

func TestGetUserByID(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)

	user1 := createRandomUser(t, tx)
	user2, err := testStore.WithTx(tx).GetUserById(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)

	tx.Rollback()
}

func TestUpdateUser(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)

	user1 := createRandomUser(t, tx)

	arg := UpdateUserParams{
		ID:       user1.ID,
		Name:     util.RandomName(),
		Password: util.RandomPassword(10),
		Email:    util.RandomEmail(),
		Avatar:   sql.NullString{String: util.RandomString(20), Valid: true},
	}

	user2, err := testStore.WithTx(tx).UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.Name, user2.Name)
	require.Equal(t, arg.Email, user2.Email)
	require.Equal(t, arg.Password, user2.Password)
	require.Equal(t, arg.Avatar, user2.Avatar)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)

	tx.Rollback()
}

func TestListUsers(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)

	for i := 0; i < 10; i++ {
		createRandomUser(t, tx)
	}

	arg := ListUsersParams{
		Limit:  10,
		Offset: 0,
	}

	users, err := testStore.WithTx(tx).ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 10)

	for _, user := range users {
		require.NotEmpty(t, user)
		require.NotZero(t, user.ID)
		require.NotZero(t, user.CreatedAt)
		require.NotZero(t, user.UpdatedAt)
	}

	tx.Rollback()
}

func TestSoftDeleteUser(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)

	user1 := createRandomUser(t, tx)

	err := testStore.WithTx(tx).DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testStore.WithTx(tx).GetUserById(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
	require.NotNil(t, user2.DeletedAt)

	tx.Rollback()
}
