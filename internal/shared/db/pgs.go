package db

import (
	"database/sql"
	"fmt"
	"log"

	c "angi.id/internal/shared"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectPostgres() (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&search_path=angi_id",
		c.Acfg.Database.User, c.Acfg.Database.Password, c.Acfg.Database.Host, c.Acfg.Database.Port, c.Acfg.Database.Name)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Error connecting to Postgres: %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging Postgres: %v", err)
		return nil, err
	}

	log.Println("Connected to Postgres successfully")
	return db, nil
}
