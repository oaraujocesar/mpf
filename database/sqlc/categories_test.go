package database

import (
	"context"
	"database/sql"
	"testing"

	"github.com/oaraujocesar/mpf/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	categoryName := util.RandomName()
	category, err := testQueries.CreateCategory(context.Background(), categoryName)

	require.NoError(t, err)
	require.NotEmpty(t, category)
	require.Equal(t, categoryName, category.Name)
	require.NotZero(t, category.ID)
	require.NotZero(t, category.CreatedAt)
	require.NotZero(t, category.UpdatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	setupTest(migrations)
	category := createRandomCategory(t)

	require.NotEmpty(t, category)
	teardownTest(migrations)
}

func TestGetCategoryById(t *testing.T) {
	setupTest(migrations)
	category := createRandomCategory(t)

	category2, err := testQueries.GetCategoryById(context.Background(), category.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category.ID, category2.ID)
	require.Equal(t, category.Name, category2.Name)
	require.Equal(t, category.CreatedAt, category2.CreatedAt)
	require.Equal(t, category.UpdatedAt, category2.UpdatedAt)

	teardownTest(migrations)
}

func TestSoftDeleteCategory(t *testing.T) {
	setupTest(migrations)
	category := createRandomCategory(t)

	err := testQueries.DeleteCategory(context.Background(), category.ID)
	require.NoError(t, err)

	category2, err := testQueries.GetCategoryById(context.Background(), category.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, category2)
	require.NotNil(t, category2.DeletedAt)

	teardownTest(migrations)
}

func TestListCategories(t *testing.T) {
	setupTest(migrations)
	for i := 0; i < 10; i++ {
		createRandomCategory(t)
	}

	args := ListCategoriesParams{
		Limit:  10,
		Offset: 0,
	}

	categories, err := testQueries.ListCategories(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, categories)
	require.Len(t, categories, 10)
	for _, category := range categories {
		require.NotEmpty(t, category)
		require.NotZero(t, category.ID)
		require.NotZero(t, category.CreatedAt)
		require.NotZero(t, category.UpdatedAt)
	}

	teardownTest(migrations)
}

func TestUpdateCategory(t *testing.T) {
	setupTest(migrations)
	category := createRandomCategory(t)

	newCategoryName := util.RandomName()

	arg := UpdateCategoryParams{
		ID:   category.ID,
		Name: newCategoryName,
	}

	category2, err := testQueries.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category.ID, category2.ID)
	require.NotEqual(t, category.Name, category2.Name)
	require.Equal(t, newCategoryName, category2.Name)
	require.Equal(t, category.CreatedAt, category2.CreatedAt)
	require.NotEqual(t, category.UpdatedAt, category2.UpdatedAt)
	teardownTest(migrations)
}
