// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateCategories(ctx context.Context, arg CreateCategoriesParams) (Category, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccounts(ctx context.Context, id int32) error
	DeleteCategories(ctx context.Context, id int32) error
	GetAccount(ctx context.Context, id int32) (Account, error)
	GetAccounts(ctx context.Context, arg GetAccountsParams) ([]GetAccountsRow, error)
	GetAccountsGraph(ctx context.Context, arg GetAccountsGraphParams) (int64, error)
	GetAccountsReports(ctx context.Context, arg GetAccountsReportsParams) (int64, error)
	GetCategories(ctx context.Context, arg GetCategoriesParams) ([]Category, error)
	GetCategoriesById(ctx context.Context, id int32) (Category, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserById(ctx context.Context, id int32) (User, error)
	UpdateAccounts(ctx context.Context, arg UpdateAccountsParams) (Account, error)
	UpdateCategories(ctx context.Context, arg UpdateCategoriesParams) (Category, error)
}

var _ Querier = (*Queries)(nil)
