package service

import (
	accountdto "basic-gin/internal/dto/accountDto"
	accountMapper "basic-gin/internal/mapper"
	"basic-gin/internal/model"
	"basic-gin/internal/repository"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

type AccountService struct {
	accountRepository *repository.AccountRepository
	accountMapper     *accountMapper.AccountMapper
}

func NewAccountService(repo *repository.AccountRepository) *AccountService {
	return &AccountService{
		accountRepository: repo,
		accountMapper:     &accountMapper.AccountMapper{},
	}
}

func (s *AccountService) GetAccountsByClientId(clientId int) ([]*accountdto.ResponseAccountDto, error) {
	if clientId < 0 {
		return nil, fmt.Errorf("id cant be negative number")
	}

	accounts, err := s.accountRepository.GetAccountsByClientId(clientId)

	if err != nil {
		return nil, fmt.Errorf("get accounts by client id: %w", err)
	}

	return s.accountMapper.ToResponseSlice(accounts), nil
}

func (s *AccountService) GetAccountById(id int) (*accountdto.ResponseAccountDto, error) {
	if id < 0 {
		return nil, fmt.Errorf("id cant be negative number")
	}

	account, err := s.accountRepository.GetAccountById(id)

	if err != nil {
		return nil, fmt.Errorf("get account by id: %w", err)
	}

	return s.accountMapper.ToResponse(account), nil
}

func (s *AccountService) CreateAccount(clientId int) (*accountdto.ResponseAccountDto, error) {
	if clientId < 0 {
		return nil, fmt.Errorf("invalid clientId")
	}

	account := createAccountShell(clientId)

	savedAccount, err := s.accountRepository.CreateAccount(&account)

	if err != nil {
		return nil, fmt.Errorf("create account: %w", err)
	}

	return s.accountMapper.ToResponse(savedAccount), nil
}

func (s *AccountService) UpdateAccount(updateDto accountdto.UpdateAccountDto) (*accountdto.ResponseAccountDto, error) {

	existingAccountDto, err := s.GetAccountById(updateDto.Id)

	if err != nil {
		return nil, fmt.Errorf("account not found: %w", err)
	}

	existingAccount := s.accountMapper.ToEntityFromResponse(*existingAccountDto)

	if !updateDto.Deposit {
		if existingAccount.Balance < updateDto.Amount {
			return nil, fmt.Errorf("you dont have enough funds for this transaction")
		}

		existingAccount.Balance -= updateDto.Amount
	} else {
		existingAccount.Balance += updateDto.Amount
	}

	updatedAccount, err := s.accountRepository.UpdateAccount(existingAccount)

	if err != nil {
		return nil, fmt.Errorf("update account: %w", err)
	}

	return s.accountMapper.ToResponse(updatedAccount), nil
}

func createAccountShell(clientId int) model.Account {
	return model.Account{
		ClientId:      clientId,
		AccountNumber: generateAccountNumber(),
		Balance:       0,
		Created_at:    time.Now().UTC(),
	}
}

func generateAccountNumber() string {
	const length = 16
	const digits = "0123456789"

	b := make([]byte, length)
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			panic(fmt.Errorf("failed to generate random digit: %w", err))
		}
		b[i] = digits[n.Int64()]
	}

	return string(b)
}
