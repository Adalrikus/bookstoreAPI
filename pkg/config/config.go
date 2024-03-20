package config

import (
	"context"
	"database/sql"
	"log"

	"github.com/adalrikus/bookstoreAPI/pkg/models"
	_ "github.com/mattn/go-sqlite3"
)

const schema = `
CREATE TABLE books (
  id        INTEGER PRIMARY KEY,
  Title     TEXT    NOT NULL,
  Author    TEXT    NOT NULL,
  Publisher TEXT    NOT NULL
);
`

func Connect(ctx context.Context) (*models.Queries, error) {
	db, err := sql.Open("sqlite3", "sqlite3.db")
	if err != nil {
		return nil, err
	}

	// create tables
	if _, err := db.ExecContext(ctx, string(schema)); err != nil {
		log.Println("[WARN]", err)
	}

	queries := models.New(db)
	return queries, nil
}
