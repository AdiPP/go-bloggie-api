package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Renos-id/go-starter-template/database"
	"github.com/golang-migrate/migrate/v4"
	pg_migrate "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var migrateUpCmd *cobra.Command

func init() {
	migrateUpCmd = &cobra.Command{
		Use:   "up",
		Short: "migrate to v1 command",
		Long:  "migrate to install version 1 of our application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate up command")
			dbConn := database.Open()
			dbDriver, err := pg_migrate.WithInstance(dbConn.DB, &pg_migrate.Config{
				MigrationsTable: os.Getenv("MIGRATION_TABLE_NAME"),
			})
			if err != nil {
				log.Fatal("Instance Error: ", err.Error())
			}
			fileSource, err := (&file.File{}).Open("file://database/migrations")
			if err != nil {
				log.Fatal("Opening file Error: ", err.Error())
			}
			m, err := migrate.NewWithInstance("file", fileSource, os.Getenv("DB_DATABASE"), dbDriver)
			if err != nil {
				log.Fatal("migrate error: ", err)
			}

			if err = m.Up(); err != nil {
				log.Fatal("migrate up error: ", err)
			}

			fmt.Println("Migrate up done with success")
		},
	}
	MigrateCmd.AddCommand(migrateUpCmd)
}
