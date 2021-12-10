package client

import (
	"github.com/gofrs/uuid"
)

var heyaClients = map[uuid.UUID][]uuid.UUID{}

func (cs *ClientServiceImpl) GetHeyaClientsIDByHeyaID(heyaID uuid.UUID) []uuid.UUID {
	return heyaClients[heyaID]
}

func (cs *ClientServiceImpl) AddHeyaClient(heyaID uuid.UUID, clientID uuid.UUID) {
	heyaClients[heyaID] = append(heyaClients[heyaID], clientID)
}

func (cs *ClientServiceImpl) DeleteHeyaClient(heyaID uuid.UUID, clientID uuid.UUID) error {
	for i, v := range heyaClients[heyaID] {
		if v == clientID {
			heyaClients[heyaID][i] = heyaClients[heyaID][len(heyaClients)-1]
			heyaClients[heyaID] = heyaClients[heyaID][:len(heyaClients)-1]
			return nil
		}
	}

	return ErrNotFound
}
