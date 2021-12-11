package model

import (
	"sync"

	"github.com/gofrs/uuid"
)

type HeyaClientsMap struct {
	Clients map[uuid.UUID]HeyaClients
	sync.RWMutex
}

type HeyaClients map[uuid.UUID]HeyaClient

type HeyaClient struct {
	ID                 uuid.UUID
	UserID             uuid.UUID
	EditingHiqidashiID uuid.UUID
}
