package service

import (
	clientdto "basic-gin/internal/dto/clientDto"
	clientMapper "basic-gin/internal/mapper"
	"basic-gin/internal/model"
	"basic-gin/internal/repository"
	"context"
	"fmt"
	"strings"
)

type ClientService struct {
	clientRepository *repository.ClientRepository
	clientMapper     *clientMapper.ClientMapper
}

func NewClientService(repo *repository.ClientRepository) *ClientService {
	return &ClientService{
		clientRepository: repo,
		clientMapper:     &clientMapper.ClientMapper{},
	}
}

func (s *ClientService) GetClients(ctx context.Context) ([]*clientdto.ResponseClientDto, error) {
	clients, err := s.clientRepository.GetClients(ctx)

	if err != nil {
		return nil, fmt.Errorf("get clients: %w", err)
	}

	response := s.clientMapper.ToResponseSlice(clients)

	return response, nil
}

func (s *ClientService) GetClientById(ctx context.Context, id int) (*clientdto.ResponseClientDto, error) {
	if id < 0 {
		return nil, fmt.Errorf("id parameter cant be nil or negative")
	}

	client, err := s.clientRepository.GetClientById(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("get client by id %d: :%w", id, err)
	}

	response := s.clientMapper.ToResponse(client)

	return &response, nil
}

func (s *ClientService) CreateClient(ctx context.Context, saveDto *clientdto.SaveClientDto) (*clientdto.ResponseClientDto, error) {
	if saveDto == nil {
		return nil, fmt.Errorf("client dto cant be nil")
	}

	client := s.clientMapper.ToEntityFromSaveDto(saveDto)

	if err := validateClient(&client); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	savedClient, err := s.clientRepository.CreateClient(ctx, &client)

	if err != nil {
		return nil, fmt.Errorf("create client: %w", err)
	}

	response := s.clientMapper.ToResponse(savedClient)
	return &response, nil
}

func (s *ClientService) UpdateClient(ctx context.Context, updateDto *clientdto.UpdateClientDto) (*clientdto.ResponseClientDto, error) {
	if updateDto == nil {
		return nil, fmt.Errorf("client dto cant be nil")
	}

	clientForUpdate := s.clientMapper.ToEntityFromUpdateDto(updateDto)

	if _, err := s.clientRepository.GetClientById(ctx, clientForUpdate.ID); err != nil {
		return nil, fmt.Errorf("get client by id: %w", err)
	}

	if err := validateClient(&clientForUpdate); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	updatedClient, err := s.clientRepository.UpdateClient(ctx, &clientForUpdate)

	if err != nil {
		return nil, fmt.Errorf("update client: %w", err)
	}

	response := s.clientMapper.ToResponse(updatedClient)
	return &response, nil
}

func (s *ClientService) DeleteClient(ctx context.Context, id int) error {
	if id < 0 {
		return fmt.Errorf("id cant be zero or negative number")
	}

	if _, err := s.clientRepository.GetClientById(ctx, id); err != nil {
		return fmt.Errorf("get client by id: %w", err)
	}

	if err := s.clientRepository.DeleteClient(ctx, id); err != nil {
		return fmt.Errorf("delete client: %w", err)
	}

	return nil
}

func validateClient(client *model.Client) error {
	if client.FirstName == "" {
		return fmt.Errorf("first name cant be empty")
	}

	if client.LastName == "" {
		return fmt.Errorf("last name cant be empty")
	}

	if client.Email == "" {
		return fmt.Errorf("email cant be empty")
	}

	if !strings.Contains(client.Email, "@") {
		return fmt.Errorf("invalid email format")
	}

	if client.ResidenceAddress == "" {
		return fmt.Errorf("residence address cant be empty")
	}

	if client.BirthDate == "" {
		return fmt.Errorf("birth date cant be empty")
	}

	return nil
}
