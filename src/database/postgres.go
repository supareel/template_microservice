package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gomicro/src/ent"
	"log"
	"sync"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresDB interface {
	GetClient() (*ent.Client, error)
	MigrateDB(dbTimeOut int8)
}

type Database struct {
	client      *ent.Client
	databaseUrl string
	timeout     time.Duration
}

var (
	dbConnOnce sync.Once
	conn       *Database
)

func GetConnection(dbUsername string, dbPass string, dbServer string, dbName string, dbSsl string,
	dbPort int, dbTimeout int8) (*Database, error) {
	// Lock and unlock the entire GetConnection function
	dbConnOnce.Do(func() {
		conn = &Database{}
		fmt.Println("Initializing database connection")
		conn.timeout = time.Duration(dbTimeout)
		conn.databaseUrl = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", dbUsername, dbPass, dbServer, dbPort, dbName, dbSsl)
		fmt.Printf("\nDATABASE URL : %s\n\n", conn.databaseUrl)

		newdb, err := sql.Open("postgres", conn.databaseUrl)
		if err != nil {
			return
		}
		if err = newdb.Ping(); err != nil {
			return
		}
		drv := entsql.OpenDB(dialect.Postgres, newdb)
		conn.client = ent.NewClient(ent.Driver(drv))
	})

	return conn, nil
}

func (db *Database) GetClient() (*ent.Client, error) {
	if db == nil {
		return nil, errors.New("Database not instantiated")
	}
	return db.client, nil
}

// create schema in db
func (db *Database) CloseClient() {
	// run migration
	if err := db.client.Close(); err != nil {
		log.Fatal(err)
	}
}

// create schema in db
func (db *Database) MigrateDB(dbTimeOut int8) {
	fmt.Println("==========================================================")
	fmt.Println("=============     DB MIGRATION RUNNING     ===============")
	fmt.Println("==========================================================")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(dbTimeOut)*time.Second)
	defer cancel()
	// run migration
	if err := db.client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
}
