package client

import (
	"sync"

	"github.com/gofrs/uuid"
)

var editorClients = map[uuid.UUID][]uuid.UUID{}
var editorClientsMutex = &sync.RWMutex{}

func (cs *ClientServiceImpl) GetEditorClientsIDByHiqidashiID(hiqidashiID uuid.UUID) []uuid.UUID {
	editorClientsMutex.RLock()
	defer editorClientsMutex.RUnlock()
	return editorClients[hiqidashiID]
}

func (cs *ClientServiceImpl) AddEditorClient(hiqidashiID uuid.UUID, clientID uuid.UUID) {
	editorClientsMutex.Lock()
	defer editorClientsMutex.Unlock()
	editorClients[hiqidashiID] = append(editorClients[hiqidashiID], clientID)
}

func (cs *ClientServiceImpl) DeleteEditorClient(hiqidashiID uuid.UUID, clientID uuid.UUID) error {
	for i, v := range editorClients[hiqidashiID] {
		if v == clientID {
			editorClientsMutex.Lock()
			defer editorClientsMutex.Unlock()
			editorClients[hiqidashiID][i] = editorClients[hiqidashiID][len(editorClients[hiqidashiID])-1]
			editorClients[hiqidashiID] = editorClients[hiqidashiID][:len(editorClients[hiqidashiID])-1]
			return nil
		}
	}

	return ErrNotFound
}
