package mapper

import (
	clientdto "basic-gin/internal/dto/clientDto"
	"basic-gin/internal/model"
)

type ClientMapper struct{}

func (m *ClientMapper) ToEntityFromSaveDto(saveDto *clientdto.SaveClientDto) model.Client {
	return model.Client{
		FirstName:        saveDto.FirstName,
		LastName:         saveDto.LastName,
		Email:            saveDto.Email,
		ResidenceAddress: saveDto.ResidenceAddress,
		BirthDate:        saveDto.BirthDate,
	}
}

func (m *ClientMapper) ToEntityFromUpdateDto(updateDto *clientdto.UpdateClientDto) model.Client {
	return model.Client{
		ID:               updateDto.Id,
		FirstName:        updateDto.FirstName,
		LastName:         updateDto.LastName,
		Email:            updateDto.Email,
		ResidenceAddress: updateDto.ResidenceAddress,
		BirthDate:        updateDto.BirthDate,
	}
}

func (m *ClientMapper) ToResponse(client *model.Client) clientdto.ResponseClientDto {
	return clientdto.ResponseClientDto{
		ID:               client.ID,
		FirstName:        client.FirstName,
		LastName:         client.LastName,
		Email:            client.Email,
		ResidenceAddress: client.ResidenceAddress,
		BirthDate:        client.BirthDate,
	}
}

func (m *ClientMapper) ToResponseSlice(clients []*model.Client) []*clientdto.ResponseClientDto {
	responseSlice := make([]*clientdto.ResponseClientDto, 0, len(clients))

	for i := 0; i < len(clients); i++ {
		client := clients[i]
		resp := m.ToResponse(client)
		responseSlice = append(responseSlice, &resp)
	}

	return responseSlice
}
