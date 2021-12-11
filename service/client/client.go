package client

import (
	"errors"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

var ErrNotFound = errors.New("client not found")

type ClientService interface {
	GetHeyaClientsIDByHeyaID(heyaID uuid.UUID) model.HeyaClients
	AddHeyaClient(heyaID uuid.UUID, client model.HeyaClient)
	DeleteHeyaClient(heyaID uuid.UUID, clientID uuid.UUID) error

	GetEditorClientsIDByHiqidashiID(hiqidashiID uuid.UUID) []uuid.UUID
	AddEditorClient(hiqidashiID uuid.UUID, clientID uuid.UUID)
	DeleteEditorClient(hiqidashiID uuid.UUID, clientID uuid.UUID) error
}

type ClientServiceImpl struct{}

func NewClientService() ClientService {
	return &ClientServiceImpl{}
}
