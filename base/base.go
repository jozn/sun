package base

import (
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var __DEV__ bool = true

type AppErr error
