package mapper

import (
	accountdto "basic-gin/internal/dto/accountDto"
	"basic-gin/internal/model"
)

type AccountMapper struct {
}

func (m *AccountMapper) ToEntityFromResponse(dto accountdto.ResponseAccountDto) *model.Account {
	return &model.Account{
		Id:            dto.Id,
		ClientId:      dto.ClientId,
		AccountNumber: dto.AccountNumber,
		Balance:       dto.Balance,
		Created_at:    dto.Created_at,
	}
}

func (m *AccountMapper) ToResponse(account *model.Account) *accountdto.ResponseAccountDto {
	return &accountdto.ResponseAccountDto{
		Id:            account.Id,
		ClientId:      account.ClientId,
		AccountNumber: account.AccountNumber,
		Balance:       account.Balance,
		Created_at:    account.Created_at,
	}
}

func (m *AccountMapper) ToResponseSlice(accounts []*model.Account) []*accountdto.ResponseAccountDto {
	responseSlice := make([]*accountdto.ResponseAccountDto, 0, len(accounts))

	for i := 0; i < len(accounts); i++ {
		account := accounts[i]
		accountResponse := m.ToResponse(account)
		responseSlice = append(responseSlice, accountResponse)
	}

	return responseSlice
}
