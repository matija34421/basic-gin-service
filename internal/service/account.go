package service

import (
	accountdto "basic-gin/internal/dto/accountDto"
	accountMapper "basic-gin/internal/mapper"
	"basic-gin/internal/model"
	"basic-gin/internal/repository"
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

type AccountService struct {
	accountRepository *repository.AccountRepository
	accountMapper     *accountMapper.AccountMapper
	clientRepository  *repository.ClientRepository
}

func NewAccountService(accountRepo *repository.AccountRepository, clientRepo *repository.ClientRepository) *AccountService {
	return &AccountService{
		accountRepository: accountRepo,
		accountMapper:     &accountMapper.AccountMapper{},
		clientRepository:  clientRepo,
	}
}

func (s *AccountService) GetAccountsByClientId(ctx context.Context, clientId int) ([]*accountdto.ResponseAccountDto, error) {
	if clientId < 0 {
		return nil, fmt.Errorf("id cant be negative number")
	}

	var (
		wg      sync.WaitGroup
		errOnce error
		errMu   sync.Mutex

		client   *model.Client
		accounts []*model.Account
	)

	setErr := func(e error) {
		if e == nil {
			return
		}
		errMu.Lock()
		if errOnce == nil {
			errOnce = e
		}
		errMu.Unlock()
	}

	wg.Add(2)

	go func() {
		defer wg.Done()

		c, err := s.clientRepository.GetClientById(ctx, clientId)

		if err != nil {
			setErr(fmt.Errorf("get client by id: %w", err))
			return
		}

		client = c
	}()

	go func() {
		defer wg.Done()

		a, err := s.accountRepository.GetAccountsByClientId(ctx, clientId)

		if err != nil {
			setErr(fmt.Errorf("get accounts by client id: %w", err))
			return
		}

		accounts = a
	}()

	wg.Wait()

	if errOnce != nil {
		return nil, errOnce
	}

	if client == nil {
		return nil, fmt.Errorf("client not found")
	}

	return s.accountMapper.ToResponseSlice(accounts), nil
}

func (s *AccountService) GetAccountById(ctx context.Context, id int) (*accountdto.ResponseAccountDto, error) {
	if id < 0 {
		return nil, fmt.Errorf("id cant be negative number")
	}

	account, err := s.accountRepository.GetAccountById(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("get account by id: %w", err)
	}

	return s.accountMapper.ToResponse(account), nil
}

func (s *AccountService) CreateAccount(ctx context.Context, clientId int) (*accountdto.ResponseAccountDto, error) {
	if clientId < 0 {
		return nil, fmt.Errorf("invalid clientId")
	}

	account := createAccountShell(clientId)

	savedAccount, err := s.accountRepository.CreateAccount(ctx, &account)

	if err != nil {
		return nil, fmt.Errorf("create account: %w", err)
	}

	return s.accountMapper.ToResponse(savedAccount), nil
}

func (s *AccountService) UpdateAccount(ctx context.Context, updateDto accountdto.UpdateAccountDto) (*accountdto.ResponseAccountDto, error) {

	existingAccountDto, err := s.GetAccountById(ctx, updateDto.Id)

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

	updatedAccount, err := s.accountRepository.UpdateAccount(ctx, existingAccount)

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
