package repository

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository struct {
	*sqlx.DB
	*CouriersRepository
	*OrdersRepository
}

func NewRepository(dsn string) (*Repository, error) {
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatal("Cannot connect to Postgres")
	}

	return &Repository{
		DB:                 db,
		CouriersRepository: &CouriersRepository{db},
		OrdersRepository:   &OrdersRepository{db},
	}, nil
}
