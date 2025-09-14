package repository

import (
	"basic-gin/internal/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) GetAccountsByClientId(ctx context.Context, clientId int) ([]*model.Account, error) {
	var accounts []*model.Account

	rows, err := r.db.QueryContext(ctx, "select id,client_id,account_number,balance,created_at from accounts where client_id = $1", clientId)

	if err != nil {
		return nil, fmt.Errorf("couldnt retrieve accounts %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var account model.Account
		if err = rows.Scan(&account.Id, &account.ClientId, &account.AccountNumber, &account.Balance, &account.Created_at); err != nil {
			return nil, fmt.Errorf("error scanning the row: %w", err)
		}

		accounts = append(accounts, &account)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row error: %w", err)
	}

	return accounts, nil
}

func (r *AccountRepository) GetAccountById(ctx context.Context, id int) (*model.Account, error) {
	var account model.Account

	row := r.db.QueryRowContext(ctx, "select id,client_id,account_number,balance,created_at from accounts where id=$1", id)

	if err := row.Scan(&account.Id, &account.ClientId, &account.AccountNumber, &account.Balance, &account.Created_at); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("couldnt find account with an id: %d", id)
		}
		return nil, fmt.Errorf("error scanning the row: %w", err)
	}

	return &account, nil
}

func (r *AccountRepository) CreateAccount(ctx context.Context, account *model.Account) (*model.Account, error) {
	if err := r.db.QueryRowContext(
		ctx,
		`INSERT INTO accounts (client_id, account_number, balance, created_at) 
		 VALUES ($1, $2, $3, $4) 
		 RETURNING id`,
		account.ClientId, account.AccountNumber, account.Balance, account.Created_at,
	).Scan(&account.Id); err != nil {
		return nil, fmt.Errorf("couldn't save an account: %w", err)
	}

	return account, nil
}

func (r *AccountRepository) UpdateAccount(ctx context.Context, account *model.Account) (*model.Account, error) {
	var updatedAccount model.Account

	if err := r.db.QueryRowContext(ctx, "update accounts set balance = $1 where id = $2 returning id,client_id,account_number,balance,created_at", account.Balance, account.Id).Scan(&updatedAccount.Id, &updatedAccount.ClientId, &updatedAccount.AccountNumber, &updatedAccount.Balance, &updatedAccount.Created_at); err != nil {
		return nil, fmt.Errorf("couldnt update an account: %w", err)
	}

	return &updatedAccount, nil
}
