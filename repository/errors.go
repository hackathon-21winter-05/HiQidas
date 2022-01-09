package repository

import "errors"

var (
	// ErrNoRecordDeleted 削除するレコードがありませんでした
	ErrNoRecordDeleted = errors.New("no record deleted")
	// ErrNoRecordUpdated 更新するレコードがありませんでした
	ErrNoRecordUpdated = errors.New("no record updated")
	// ErrNillUUID UUIDがnullです
	ErrNillUUID = errors.New("nil uuid")
	// ErrEmptyString 許可されない空文字が存在します
	ErrEmptyString = errors.New("empty string")
	// ErrNotFound 取得したレコードがありませんでした
	ErrNotFound = errors.New("record not found")
	// ErrInconsistent 与えられたデータの中で一致すべき部分が一致していません
	ErrInconsistent = errors.New("inconsistent data")
)
