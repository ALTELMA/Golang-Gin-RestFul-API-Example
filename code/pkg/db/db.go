package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"

	// "github.com/jackc/pgx/v4/pgxpool"

	"open-market.com/user-api/pkg/helpers"
)

// DB function
func DB() *DB.sql {

	schema := os.Getenv("DB_SCHEMA")
	port := os.Getenv("DB_PORT")
	connection := os.Getenv("DB_CONNECTION")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_DATABASE")
	sslMode := os.Getenv("DB_SSL_MODE")
	sslCertificate := os.Getenv("DB_SSL_CERTIFICATE")
	sslPrivateKey := os.Getenv("DB_SSL_PRIVATE_KEY")
	sslRootCert := os.Getenv("DB_SSL_ROOT_CA")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s sslcert=%s sslkey=%s sslrootcert=%s",
		host,
		port,
		user,
		password,
		dbName,
		sslMode,
		schema,
		sslCertificate,
		sslPrivateKey,
		sslRootCert,
	)

	// db, err := pgx.Connect(context.Background(), connStr)
	// db, err := pgxpool.Connect(context.Background(), connStr)
	db, err := sql.Open(connection, connStr)
	helpers.CheckErr(err)

	err = db.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(1800 * time.Second)

	return db
}
