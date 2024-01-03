package database

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/oaraujocesar/mpf/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T, tx *sql.Tx) Category {
	categoryName := util.RandomName()
	category, err := testStore.WithTx(tx).CreateCategory(context.Background(), categoryName)

	require.NoError(t, err)
	require.NotEmpty(t, category)
	require.Equal(t, categoryName, category.Name)
	require.NotZero(t, category.ID)
	require.NotZero(t, category.CreatedAt)
	require.NotZero(t, category.UpdatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)

	createRandomCategory(t, tx)

	tx.Rollback()
}

func TestGetCategoryById(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)

	category := createRandomCategory(t, tx)

	category2, err := testStore.WithTx(tx).GetCategoryById(context.Background(), category.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category.ID, category2.ID)
	require.Equal(t, category.Name, category2.Name)
	require.Equal(t, category.CreatedAt, category2.CreatedAt)
	require.Equal(t, category.UpdatedAt, category2.UpdatedAt)

	tx.Rollback()
}

func TestSoftDeleteCategory(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)

	category := createRandomCategory(t, tx)

	err := testStore.WithTx(tx).DeleteCategory(context.Background(), category.ID)
	require.NoError(t, err)

	category2, err := testStore.WithTx(tx).GetCategoryById(context.Background(), category.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, category2)
	require.NotNil(t, category2.DeletedAt)

	tx.Rollback()
}

func TestListCategories(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)

	for i := 0; i < 10; i++ {
		createRandomCategory(t, tx)
	}

	args := ListCategoriesParams{
		Limit:  10,
		Offset: 0,
	}

	categories, err := testStore.WithTx(tx).ListCategories(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, categories)
	require.Len(t, categories, 10)
	for _, category := range categories {
		require.NotEmpty(t, category)
		require.NotZero(t, category.ID)
		require.NotZero(t, category.CreatedAt)
		require.NotZero(t, category.UpdatedAt)
	}

	tx.Rollback()
}

func TestUpdateCategory(t *testing.T) {
	tx, _ := testStore.db.BeginTx(context.Background(), nil)

	category := createRandomCategory(t, tx)

	newCategoryName := util.RandomName()

	arg := UpdateCategoryParams{
		ID:   category.ID,
		Name: newCategoryName,
	}

	category2, err := testStore.WithTx(tx).UpdateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category.ID, category2.ID)
	require.NotEqual(t, category.Name, category2.Name)
	require.Equal(t, newCategoryName, category2.Name)
	require.Equal(t, category.CreatedAt, category2.CreatedAt)
	require.WithinDuration(t, category.UpdatedAt, category2.UpdatedAt, time.Second)

	tx.Rollback()
}
