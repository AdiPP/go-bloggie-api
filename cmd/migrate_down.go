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

var migrateDownCmd *cobra.Command

func init() {
	migrateDownCmd = &cobra.Command{
		Use:   "down",
		Short: "migrate from v2 to v1",
		Long:  "command to downgrade database from v2 to v1",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate Down command")
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

			if err = m.Down(); err != nil {
				log.Fatal("migrate Down error: ", err)
			}

			fmt.Println("Migrate down done with success")
		},
	}
	MigrateCmd.AddCommand(migrateDownCmd)
}
