package client

import (
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

var (
	heyaClientsMap = model.HeyaClientsMap{
		Clients: make(map[uuid.UUID]model.HeyaClients),
	}
)

func (cs *ClientServiceImpl) GetHeyaClientsIDByHeyaID(heyaID uuid.UUID) model.HeyaClients {
	heyaClientsMap.RLock()
	defer heyaClientsMap.RUnlock()

	return heyaClientsMap.Clients[heyaID]
}

func (cs *ClientServiceImpl) AddHeyaClient(heyaID uuid.UUID, client model.HeyaClient) {
	heyaClientsMap.Lock()
	defer heyaClientsMap.Unlock()

	heyaClientsMap.Clients[heyaID][client.ID] = client
}

func (cs *ClientServiceImpl) DeleteHeyaClient(heyaID uuid.UUID, clientID uuid.UUID) error {
	heyaClientsMap.Lock()
	defer heyaClientsMap.Unlock()

	delete(heyaClientsMap.Clients[heyaID], clientID)

	return ErrNotFound
}
