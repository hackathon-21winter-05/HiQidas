package client

import (
	"errors"

	"github.com/gofrs/uuid"
)

var ErrNotFound = errors.New("client not found")

type ClientService interface {
	GetHeyaClientsIDByHeyaID(heyaID uuid.UUID) []uuid.UUID
	AddHeyaClient(heyaID uuid.UUID, clientID uuid.UUID)
	DeleteHeyaClient(heyaID uuid.UUID, clientID uuid.UUID) error

	GetEditorClientsIDByHiqidashiID(hiqidashiID uuid.UUID) []uuid.UUID
	AddEditorClient(hiqidashiID uuid.UUID, clientID uuid.UUID)
	DeleteEditorClient(hiqidashiID uuid.UUID, clientID uuid.UUID) error
}

type ClientServiceImpl struct{}

func NewClientService() ClientService {
	return &ClientServiceImpl{}
}
