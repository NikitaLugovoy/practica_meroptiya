package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func AddDataBase() *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), "postgres://postgres:n1i2k3i4t5a6@localhost:5432/meropriyatiya")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to postgres")
	return pool
}

func RunMigration(pool *pgxpool.Pool) {

	files := []string{
		"migrations/created_table_user.sql",
		"migrations/created_table_group.sql",
		"migrations/created_table_event.sql",
		"migrations/created_table_event_image.sql",
		"migrations/created_table_event_participants.sql",
		"migrations/created_table_group_users.sql",
		"migrations/created_table_events_result.sql",
	}

	for _, file := range files {
		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		_, err = pool.Exec(context.Background(), string(sqlBytes))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Migration applied")
	}

}
