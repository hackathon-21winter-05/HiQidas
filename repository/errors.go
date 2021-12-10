package repository

import "errors"

var (
	// ErrNoRecordDeleted 削除するレコードがありませんでした
	ErrNoRecordDeleted = errors.New("no record deleted")
	// ErrNoRecordUpdated 更新するレコードがありませんでした
	ErrNoRecordUpdated = errors.New("no record updated")
	// ErrNillUUID UUIDがnullです
	ErrNillUUID = errors.New("nil uuid")
)
