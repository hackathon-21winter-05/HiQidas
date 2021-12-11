package client

import (
	"sync"

	"github.com/gofrs/uuid"
)

type heyaClientsMap struct {
	cli map[uuid.UUID][]uuid.UUID
	sync.RWMutex
}

var heyaClients = heyaClientsMap{
	cli: make(map[uuid.UUID][]uuid.UUID),
}

func (cs *ClientServiceImpl) GetHeyaClientsIDByHeyaID(heyaID uuid.UUID) []uuid.UUID {
	heyaClients.RLock()
	defer heyaClients.RUnlock()

	return heyaClients.cli[heyaID]
}

func (cs *ClientServiceImpl) AddHeyaClient(heyaID uuid.UUID, clientID uuid.UUID) {
	heyaClients.Lock()
	defer heyaClients.Unlock()

	heyaClients.cli[heyaID] = append(heyaClients.cli[heyaID], clientID)
}

func (cs *ClientServiceImpl) DeleteHeyaClient(heyaID uuid.UUID, clientID uuid.UUID) error {
	heyaClients.Lock()
	defer heyaClients.Unlock()

	for i, v := range heyaClients.cli[heyaID] {
		if v == clientID {
			heyaClients.cli[heyaID][i] = heyaClients.cli[heyaID][len(heyaClients.cli[heyaID])-1]
			heyaClients.cli[heyaID] = heyaClients.cli[heyaID][:len(heyaClients.cli[heyaID])-1]
			return nil
		}
	}

	return ErrNotFound
}
