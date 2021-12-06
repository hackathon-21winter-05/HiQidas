package database

import "errors"

var (
	// ErrNoRecordDeleted 削除するレコードがありませんでした
	ErrNoRecordDeleted = errors.New("no Record Deleted")
)


