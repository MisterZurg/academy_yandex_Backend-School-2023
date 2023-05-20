package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"io"
	"os"
)

func InitTables(db *sqlx.DB) error {
	f, err := os.Open("internal/db/postgres/postgres_init.sql")
	defer f.Close()
	if err != nil {
		return fmt.Errorf("cannot open file postgres_init.sql")
	}

	c, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("cannot read file postgres_init.sql")
	}

	// Create Tables
	_, err = db.Exec(string(c))

	return err
}
