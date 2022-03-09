package main

import (
	"context"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/smith-30/goparco/ent_playground/ent"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent.db?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	//
	// 初回用
	//
	// オートマイグレーションツールを実行する
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	//
	// migration
	//
	// ctx := context.Background()
	// // Create a local migration directory.
	// dir, err := migrate.NewLocalDir("migrations")
	// if err != nil {
	// 	log.Fatalf("failed creating atlas migration directory: %v", err)
	// }
	// // Write migration diff.
	// err = client.Schema.Diff(ctx, schema.WithDir(dir))
	// if err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }
}
