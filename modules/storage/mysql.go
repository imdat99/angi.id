package storage

import (
	"database/sql"
	"fmt"
	"log"

	c "angi.account/config"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func ConnectMySQL() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Acfg.Database.User, c.Acfg.Database.Password, c.Acfg.Database.Host, c.Acfg.Database.Port, c.Acfg.Database.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to MySQL: %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging MySQL: %v", err)
		return nil, err
	}

	log.Println("Connected to MySQL successfully")
	return db, nil
}
