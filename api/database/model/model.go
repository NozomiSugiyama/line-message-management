package model

import (
	"errors"
)

var (
	UserTableName     = "public.users"
	NonceTableName    = "public.nonces"
	LineUserTableName = "public.line_users"
)

var (
	// ErrRecordNotFound returns a "record not found error". Occurs only when attempting to query the database with a struct; querying with a slice won't return this error
	ErrRecordNotFound = errors.New("record not found")
)
