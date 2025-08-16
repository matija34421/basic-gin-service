package repository

import (
	"basic-gin/internal/model"
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

func (r *AccountRepository) GetAccounts() ([]*model.Account, error) {
	var accounts []*model.Account

	rows, err := r.db.Query("select  id,client_id,account_number,balance,created_at from accounts")

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

func (r *AccountRepository) GetAccountById(id int) (*model.Account, error) {
	var account model.Account

	row := r.db.QueryRow("select id,client_id,account_number,balance,created_at from accounts where id=$1", id)

	if err := row.Scan(&account.Id, &account.ClientId, &account.AccountNumber, &account.Balance, &account.Created_at); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("couldnt find account with an id: %d", id)
		}
		return nil, fmt.Errorf("error scanning the row: %w", err)
	}

	return &account, nil
}

func (r *AccountRepository) CreateAccount(account *model.Account) (*model.Account, error) {
	if err := r.db.QueryRow(`insert into accounts(id,client_id,account_number,balance,created_at) values($1,$2,$3,$4,$5) returning id`, account.Id, account.ClientId, account.AccountNumber, account.Balance, account.Created_at).Scan(&account.Id); err != nil {
		return nil, fmt.Errorf("couldnt save an account: %w", err)
	}

	return account, nil
}

func (r *AccountRepository) UpdateAccount(account *model.Account) (*model.Account, error) {
	var updatedAccount model.Account

	if err := r.db.QueryRow("update accounts set account_number = $1,balance = $2 where id = $3 returning id,client_id,account_number,balance,created_at", &account.AccountNumber, &account.Balance, &account.Id).Scan(&updatedAccount.Id, &updatedAccount.ClientId, &updatedAccount.AccountNumber, &updatedAccount.Balance, &updatedAccount.Created_at); err != nil {
		return nil, fmt.Errorf("couldnt update an account: %w", err)
	}

	return &updatedAccount, nil
}
