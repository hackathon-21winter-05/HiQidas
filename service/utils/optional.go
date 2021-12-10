package utils

import (
	"database/sql"

	"github.com/gofrs/uuid"
)

func NullStringFrom(s string) sql.NullString {
	return sql.NullString{String: s, Valid: true}
}

func NullString() sql.NullString {
	return sql.NullString{String: "", Valid: false}
}

func NullUUIDFrom(id uuid.UUID) uuid.NullUUID {
	return uuid.NullUUID{UUID: id, Valid: true}
}

func NullUUID() uuid.NullUUID {
	return uuid.NullUUID{UUID: uuid.Nil, Valid: false}
}
