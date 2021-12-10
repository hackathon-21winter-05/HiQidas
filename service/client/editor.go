package client

import (
	"github.com/gofrs/uuid"
)

var editorClients = map[uuid.UUID][]uuid.UUID{}

func (cs *ClientServiceImpl) GetEditorClientsIDByHiqidashiID(hiqidashiID uuid.UUID) []uuid.UUID {
	return editorClients[hiqidashiID]
}

func (cs *ClientServiceImpl) AddEditorClient(hiqidashiID uuid.UUID, clientID uuid.UUID) {
	editorClients[hiqidashiID] = append(editorClients[hiqidashiID], clientID)
}

func (cs *ClientServiceImpl) DeleteEditorClient(hiqidashiID uuid.UUID, clientID uuid.UUID) error {
	for i, v := range editorClients[hiqidashiID] {
		if v == clientID {
			editorClients[hiqidashiID][i] = editorClients[hiqidashiID][len(editorClients)-1]
			editorClients[hiqidashiID] = editorClients[hiqidashiID][:len(editorClients)-1]
			return nil
		}
	}

	return ErrNotFound
}
