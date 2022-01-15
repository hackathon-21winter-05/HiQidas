package repository

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type CredentialRepository interface {
	GetCredentialByMailAddress(ctx context.Context, mailAddress string) (*model.Credential, error)
	CreateCredential(ctx context.Context, credential *model.Credential) error
	DeleteCredentialByUserID(ctx context.Context, userID uuid.UUID) error
	UpdateCredentialByUserID(ctx context.Context, credential *model.Credential) error
}
