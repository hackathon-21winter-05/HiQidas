package client

import (
	"sync"

	"github.com/gofrs/uuid"
)

var heyaClients = map[uuid.UUID][]uuid.UUID{}
var heyaClientsMutex = &sync.RWMutex{}

func (cs *ClientServiceImpl) GetHeyaClientsIDByHeyaID(heyaID uuid.UUID) []uuid.UUID {
	heyaClientsMutex.RLock()
	defer heyaClientsMutex.RUnlock()
	return heyaClients[heyaID]
}

func (cs *ClientServiceImpl) AddHeyaClient(heyaID uuid.UUID, clientID uuid.UUID) {
	heyaClientsMutex.Lock()
	defer heyaClientsMutex.Unlock()
	heyaClients[heyaID] = append(heyaClients[heyaID], clientID)
}

func (cs *ClientServiceImpl) DeleteHeyaClient(heyaID uuid.UUID, clientID uuid.UUID) error {
	for i, v := range heyaClients[heyaID] {
		if v == clientID {
			heyaClientsMutex.Lock()
			defer heyaClientsMutex.Unlock()
			heyaClients[heyaID][i] = heyaClients[heyaID][len(heyaClients[heyaID])-1]
			heyaClients[heyaID] = heyaClients[heyaID][:len(heyaClients[heyaID])-1]
			return nil
		}
	}

	return ErrNotFound
}
