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

var Conn *ent.Client

func ConnectToDB() {
	DbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASS"), os.Getenv("DB_SERVER"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL"))

	fmt.Printf("DATABASE URL : %s\n\n", DbUrl)

	db, err := sql.Open("pgx", DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	Conn = ent.NewClient(ent.Driver(drv))

	ctx := context.Background()
	if err := Conn.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
	demo_query()
}

func demo_query() {
	ctx := context.Background()
	users, err := Conn.Accounts.Query().All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(users)
}
