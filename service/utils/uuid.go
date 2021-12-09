package utils

import "github.com/gofrs/uuid"

func GetUUID() uuid.UUID {
	return uuid.Must(uuid.NewV4())
}
