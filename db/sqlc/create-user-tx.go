package db

import (
	"context"
)

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error
}

type CreateUserTxResult struct {
	User User
}

func (store *SQLStore) CreateUserTransaction(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.execTx(ctx, func(queries *Queries) error {
		var err error

		result.User, err = queries.CreateUser(ctx, arg.CreateUserParams)

		if err != nil {
			return err
		}
		return arg.AfterCreate(result.User)
	})

	return result, err
}
