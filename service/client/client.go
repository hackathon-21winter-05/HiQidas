package client

import (
	"errors"

	"github.com/gofrs/uuid"
)

var (
	ErrNotFound = errors.New("client not found")

	heyaClients = map[uuid.UUID][]uuid.UUID{}
)

type ClientService interface {
	GetClientsIDByHeyaID(heyaID uuid.UUID) []uuid.UUID
	AddClient(heyaID uuid.UUID, clientID uuid.UUID)
	DeleteClient(heyaID uuid.UUID, clientID uuid.UUID) error
}

type ClientServiceImpl struct{}

func NewClientService() ClientService {
	return &ClientServiceImpl{}
}

func (cs *ClientServiceImpl) GetClientsIDByHeyaID(heyaID uuid.UUID) []uuid.UUID {
	return heyaClients[heyaID]
}

func (cs *ClientServiceImpl) AddClient(heyaID uuid.UUID, clientID uuid.UUID) {
	heyaClients[heyaID] = append(heyaClients[heyaID], clientID)
}

func (cs *ClientServiceImpl) DeleteClient(heyaID uuid.UUID, clientID uuid.UUID) error {
	for i, v := range heyaClients[heyaID] {
		if v == clientID {
			heyaClients[heyaID][i] = heyaClients[heyaID][len(heyaClients)-1]
			heyaClients[heyaID] = heyaClients[heyaID][:len(heyaClients)-1]
			return nil
		}
	}

	return ErrNotFound
}
