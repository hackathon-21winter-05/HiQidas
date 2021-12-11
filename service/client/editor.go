package client

import (
	"sync"

	"github.com/gofrs/uuid"
)

type editorClientsMap struct {
	cli map[uuid.UUID][]uuid.UUID
	sync.RWMutex
}

var editorClients = editorClientsMap{
	cli: make(map[uuid.UUID][]uuid.UUID),
}

func (cs *ClientServiceImpl) GetEditorClientsIDByHiqidashiID(hiqidashiID uuid.UUID) []uuid.UUID {
	editorClients.RLock()
	defer editorClients.RUnlock()

	return editorClients.cli[hiqidashiID]
}

func (cs *ClientServiceImpl) AddEditorClient(hiqidashiID uuid.UUID, clientID uuid.UUID) {
	editorClients.Lock()
	defer editorClients.Unlock()

	editorClients.cli[hiqidashiID] = append(editorClients.cli[hiqidashiID], clientID)
}

func (cs *ClientServiceImpl) DeleteEditorClient(hiqidashiID uuid.UUID, clientID uuid.UUID) error {
	editorClients.Lock()
	defer editorClients.Unlock()

	for i, v := range editorClients.cli[hiqidashiID] {
		if v == clientID {
			editorClients.cli[hiqidashiID][i] = editorClients.cli[hiqidashiID][len(editorClients.cli[hiqidashiID])-1]
			editorClients.cli[hiqidashiID] = editorClients.cli[hiqidashiID][:len(editorClients.cli[hiqidashiID])-1]
			return nil
		}
	}

	return ErrNotFound
}
