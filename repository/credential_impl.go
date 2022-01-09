package repository

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

// GetCredentialByUserID UserIDからCredentialを取得する
func (repo *GormRepository) GetCredentialByUserID(ctx context.Context, userID uuid.UUID) (*model.Credential, error) {
	db, err := repo.getDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	var credential *model.Credential
	err = db.
		Where("user_id = ?", userID).
		Joins("User").
		First(credential).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get credential by userID : %w", err)
	}

	return credential, nil
}

// CreateCredential Credentialを作成
func (repo *GormRepository) CreateCredential(ctx context.Context, credential *model.Credential) error {
	if credential.UserID == uuid.Nil || credential.User.ID == uuid.Nil || credential.UserID != credential.User.ID {
		return ErrNillUUID
	}
	if credential.MailAddress == "" || credential.HashedPass == "" || credential.User.Name == "" {
		return ErrEmptyString
	}
	if credential.UserID != credential.User.ID {
		return ErrInconsistent
	}

	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	err = db.Create(&credential).Error
	if err != nil {
		return fmt.Errorf("failed to create credential :%w", err)
	}

	return nil
}

// DeleteCredentialByUserID Credentialを削除する
func (repo *GormRepository) DeleteCredentialByUserID(ctx context.Context, userID uuid.UUID) error {
	if userID == uuid.Nil {
		return ErrNillUUID
	}

	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	result := db.
		Where("user_id = ?", userID).
		Delete(&model.Credential{})
	err = result.Error
	if err != nil {
		return fmt.Errorf("failed to delete credential :%w", err)
	}
	if result.RowsAffected == 0 {
		return ErrNoRecordDeleted
	}

	return nil
}

// UpdateCredentialByUserID ユーザーの認証情報を更新
func (repo *GormRepository) UpdateCredentialByUserID(ctx context.Context, credential *model.NullCredential) error {
	if credential.UserID == uuid.Nil || credential.User.ID == uuid.Nil {
		return ErrNillUUID
	}
	if credential.UserID != credential.User.ID {
		return ErrInconsistent
	}

	db, err := repo.getDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}

	userMap := map[string]interface{}{}
	userMap["id"] = credential.User.ID
	if credential.User.Name.Valid {
		userMap["name"] = credential.User.Name
	}

	credentialMap := map[string]interface{}{}
	credentialMap["user_id"] = credential.UserID
	if credential.MailAddress.Valid {
		credentialMap["mail_address"] = credential.MailAddress
	}
	if credential.HashedPass.Valid {
		credentialMap["hashed_pass"] = credential.HashedPass
	}
	credentialMap["user"] = userMap

	result := db.
		Model(model.Credential{}).
		Where("user_id = ?", credential.UserID).
		Updates(credentialMap)
	err = result.Error
	if err != nil {
		return fmt.Errorf("failed to update credential : %w", err)
	}
	if result.RowsAffected == 0 {
		return ErrNoRecordUpdated
	}

	return nil
}
