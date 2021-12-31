package database

import (
	"context"
	"database/sql"
	"fmt"
	"gomicro/ent"
	"log"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
			log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}


func ConnectToDB() {
	DbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", 
	os.Getenv("DB_USERNAME"),	os.Getenv("DB_PASS"),	os.Getenv("DB_SERVER"),	
	os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL"))
	fmt.Printf("DATABASE : %s\n", DbUrl)
	client := Open(DbUrl)

	// Your code. For example:
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
			log.Fatal(err)
	}
	users, err := client.Accounts.Query().All(ctx)
	if err != nil {
			log.Fatal(err)
	}
	log.Println(users)
}
