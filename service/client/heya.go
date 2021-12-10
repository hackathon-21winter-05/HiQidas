package client

import (
	"errors"

	"github.com/gofrs/uuid"
)

var (
	ErrNotFound = errors.New("client not found")

	heyaClients = map[uuid.UUID][]uuid.UUID{}
)

type HeyaClientService interface {
	GetHeyaClientsIDByHeyaID(heyaID uuid.UUID) []uuid.UUID
	AddHeyaClient(heyaID uuid.UUID, clientID uuid.UUID)
	DeleteHeyaClient(heyaID uuid.UUID, clientID uuid.UUID) error
}

type HeyaClientServiceImpl struct{}

func NewClientService() HeyaClientService {
	return &HeyaClientServiceImpl{}
}

func (cs *HeyaClientServiceImpl) GetHeyaClientsIDByHeyaID(heyaID uuid.UUID) []uuid.UUID {
	return heyaClients[heyaID]
}

func (cs *HeyaClientServiceImpl) AddHeyaClient(heyaID uuid.UUID, clientID uuid.UUID) {
	heyaClients[heyaID] = append(heyaClients[heyaID], clientID)
}

func (cs *HeyaClientServiceImpl) DeleteHeyaClient(heyaID uuid.UUID, clientID uuid.UUID) error {
	for i, v := range heyaClients[heyaID] {
		if v == clientID {
			heyaClients[heyaID][i] = heyaClients[heyaID][len(heyaClients)-1]
			heyaClients[heyaID] = heyaClients[heyaID][:len(heyaClients)-1]
			return nil
		}
	}

	return ErrNotFound
}
