package database

import "errors"

var (
	// ErrNoRecordDeleted 削除するレコードがありませんでした
	ErrNoRecordDeleted = errors.New("no Record Deleted")
	// ErrNoRecordUpdated 更新するレコードがありませんでした
	ErrNoRecordUpdated = errors.New("no Record Updated")
)


