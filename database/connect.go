package database

import (
	"context"
	"gomicro/ent"
	"gomicro/ent/migrate"
	"log"
)

func ConnectToDB() {

	client, err := ent.Open("postgres", "root:root@tcp(localhost:5432)/simple_bank")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
