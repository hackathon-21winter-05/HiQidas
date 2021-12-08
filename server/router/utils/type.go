package utils

import "github.com/gofrs/uuid"

func UuidsToStrings(IDs []uuid.UUID) []string {
	var res []string

	for _, ID := range IDs {
		res = append(res, ID.String())
	}

	return res
}
