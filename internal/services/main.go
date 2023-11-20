package services

import (
	"database/sql"
	"time"
)

var db*sql.DB

const dbTimeout = 5 * time.Second

type Models struct {
	Coffee Coffee
	JsonResponse JsonResponse
}

func New(dbPool * sql.DB) Models{
	db = dbPool
	return Models{}
}