// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateAuthor(ctx context.Context, arg CreateAuthorParams) (User, error)
	CreateCategory(ctx context.Context, name string) (Category, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	DeleteAuthor(ctx context.Context, id pgtype.UUID) error
	DeleteCategory(ctx context.Context, id pgtype.UUID) error
	DeleteEntry(ctx context.Context, id pgtype.UUID) error
	GetAuthor(ctx context.Context, email string) (User, error)
	GetCategory(ctx context.Context, name string) (Category, error)
	GetEntry(ctx context.Context, id pgtype.UUID) (Entry, error)
	ListAuthors(ctx context.Context, arg ListAuthorsParams) ([]User, error)
	ListCategories(ctx context.Context, arg ListCategoriesParams) ([]Category, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (User, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
	UpdateEntry(ctx context.Context, arg UpdateEntryParams) (Entry, error)
}

var _ Querier = (*Queries)(nil)
