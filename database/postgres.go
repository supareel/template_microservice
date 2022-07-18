package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gomicro/ent"
	"gomicro/pkg"
	"log"
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
	client   *ent.Client
	database string
	timeout  time.Duration
}

func (db *Database) GetClient() (*ent.Client, error) {
	if db != nil {
		return db.client, errors.New("Database already instanciated")
	}
	return db.client, nil
}

// create schema in db
func (db *Database) MigrateDB(dbTimeOut int8) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(dbTimeOut)*time.Second)
	defer cancel()
	// run migration
	if err := db.client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
}

func NewPostgresClient(
	dbUsername string, dbPass string, dbServer string, dbName string,
	dbSsl string, dbPort int, dbTimeout int8) (Database, error) {
	var db Database
		db.timeout = time.Duration(dbTimeout)
		db.database = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", dbUsername, dbPass, dbServer, dbPort, dbName, dbSsl)
	fmt.Printf("\nDATABASE URL : %s\n\n", db.database)
	if db.client != nil {
		return db, nil
	}
	newdb, err := sql.Open("pgx", db.database)
	if err != nil {
		pkg.FancyHandleError(err)
		return db, err
	}
	err = newdb.Ping()
	pkg.FancyHandleError(err)
	drv := entsql.OpenDB(dialect.Postgres, newdb)
	db.client = ent.NewClient(ent.Driver(drv))
	return db, err
}
